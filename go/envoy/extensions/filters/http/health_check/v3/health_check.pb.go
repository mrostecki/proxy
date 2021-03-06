// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/health_check/v3/health_check.proto

package envoy_extensions_filters_http_health_check_v3

import (
	fmt "fmt"
	v31 "github.com/cilium/proxy/go/envoy/config/route/v3"
	v3 "github.com/cilium/proxy/go/envoy/type/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// [#next-free-field: 6]
type HealthCheck struct {
	// Specifies whether the filter operates in pass through mode or not.
	PassThroughMode *wrappers.BoolValue `protobuf:"bytes,1,opt,name=pass_through_mode,json=passThroughMode,proto3" json:"pass_through_mode,omitempty"`
	// If operating in pass through mode, the amount of time in milliseconds
	// that the filter should cache the upstream response.
	CacheTime *duration.Duration `protobuf:"bytes,3,opt,name=cache_time,json=cacheTime,proto3" json:"cache_time,omitempty"`
	// If operating in non-pass-through mode, specifies a set of upstream cluster
	// names and the minimum percentage of servers in each of those clusters that
	// must be healthy or degraded in order for the filter to return a 200.
	ClusterMinHealthyPercentages map[string]*v3.Percent `protobuf:"bytes,4,rep,name=cluster_min_healthy_percentages,json=clusterMinHealthyPercentages,proto3" json:"cluster_min_healthy_percentages,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Specifies a set of health check request headers to match on. The health check filter will
	// check a request’s headers against all the specified headers. To specify the health check
	// endpoint, set the ``:path`` header to match on.
	Headers              []*v31.HeaderMatcher `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_554d4249f82380e7, []int{0}
}

func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetPassThroughMode() *wrappers.BoolValue {
	if m != nil {
		return m.PassThroughMode
	}
	return nil
}

func (m *HealthCheck) GetCacheTime() *duration.Duration {
	if m != nil {
		return m.CacheTime
	}
	return nil
}

func (m *HealthCheck) GetClusterMinHealthyPercentages() map[string]*v3.Percent {
	if m != nil {
		return m.ClusterMinHealthyPercentages
	}
	return nil
}

func (m *HealthCheck) GetHeaders() []*v31.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthCheck)(nil), "envoy.extensions.filters.http.health_check.v3.HealthCheck")
	proto.RegisterMapType((map[string]*v3.Percent)(nil), "envoy.extensions.filters.http.health_check.v3.HealthCheck.ClusterMinHealthyPercentagesEntry")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/health_check/v3/health_check.proto", fileDescriptor_554d4249f82380e7)
}

var fileDescriptor_554d4249f82380e7 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0x26, 0xd9, 0xdd, 0xda, 0xce, 0x1e, 0x5c, 0x73, 0xd0, 0x75, 0xd5, 0xda, 0x8a, 0x87, 0x1e,
	0xea, 0x04, 0x76, 0x45, 0x4a, 0x0b, 0x22, 0x5b, 0x05, 0x11, 0x16, 0x96, 0x50, 0x04, 0x41, 0x08,
	0xd3, 0xe4, 0x6d, 0x32, 0x34, 0x99, 0x09, 0x33, 0x93, 0xd8, 0xdc, 0x3c, 0x8a, 0x7f, 0x82, 0x57,
	0x8f, 0xfe, 0x07, 0xde, 0x05, 0xaf, 0xfe, 0x3b, 0x9e, 0x64, 0x7e, 0x2c, 0xdd, 0xa5, 0x6a, 0xe9,
	0x6d, 0x92, 0xf7, 0xbd, 0xef, 0xfb, 0xde, 0xf7, 0x1e, 0x7a, 0x01, 0xac, 0xe1, 0x6d, 0x08, 0xe7,
	0x0a, 0x98, 0xa4, 0x9c, 0xc9, 0x70, 0x41, 0x0b, 0x05, 0x42, 0x86, 0xb9, 0x52, 0x55, 0x98, 0x03,
	0x29, 0x54, 0x1e, 0x27, 0x39, 0x24, 0x67, 0x61, 0x33, 0x59, 0xfb, 0xc6, 0x95, 0xe0, 0x8a, 0x07,
	0x4f, 0x0c, 0x03, 0xbe, 0x60, 0xc0, 0x8e, 0x01, 0x6b, 0x06, 0xbc, 0xd6, 0xd1, 0x4c, 0x46, 0xfb,
	0x56, 0x30, 0xe1, 0x6c, 0x41, 0xb3, 0x50, 0xf0, 0x5a, 0x81, 0x26, 0x36, 0x8f, 0x38, 0xe1, 0x65,
	0xc5, 0x19, 0x30, 0x25, 0x2d, 0xf9, 0xe8, 0x9e, 0x45, 0xab, 0xb6, 0x32, 0xa8, 0x0a, 0x44, 0x02,
	0x4c, 0xb9, 0xe2, 0x76, 0xc6, 0x79, 0x56, 0x40, 0x68, 0xbe, 0x4e, 0xeb, 0x45, 0x98, 0xd6, 0x82,
	0x28, 0xca, 0xd9, 0xbf, 0xea, 0x1f, 0x04, 0xa9, 0x2a, 0xed, 0xcc, 0xd6, 0x1f, 0xd4, 0x69, 0x45,
	0x42, 0xc2, 0x18, 0x57, 0xa6, 0x4d, 0x86, 0x52, 0x11, 0x55, 0x2f, 0xcb, 0xbb, 0x97, 0xca, 0x0d,
	0x08, 0x3d, 0x21, 0x65, 0x99, 0x83, 0xdc, 0x69, 0x48, 0x41, 0x53, 0xa2, 0x47, 0x70, 0x0f, 0x5b,
	0x78, 0xf4, 0xad, 0x8b, 0xfa, 0xaf, 0xcd, 0xe4, 0xc7, 0x7a, 0xf0, 0x60, 0x8e, 0x6e, 0x55, 0x44,
	0xca, 0x58, 0xe5, 0x82, 0xd7, 0x59, 0x1e, 0x97, 0x3c, 0x85, 0xa1, 0xb7, 0xe3, 0xed, 0xf5, 0xc7,
	0x23, 0x6c, 0x6d, 0xe2, 0xa5, 0x4d, 0x3c, 0xe5, 0xbc, 0x78, 0x4b, 0x8a, 0x1a, 0xa6, 0x9b, 0xbf,
	0xa7, 0xbd, 0xcf, 0x9e, 0x3f, 0xf0, 0xa2, 0x9b, 0xba, 0xfd, 0xc4, 0x76, 0xcf, 0x78, 0x0a, 0xc1,
	0x01, 0x42, 0x09, 0x49, 0x72, 0x88, 0x15, 0x2d, 0x61, 0xd8, 0x31, 0x54, 0x77, 0x2f, 0x51, 0xbd,
	0x74, 0x89, 0x44, 0x5b, 0x06, 0x7c, 0x42, 0x4b, 0x08, 0xbe, 0x7a, 0xe8, 0x61, 0x52, 0xd4, 0x52,
	0x81, 0x88, 0x4b, 0xca, 0x62, 0xbb, 0xa1, 0x36, 0x76, 0xe1, 0x92, 0x0c, 0xe4, 0xb0, 0xbb, 0xd3,
	0xd9, 0xeb, 0x8f, 0xdf, 0xe3, 0x6b, 0xed, 0x16, 0xaf, 0x4c, 0x8c, 0x8f, 0xad, 0xc2, 0x8c, 0x32,
	0xfb, 0xb7, 0x9d, 0x5f, 0xd0, 0xbf, 0x62, 0x4a, 0xb4, 0xd1, 0xfd, 0xe4, 0x3f, 0x90, 0xe0, 0x39,
	0xba, 0x91, 0x03, 0x49, 0x41, 0xc8, 0x61, 0xcf, 0x98, 0x79, 0xec, 0xcc, 0xd8, 0xcb, 0xc1, 0xe6,
	0x60, 0x9c, 0x68, 0x0a, 0x62, 0x46, 0x54, 0x92, 0x83, 0x88, 0x96, 0x4d, 0xa3, 0x0c, 0xed, 0x5e,
	0x69, 0x21, 0x18, 0xa0, 0xce, 0x19, 0xb4, 0x66, 0x11, 0x5b, 0x91, 0x7e, 0x06, 0xfb, 0xa8, 0xd7,
	0xe8, 0xe8, 0x87, 0xbe, 0x49, 0xf4, 0xb6, 0x13, 0xd5, 0x07, 0xa8, 0xc5, 0x1c, 0x43, 0x64, 0x41,
	0x87, 0xfe, 0x81, 0x77, 0x78, 0xf4, 0xe5, 0xc7, 0xa7, 0xed, 0x67, 0xe8, 0xe9, 0x9a, 0x3b, 0x1b,
	0xd3, 0xdf, 0x52, 0x1a, 0xaf, 0xa6, 0xf4, 0xa6, 0xbb, 0xe9, 0x0f, 0x3a, 0xd3, 0x77, 0xdf, 0x3f,
	0xfe, 0xfc, 0xb5, 0xe1, 0x0f, 0x7c, 0x74, 0x44, 0xb9, 0x55, 0xac, 0x04, 0x3f, 0x6f, 0xaf, 0x17,
	0xff, 0x74, 0xb0, 0xc2, 0x3c, 0xd7, 0x17, 0x30, 0xf7, 0x4e, 0x37, 0xcc, 0x29, 0x4c, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xca, 0xf9, 0x43, 0x67, 0xe8, 0x03, 0x00, 0x00,
}
