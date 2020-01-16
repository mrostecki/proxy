#include <dlfcn.h>

#include "envoy/network/listen_socket.h"
#include "envoy/registry/registry.h"
#include "envoy/server/filter_config.h"

#include "common/buffer/buffer_impl.h"
#include "common/common/assert.h"
#include "common/common/fmt.h"

#include "cilium/api/network_filter.pb.validate.h"
#include "cilium/network_filter.h"
#include "cilium/network_policy.h"
#include "cilium/socket_option.h"

namespace Envoy {
namespace Server {
namespace Configuration {

/**
 * Config registration for the bpf metadata filter. @see
 * NamedNetworkFilterConfigFactory.
 */
class CiliumNetworkConfigFactory : public NamedNetworkFilterConfigFactory {
public:
  // NamedNetworkFilterConfigFactory
  Network::FilterFactoryCb
  createFilterFactoryFromProto(const Protobuf::Message& proto_config,
			       FactoryContext& context) override {
    auto config = std::make_shared<Filter::CiliumL3::Config>(MessageUtil::downcastAndValidate<const ::cilium::NetworkFilter&>(proto_config, context.messageValidationVisitor()), context);
    return [config](Network::FilterManager &filter_manager) mutable -> void {
      filter_manager.addFilter(std::make_shared<Filter::CiliumL3::Instance>(config));
    };
  }

  Network::FilterFactoryCb
  createFilterFactory(const Json::Object& json_config, FactoryContext& context) {
    auto config = std::make_shared<Filter::CiliumL3::Config>(json_config, context);
    return [config](Network::FilterManager &filter_manager) mutable -> void {
      filter_manager.addFilter(std::make_shared<Filter::CiliumL3::Instance>(config));
    };
  }
  ProtobufTypes::MessagePtr createEmptyConfigProto() override {
    return std::make_unique<::cilium::NetworkFilter>();
  }

  std::string name() override { return "cilium.network"; }
};

/**
 * Static registration for the bpf metadata filter. @see RegisterFactory.
 */
static Registry::RegisterFactory<CiliumNetworkConfigFactory, NamedNetworkFilterConfigFactory>
    registered_;

} // namespace Configuration
} // namespace Server

namespace Filter {
namespace CiliumL3 {

Config::Config(const ::cilium::NetworkFilter& config, Server::Configuration::FactoryContext&) {
  if (config.proxylib().length() > 0) {
    proxylib_ = std::make_shared<Cilium::GoFilter>(config.proxylib(), config.proxylib_params());
  }
  if (config.policy_name() != "" || config.l7_proto() != "") {
    throw EnvoyException(fmt::format("network: 'policy_name' and 'go_proto' are no longer supported: \'{}\'", config.DebugString()));
  }
}
  
Config::Config(const Json::Object&, Server::Configuration::FactoryContext&) {} // Dummy, not used.

Network::FilterStatus Instance::onNewConnection() {
  ENVOY_LOG(debug, "Cilium Network: onNewConnection");
  auto& conn = callbacks_->connection();
  const auto option = Cilium::GetSocketOption(conn.socketOptions());
  if (option) {
    maps_ = option->GetProxyMap();
    if (maps_) {
      // Insert connection callback to delete the proxymap entry once the connection is closed.
      proxy_port_ = option->proxy_port_;
      if (proxy_port_ != 0) {
	conn.addConnectionCallbacks(*this);
	ENVOY_CONN_LOG(debug, "Cilium Network: Added connection callbacks to delete proxymap entry later", conn);
      }
    }

    const std::string& policy_name = option->pod_ip_;
    std::string l7proto;
    if (config_->proxylib_.get() != nullptr && option->policy_ &&
	(option->policy_->useProxylib(option->ingress_, option->port_,
				      option->ingress_ ? option->identity_ : option->destination_identity_,
				      l7proto))) {
      go_parser_ = config_->proxylib_->NewInstance(conn, l7proto, option->ingress_, option->identity_,
						   option->destination_identity_, conn.remoteAddress()->asString(),
						   conn.localAddress()->asString(), policy_name);
      if (go_parser_.get() == nullptr) {
	ENVOY_CONN_LOG(warn, "Cilium Network: Go parser \"{}\" not found", conn, l7proto);
	return Network::FilterStatus::StopIteration;
      }
    }
  } else {
    ENVOY_CONN_LOG(warn, "Cilium Network: Cilium Socket Option not found", conn);
  }

  return Network::FilterStatus::Continue;
}

Network::FilterStatus Instance::onData(Buffer::Instance& data, bool end_stream) {
  if (go_parser_) {
    FilterResult res = go_parser_->OnIO(false, data, end_stream); // 'false' marks original direction data
    ENVOY_CONN_LOG(trace, "Cilium Network::onData: \'GoFilter::OnIO\' returned {}", callbacks_->connection(), res);

    if (res != FILTER_OK) {
      // Drop the connection due to an error
      go_parser_->Close();
      return Network::FilterStatus::StopIteration;
    }

    if (go_parser_->WantReplyInject()) {
      auto& conn = callbacks_->connection();
      ENVOY_CONN_LOG(trace, "Cilium Network::onData: calling write() on an empty buffer", conn);

      // We have no idea when, if ever new data will be received on the
      // reverse direction. Connection write on an empty buffer will cause
      // write filter chain to be called, and gives our write path the
      // opportunity to inject data.
      Buffer::OwnedImpl empty;
      conn.write(empty, false);  
    }

    go_parser_->SetOrigEndStream(end_stream);
  }

  return Network::FilterStatus::Continue;
}

Network::FilterStatus Instance::onWrite(Buffer::Instance& data, bool end_stream) {
  if (go_parser_) {
    FilterResult res = go_parser_->OnIO(true, data, end_stream); // 'true' marks reverse direction data
    ENVOY_CONN_LOG(trace, "Cilium Network::OnWrite: \'GoFilter::OnIO\' returned {}", callbacks_->connection(), res);
  
    if (res != FILTER_OK) {
      // Drop the connection due to an error
      go_parser_->Close();
      return Network::FilterStatus::StopIteration;
    }

    // XXX: Unfortunately continueReading() continues from the next filter, and there seems to be no way
    // to trigger the whole filter chain to be called.

    go_parser_->SetReplyEndStream(end_stream);
  }

  return Network::FilterStatus::Continue;
}

void Instance::onEvent(Network::ConnectionEvent event) {
  if (event == Network::ConnectionEvent::RemoteClose ||
      event == Network::ConnectionEvent::LocalClose) {
    if (maps_) {
      auto& conn = callbacks_->connection();
      bool ok = maps_->removeBpfMetadata(conn, proxy_port_);
      ENVOY_CONN_LOG(debug, "Cilium Network: Connection Closed, proxymap cleanup {}", conn,
		     ok ? "succeeded" : "failed");
    }
  }
}

} // namespace CiliumL3
} // namespace Filter
} // namespace Envoy
