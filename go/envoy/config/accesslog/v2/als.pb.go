// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/accesslog/v2/als.proto

package v2

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
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

// Configuration for the built-in *envoy.http_grpc_access_log*
// :ref:`AccessLog <envoy_api_msg_config.filter.accesslog.v2.AccessLog>`. This configuration will
// populate :ref:`StreamAccessLogsMessage.http_logs
// <envoy_api_field_service.accesslog.v2.StreamAccessLogsMessage.http_logs>`.
type HttpGrpcAccessLogConfig struct {
	CommonConfig *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	// Additional request headers to log in :ref:`HTTPRequestProperties.request_headers
	// <envoy_api_field_data.accesslog.v2.HTTPRequestProperties.request_headers>`.
	AdditionalRequestHeadersToLog []string `protobuf:"bytes,2,rep,name=additional_request_headers_to_log,json=additionalRequestHeadersToLog,proto3" json:"additional_request_headers_to_log,omitempty"`
	// Additional response headers to log in :ref:`HTTPResponseProperties.response_headers
	// <envoy_api_field_data.accesslog.v2.HTTPResponseProperties.response_headers>`.
	AdditionalResponseHeadersToLog []string `protobuf:"bytes,3,rep,name=additional_response_headers_to_log,json=additionalResponseHeadersToLog,proto3" json:"additional_response_headers_to_log,omitempty"`
	// Additional response trailers to log in :ref:`HTTPResponseProperties.response_trailers
	// <envoy_api_field_data.accesslog.v2.HTTPResponseProperties.response_trailers>`.
	AdditionalResponseTrailersToLog []string `protobuf:"bytes,4,rep,name=additional_response_trailers_to_log,json=additionalResponseTrailersToLog,proto3" json:"additional_response_trailers_to_log,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *HttpGrpcAccessLogConfig) Reset()         { *m = HttpGrpcAccessLogConfig{} }
func (m *HttpGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*HttpGrpcAccessLogConfig) ProtoMessage()    {}
func (*HttpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7b431652a309a2e, []int{0}
}

func (m *HttpGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *HttpGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (m *HttpGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpGrpcAccessLogConfig.Merge(m, src)
}
func (m *HttpGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Size(m)
}
func (m *HttpGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HttpGrpcAccessLogConfig proto.InternalMessageInfo

func (m *HttpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalRequestHeadersToLog() []string {
	if m != nil {
		return m.AdditionalRequestHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseHeadersToLog() []string {
	if m != nil {
		return m.AdditionalResponseHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseTrailersToLog() []string {
	if m != nil {
		return m.AdditionalResponseTrailersToLog
	}
	return nil
}

// Configuration for the built-in *envoy.tcp_grpc_access_log* type. This configuration will
// populate *StreamAccessLogsMessage.tcp_logs*.
// [#not-implemented-hide:]
type TcpGrpcAccessLogConfig struct {
	CommonConfig         *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *TcpGrpcAccessLogConfig) Reset()         { *m = TcpGrpcAccessLogConfig{} }
func (m *TcpGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*TcpGrpcAccessLogConfig) ProtoMessage()    {}
func (*TcpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7b431652a309a2e, []int{1}
}

func (m *TcpGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *TcpGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (m *TcpGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpGrpcAccessLogConfig.Merge(m, src)
}
func (m *TcpGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Size(m)
}
func (m *TcpGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TcpGrpcAccessLogConfig proto.InternalMessageInfo

func (m *TcpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

// Common configuration for gRPC access logs.
type CommonGrpcAccessLogConfig struct {
	// The friendly name of the access log to be returned in :ref:`StreamAccessLogsMessage.Identifier
	// <envoy_api_msg_service.accesslog.v2.StreamAccessLogsMessage.Identifier>`. This allows the
	// access log server to differentiate between different access logs coming from the same Envoy.
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	// The gRPC service for the access log service.
	GrpcService          *core.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CommonGrpcAccessLogConfig) Reset()         { *m = CommonGrpcAccessLogConfig{} }
func (m *CommonGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*CommonGrpcAccessLogConfig) ProtoMessage()    {}
func (*CommonGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7b431652a309a2e, []int{2}
}

func (m *CommonGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *CommonGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (m *CommonGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonGrpcAccessLogConfig.Merge(m, src)
}
func (m *CommonGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Size(m)
}
func (m *CommonGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonGrpcAccessLogConfig proto.InternalMessageInfo

func (m *CommonGrpcAccessLogConfig) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *CommonGrpcAccessLogConfig) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterType((*HttpGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v2.HttpGrpcAccessLogConfig")
	proto.RegisterType((*TcpGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v2.TcpGrpcAccessLogConfig")
	proto.RegisterType((*CommonGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v2.CommonGrpcAccessLogConfig")
}

func init() {
	proto.RegisterFile("envoy/config/accesslog/v2/als.proto", fileDescriptor_e7b431652a309a2e)
}

var fileDescriptor_e7b431652a309a2e = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0xbf, 0xae, 0xd3, 0x30,
	0x14, 0xc6, 0x95, 0xb4, 0x40, 0xeb, 0x16, 0x09, 0x65, 0xa0, 0x7f, 0x24, 0x4a, 0x49, 0x2b, 0xd1,
	0x29, 0x91, 0x02, 0x0b, 0x63, 0xdb, 0x81, 0x0a, 0x15, 0x54, 0x85, 0x4e, 0x2c, 0x91, 0x71, 0x8c,
	0xb1, 0xe4, 0xe4, 0x18, 0xdb, 0x44, 0x74, 0x62, 0x67, 0xe2, 0x79, 0x98, 0x78, 0x18, 0x16, 0xde,
	0x02, 0xc5, 0x2e, 0xf7, 0xa6, 0xf7, 0xb6, 0xf3, 0xdd, 0x22, 0x9d, 0xdf, 0xf9, 0x7d, 0x5f, 0xac,
	0x83, 0x66, 0xb4, 0xac, 0xe0, 0x10, 0x13, 0x28, 0x3f, 0x71, 0x16, 0x63, 0x42, 0xa8, 0xd6, 0x02,
	0x58, 0x5c, 0x25, 0x31, 0x16, 0x3a, 0x92, 0x0a, 0x0c, 0x04, 0x23, 0x0b, 0x45, 0x0e, 0x8a, 0xae,
	0xa0, 0xa8, 0x4a, 0xc6, 0x73, 0xb7, 0x8f, 0x25, 0xaf, 0x57, 0x08, 0x28, 0x1a, 0x33, 0x25, 0x49,
	0xa6, 0xa9, 0xaa, 0x38, 0xa1, 0x4e, 0x30, 0x1e, 0x54, 0x58, 0xf0, 0x1c, 0x1b, 0x1a, 0xff, 0xff,
	0x70, 0x83, 0xf0, 0x8f, 0x8f, 0x06, 0x1b, 0x63, 0xe4, 0x6b, 0x25, 0xc9, 0xd2, 0x7a, 0xb7, 0xc0,
	0xd6, 0x36, 0x27, 0xa0, 0xe8, 0x21, 0x81, 0xa2, 0x80, 0x32, 0x73, 0xc1, 0x43, 0x6f, 0xea, 0x2d,
	0x7a, 0xc9, 0xcb, 0xe8, 0x62, 0x9b, 0x68, 0x6d, 0xf9, 0x33, 0xb2, 0x15, 0xfa, 0xf5, 0xf7, 0x77,
	0xeb, 0xde, 0x0f, 0xcf, 0x7f, 0xe4, 0xa5, 0x7d, 0xa7, 0x3d, 0xc6, 0x6c, 0xd0, 0x33, 0x9c, 0xe7,
	0xdc, 0x70, 0x28, 0xb1, 0xc8, 0x14, 0xfd, 0xf2, 0x95, 0x6a, 0x93, 0x7d, 0xa6, 0x38, 0xa7, 0x4a,
	0x67, 0x06, 0x32, 0x01, 0x6c, 0xe8, 0x4f, 0x5b, 0x8b, 0x6e, 0xfa, 0xe4, 0x1a, 0x4c, 0x1d, 0xb7,
	0x71, 0xd8, 0x1e, 0xb6, 0xc0, 0x82, 0x37, 0x28, 0x3c, 0x31, 0x69, 0x09, 0xa5, 0xa6, 0x37, 0x55,
	0x2d, 0xab, 0x9a, 0x34, 0x55, 0x0e, 0x3c, 0x71, 0x6d, 0xd1, 0xec, 0x9c, 0xcb, 0x28, 0xcc, 0x45,
	0x43, 0xd6, 0xb6, 0xb2, 0xa7, 0xb7, 0x65, 0xfb, 0x23, 0x68, 0x6d, 0xe1, 0x77, 0xf4, 0x78, 0x4f,
	0xee, 0xf0, 0x91, 0xc3, 0x9f, 0x1e, 0x1a, 0x5d, 0xdc, 0x0b, 0xe6, 0xa8, 0x23, 0x80, 0x65, 0x25,
	0x2e, 0xa8, 0xcd, 0xef, 0xae, 0xba, 0xb5, 0xa9, 0xad, 0xfc, 0xa9, 0x97, 0x3e, 0x10, 0xc0, 0xde,
	0xe1, 0x82, 0x06, 0x6f, 0x51, 0xbf, 0x79, 0x5a, 0x43, 0xdf, 0x36, 0x9d, 0x1c, 0x9b, 0x62, 0xc9,
	0xeb, 0x72, 0xf5, 0x05, 0x46, 0x75, 0xc6, 0x7b, 0x47, 0x9d, 0x74, 0xea, 0xb1, 0xc6, 0xe0, 0x15,
	0x7a, 0xce, 0xc1, 0x2d, 0x4b, 0x05, 0xdf, 0x0e, 0x97, 0xff, 0x78, 0xd5, 0x59, 0x0a, 0xbd, 0xab,
	0xef, 0x75, 0xe7, 0x7d, 0xf0, 0xab, 0xe4, 0xe3, 0x7d, 0x7b, 0xbc, 0x2f, 0xfe, 0x05, 0x00, 0x00,
	0xff, 0xff, 0xfc, 0xec, 0x6a, 0x4f, 0x3d, 0x03, 0x00, 0x00,
}
