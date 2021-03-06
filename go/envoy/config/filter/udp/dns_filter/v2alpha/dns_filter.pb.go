// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/udp/dns_filter/v2alpha/dns_filter.proto

package envoy_config_filter_udp_dns_filter_v2alpha

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	v2alpha "github.com/cilium/proxy/go/envoy/data/dns/v2alpha"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Configuration for the DNS filter.
type DnsFilterConfig struct {
	// The stat prefix used when emitting DNS filter statistics
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Server context configuration
	ServerConfig         *DnsFilterConfig_ServerContextConfig `protobuf:"bytes,2,opt,name=server_config,json=serverConfig,proto3" json:"server_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *DnsFilterConfig) Reset()         { *m = DnsFilterConfig{} }
func (m *DnsFilterConfig) String() string { return proto.CompactTextString(m) }
func (*DnsFilterConfig) ProtoMessage()    {}
func (*DnsFilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5252ffc76176bc87, []int{0}
}

func (m *DnsFilterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsFilterConfig.Unmarshal(m, b)
}
func (m *DnsFilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsFilterConfig.Marshal(b, m, deterministic)
}
func (m *DnsFilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsFilterConfig.Merge(m, src)
}
func (m *DnsFilterConfig) XXX_Size() int {
	return xxx_messageInfo_DnsFilterConfig.Size(m)
}
func (m *DnsFilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsFilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DnsFilterConfig proto.InternalMessageInfo

func (m *DnsFilterConfig) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *DnsFilterConfig) GetServerConfig() *DnsFilterConfig_ServerContextConfig {
	if m != nil {
		return m.ServerConfig
	}
	return nil
}

// This message contains the configuration for the Dns Filter operating
// in a server context. This message will contain the virtual hosts and
// associated addresses with which Envoy will respond to queries
type DnsFilterConfig_ServerContextConfig struct {
	// Types that are valid to be assigned to ConfigSource:
	//	*DnsFilterConfig_ServerContextConfig_InlineDnsTable
	//	*DnsFilterConfig_ServerContextConfig_ExternalDnsTable
	ConfigSource         isDnsFilterConfig_ServerContextConfig_ConfigSource `protobuf_oneof:"config_source"`
	XXX_NoUnkeyedLiteral struct{}                                           `json:"-"`
	XXX_unrecognized     []byte                                             `json:"-"`
	XXX_sizecache        int32                                              `json:"-"`
}

func (m *DnsFilterConfig_ServerContextConfig) Reset()         { *m = DnsFilterConfig_ServerContextConfig{} }
func (m *DnsFilterConfig_ServerContextConfig) String() string { return proto.CompactTextString(m) }
func (*DnsFilterConfig_ServerContextConfig) ProtoMessage()    {}
func (*DnsFilterConfig_ServerContextConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5252ffc76176bc87, []int{0, 0}
}

func (m *DnsFilterConfig_ServerContextConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsFilterConfig_ServerContextConfig.Unmarshal(m, b)
}
func (m *DnsFilterConfig_ServerContextConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsFilterConfig_ServerContextConfig.Marshal(b, m, deterministic)
}
func (m *DnsFilterConfig_ServerContextConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsFilterConfig_ServerContextConfig.Merge(m, src)
}
func (m *DnsFilterConfig_ServerContextConfig) XXX_Size() int {
	return xxx_messageInfo_DnsFilterConfig_ServerContextConfig.Size(m)
}
func (m *DnsFilterConfig_ServerContextConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsFilterConfig_ServerContextConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DnsFilterConfig_ServerContextConfig proto.InternalMessageInfo

type isDnsFilterConfig_ServerContextConfig_ConfigSource interface {
	isDnsFilterConfig_ServerContextConfig_ConfigSource()
}

type DnsFilterConfig_ServerContextConfig_InlineDnsTable struct {
	InlineDnsTable *v2alpha.DnsTable `protobuf:"bytes,1,opt,name=inline_dns_table,json=inlineDnsTable,proto3,oneof"`
}

type DnsFilterConfig_ServerContextConfig_ExternalDnsTable struct {
	ExternalDnsTable *core.DataSource `protobuf:"bytes,2,opt,name=external_dns_table,json=externalDnsTable,proto3,oneof"`
}

func (*DnsFilterConfig_ServerContextConfig_InlineDnsTable) isDnsFilterConfig_ServerContextConfig_ConfigSource() {
}

func (*DnsFilterConfig_ServerContextConfig_ExternalDnsTable) isDnsFilterConfig_ServerContextConfig_ConfigSource() {
}

func (m *DnsFilterConfig_ServerContextConfig) GetConfigSource() isDnsFilterConfig_ServerContextConfig_ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}

func (m *DnsFilterConfig_ServerContextConfig) GetInlineDnsTable() *v2alpha.DnsTable {
	if x, ok := m.GetConfigSource().(*DnsFilterConfig_ServerContextConfig_InlineDnsTable); ok {
		return x.InlineDnsTable
	}
	return nil
}

func (m *DnsFilterConfig_ServerContextConfig) GetExternalDnsTable() *core.DataSource {
	if x, ok := m.GetConfigSource().(*DnsFilterConfig_ServerContextConfig_ExternalDnsTable); ok {
		return x.ExternalDnsTable
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DnsFilterConfig_ServerContextConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DnsFilterConfig_ServerContextConfig_InlineDnsTable)(nil),
		(*DnsFilterConfig_ServerContextConfig_ExternalDnsTable)(nil),
	}
}

func init() {
	proto.RegisterType((*DnsFilterConfig)(nil), "envoy.config.filter.udp.dns_filter.v2alpha.DnsFilterConfig")
	proto.RegisterType((*DnsFilterConfig_ServerContextConfig)(nil), "envoy.config.filter.udp.dns_filter.v2alpha.DnsFilterConfig.ServerContextConfig")
}

func init() {
	proto.RegisterFile("envoy/config/filter/udp/dns_filter/v2alpha/dns_filter.proto", fileDescriptor_5252ffc76176bc87)
}

var fileDescriptor_5252ffc76176bc87 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcf, 0x8b, 0xd4, 0x30,
	0x14, 0x36, 0x55, 0x57, 0xcd, 0xb8, 0x6b, 0x89, 0x82, 0xcb, 0xe0, 0xca, 0xe0, 0x41, 0x06, 0x0f,
	0x89, 0x74, 0x2f, 0x82, 0xb7, 0xee, 0x22, 0x1e, 0x14, 0x87, 0x59, 0xef, 0xe5, 0xcd, 0x34, 0x1d,
	0x03, 0x35, 0x09, 0x49, 0x5a, 0xba, 0x37, 0xff, 0x03, 0xc1, 0x93, 0x7f, 0x8b, 0x7f, 0xc1, 0x5c,
	0xfd, 0x57, 0x3c, 0x89, 0x88, 0x48, 0x92, 0x76, 0xa8, 0xbf, 0xc0, 0x3d, 0xb5, 0x79, 0xdf, 0xfb,
	0xbe, 0xf7, 0xe5, 0x7b, 0xc1, 0x4f, 0xb9, 0x6c, 0xd5, 0x39, 0x5b, 0x2b, 0x59, 0x89, 0x0d, 0xab,
	0x44, 0xed, 0xb8, 0x61, 0x4d, 0xa9, 0x59, 0x29, 0x6d, 0xd1, 0x1f, 0xdb, 0x0c, 0x6a, 0xfd, 0x06,
	0x46, 0x25, 0xaa, 0x8d, 0x72, 0x8a, 0x3c, 0x0a, 0x64, 0x1a, 0xc9, 0xb4, 0x87, 0x9a, 0x52, 0xd3,
	0x51, 0x67, 0x4f, 0x9e, 0xde, 0x8b, 0x83, 0x40, 0x0b, 0xd6, 0x66, 0x6c, 0xad, 0x0c, 0x67, 0x2b,
	0xb0, 0x3c, 0x2a, 0x4d, 0x1f, 0x46, 0xb4, 0x04, 0x17, 0xc6, 0xfc, 0x32, 0xd2, 0xc1, 0xaa, 0x1e,
	0xfa, 0xee, 0x37, 0xa5, 0x06, 0x06, 0x52, 0x2a, 0x07, 0x4e, 0x28, 0x69, 0xd9, 0x5b, 0xb1, 0x31,
	0xe0, 0x06, 0xfc, 0xe8, 0x0f, 0xdc, 0x3a, 0x70, 0x8d, 0xed, 0xe1, 0xbb, 0x2d, 0xd4, 0xa2, 0x04,
	0xc7, 0xd9, 0xf0, 0x13, 0x81, 0x07, 0xdf, 0x13, 0x7c, 0xeb, 0x54, 0xda, 0x67, 0xc1, 0xf3, 0x49,
	0xb8, 0x0f, 0x99, 0xe3, 0x89, 0x27, 0x17, 0xda, 0xf0, 0x4a, 0x74, 0x87, 0x68, 0x86, 0xe6, 0x37,
	0xf2, 0x6b, 0xdf, 0xf2, 0x2b, 0x26, 0x49, 0xd1, 0x12, 0x7b, 0x6c, 0x11, 0x20, 0xe2, 0xf0, 0xbe,
	0xe5, 0xa6, 0xe5, 0xa6, 0x88, 0x51, 0x1c, 0x26, 0x33, 0x34, 0x9f, 0x64, 0xaf, 0xe8, 0xff, 0xe7,
	0x43, 0x7f, 0x9b, 0x4e, 0xcf, 0x82, 0xe0, 0x89, 0x92, 0x8e, 0x77, 0x2e, 0xd6, 0x96, 0x37, 0xed,
	0x50, 0xac, 0xc4, 0x66, 0xba, 0x45, 0xf8, 0xf6, 0x5f, 0xba, 0xc8, 0x0b, 0x9c, 0x0a, 0x59, 0x0b,
	0xc9, 0x8b, 0x5d, 0x7a, 0xc1, 0xfc, 0x24, 0x9b, 0xf5, 0x86, 0x7c, 0xcc, 0xde, 0xc3, 0x78, 0xf8,
	0x6b, 0xdf, 0xf7, 0xfc, 0xd2, 0xf2, 0x20, 0x72, 0x87, 0x0a, 0x79, 0x89, 0x09, 0xef, 0x1c, 0x37,
	0x12, 0xea, 0x91, 0x5e, 0xbc, 0xe0, 0x51, 0xaf, 0x07, 0x5a, 0xd0, 0x36, 0xa3, 0x7e, 0xa9, 0xf4,
	0x14, 0x1c, 0x9c, 0xa9, 0xc6, 0xac, 0xbd, 0x58, 0x3a, 0x50, 0x07, 0xb9, 0xfc, 0x0e, 0xde, 0x8f,
	0x71, 0x14, 0x36, 0x34, 0x91, 0xcb, 0x5f, 0x73, 0x94, 0x7f, 0x40, 0x5f, 0x3e, 0xfe, 0x78, 0x7f,
	0xf5, 0x31, 0xe9, 0x05, 0x3d, 0x4d, 0x5a, 0xbf, 0xbf, 0x7f, 0xa5, 0x76, 0x1c, 0x8c, 0x7f, 0x7a,
	0xb7, 0xfd, 0xbc, 0x97, 0x5c, 0x47, 0xf1, 0x9b, 0x22, 0xfc, 0x44, 0xa8, 0x28, 0xa1, 0x8d, 0xea,
	0xce, 0x2f, 0x90, 0x7f, 0x7e, 0xb0, 0x5b, 0xc0, 0xc2, 0xbf, 0x88, 0x05, 0x5a, 0xed, 0x85, 0xa7,
	0x71, 0xfc, 0x33, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x2b, 0xdf, 0x37, 0x23, 0x03, 0x00, 0x00,
}
