// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/client_ssl_auth/v2/client_ssl_auth.proto

package envoy_config_filter_network_client_ssl_auth_v2

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type ClientSSLAuth struct {
	// The :ref:`cluster manager <arch_overview_cluster_manager>` cluster that runs
	// the authentication service. The filter will connect to the service every 60s to fetch the list
	// of principals. The service must support the expected :ref:`REST API
	// <config_network_filters_client_ssl_auth_rest_api>`.
	AuthApiCluster string `protobuf:"bytes,1,opt,name=auth_api_cluster,json=authApiCluster,proto3" json:"auth_api_cluster,omitempty"`
	// The prefix to use when emitting :ref:`statistics
	// <config_network_filters_client_ssl_auth_stats>`.
	StatPrefix string `protobuf:"bytes,2,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Time in milliseconds between principal refreshes from the
	// authentication service. Default is 60000 (60s). The actual fetch time
	// will be this value plus a random jittered value between
	// 0-refresh_delay_ms milliseconds.
	RefreshDelay *duration.Duration `protobuf:"bytes,3,opt,name=refresh_delay,json=refreshDelay,proto3" json:"refresh_delay,omitempty"`
	// An optional list of IP address and subnet masks that should be white
	// listed for access by the filter. If no list is provided, there is no
	// IP white list.
	IpWhiteList          []*core.CidrRange `protobuf:"bytes,4,rep,name=ip_white_list,json=ipWhiteList,proto3" json:"ip_white_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ClientSSLAuth) Reset()         { *m = ClientSSLAuth{} }
func (m *ClientSSLAuth) String() string { return proto.CompactTextString(m) }
func (*ClientSSLAuth) ProtoMessage()    {}
func (*ClientSSLAuth) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c05e9c9b57da130, []int{0}
}

func (m *ClientSSLAuth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientSSLAuth.Unmarshal(m, b)
}
func (m *ClientSSLAuth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientSSLAuth.Marshal(b, m, deterministic)
}
func (m *ClientSSLAuth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientSSLAuth.Merge(m, src)
}
func (m *ClientSSLAuth) XXX_Size() int {
	return xxx_messageInfo_ClientSSLAuth.Size(m)
}
func (m *ClientSSLAuth) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientSSLAuth.DiscardUnknown(m)
}

var xxx_messageInfo_ClientSSLAuth proto.InternalMessageInfo

func (m *ClientSSLAuth) GetAuthApiCluster() string {
	if m != nil {
		return m.AuthApiCluster
	}
	return ""
}

func (m *ClientSSLAuth) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ClientSSLAuth) GetRefreshDelay() *duration.Duration {
	if m != nil {
		return m.RefreshDelay
	}
	return nil
}

func (m *ClientSSLAuth) GetIpWhiteList() []*core.CidrRange {
	if m != nil {
		return m.IpWhiteList
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientSSLAuth)(nil), "envoy.config.filter.network.client_ssl_auth.v2.ClientSSLAuth")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/client_ssl_auth/v2/client_ssl_auth.proto", fileDescriptor_2c05e9c9b57da130)
}

var fileDescriptor_2c05e9c9b57da130 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x3d, 0x8f, 0x13, 0x31,
	0x10, 0xd5, 0xe6, 0x8e, 0x43, 0x38, 0x04, 0x9d, 0xb6, 0x21, 0x9c, 0xf8, 0x88, 0xa8, 0x52, 0xd9,
	0x62, 0x23, 0x3a, 0x84, 0xb8, 0x24, 0xe5, 0x15, 0x51, 0xae, 0xa0, 0xb4, 0x7c, 0xd9, 0xd9, 0xcd,
	0x08, 0x63, 0x5b, 0xf6, 0xec, 0x5e, 0xd2, 0xf1, 0x0f, 0x28, 0xe1, 0xb7, 0xf0, 0x0b, 0x68, 0xf9,
	0x23, 0x14, 0x94, 0x14, 0x08, 0x79, 0xbd, 0x29, 0x08, 0xa2, 0xa0, 0xdb, 0x9d, 0xf7, 0xe6, 0xcd,
	0xbc, 0x37, 0x66, 0x4b, 0x30, 0xad, 0xdd, 0x8b, 0x8d, 0x35, 0x15, 0xd6, 0xa2, 0x42, 0x4d, 0xe0,
	0x85, 0x01, 0xba, 0xb5, 0xfe, 0x9d, 0xd8, 0x68, 0x04, 0x43, 0x32, 0x04, 0x2d, 0x55, 0x43, 0x5b,
	0xd1, 0x16, 0xc7, 0x25, 0xee, 0xbc, 0x25, 0x9b, 0xf3, 0x4e, 0x85, 0x27, 0x15, 0x9e, 0x54, 0x78,
	0xaf, 0xc2, 0x8f, 0x5b, 0xda, 0xe2, 0xe2, 0x59, 0x9a, 0xaa, 0x1c, 0x76, 0x9a, 0xd6, 0x83, 0x50,
	0x65, 0xe9, 0x21, 0x84, 0x24, 0x78, 0xf1, 0xb4, 0xb6, 0xb6, 0xd6, 0x20, 0xba, 0xbf, 0x9b, 0xa6,
	0x12, 0x65, 0xe3, 0x15, 0xa1, 0x35, 0x07, 0xbc, 0x29, 0x9d, 0x12, 0xca, 0x18, 0x4b, 0x5d, 0x39,
	0x88, 0xf7, 0x58, 0x7b, 0x45, 0xd0, 0xe3, 0x4f, 0xfe, 0xc2, 0x03, 0x29, 0x6a, 0x0e, 0xf2, 0x0f,
	0x5b, 0xa5, 0xb1, 0x54, 0x04, 0xe2, 0xf0, 0x91, 0x80, 0xe7, 0xdf, 0x33, 0x36, 0x5a, 0x74, 0xfb,
	0x5e, 0x5f, 0x5f, 0x5d, 0x36, 0xb4, 0xcd, 0x5f, 0xb0, 0xf3, 0xb8, 0xb5, 0x54, 0x0e, 0xe5, 0x46,
	0x37, 0x81, 0xc0, 0x8f, 0xb3, 0x49, 0x36, 0xbd, 0x37, 0xbf, 0xfb, 0x73, 0x7e, 0xea, 0x07, 0x93,
	0x6c, 0xfd, 0x20, 0x12, 0x2e, 0x1d, 0x2e, 0x12, 0x9c, 0x4f, 0xd9, 0x30, 0x4e, 0x93, 0xce, 0x43,
	0x85, 0xbb, 0xf1, 0xe0, 0x4f, 0x36, 0x8b, 0xd8, 0xaa, 0x83, 0xf2, 0xd7, 0x6c, 0xe4, 0xa1, 0xf2,
	0x10, 0xb6, 0xb2, 0x04, 0xad, 0xf6, 0xe3, 0x93, 0x49, 0x36, 0x1d, 0x16, 0x8f, 0x78, 0xb2, 0xcf,
	0x0f, 0xf6, 0xf9, 0xb2, 0xb7, 0xbf, 0xbe, 0xdf, 0xf3, 0x97, 0x91, 0x9e, 0xbf, 0x61, 0x23, 0x74,
	0xf2, 0x76, 0x8b, 0x04, 0x52, 0x63, 0xa0, 0xf1, 0xe9, 0xe4, 0x64, 0x3a, 0x2c, 0x1e, 0xf7, 0xf7,
	0x50, 0x0e, 0x79, 0x5b, 0xf0, 0x98, 0x2f, 0x5f, 0x60, 0xe9, 0xd7, 0xca, 0xd4, 0xb0, 0x1e, 0xa2,
	0x7b, 0x1b, 0x3b, 0xae, 0x30, 0xd0, 0xfc, 0x53, 0xf6, 0xe3, 0xf3, 0xaf, 0x8f, 0x77, 0x5e, 0xe6,
	0xb3, 0xd4, 0x02, 0x3b, 0x02, 0x13, 0x62, 0x62, 0xfd, 0x19, 0xc3, 0xbf, 0xef, 0x38, 0xfb, 0xf2,
	0xe1, 0xeb, 0xb7, 0xb3, 0xc1, 0x79, 0xc6, 0x5e, 0xa1, 0x4d, 0x23, 0x9d, 0xb7, 0xbb, 0xfd, 0x7f,
	0xbe, 0x86, 0x79, 0xde, 0x27, 0x1e, 0x74, 0x4c, 0x7c, 0x15, 0x3d, 0xaf, 0xb2, 0x9b, 0xb3, 0xce,
	0xfc, 0xec, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0x8c, 0xd0, 0x34, 0xa2, 0x02, 0x00, 0x00,
}
