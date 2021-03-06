// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/dubbo_proxy/v3/route.proto

package envoy_extensions_filters_network_dubbo_proxy_v3

import (
	fmt "fmt"
	v3 "github.com/cilium/proxy/go/envoy/config/route/v3"
	v31 "github.com/cilium/proxy/go/envoy/type/matcher/v3"
	v32 "github.com/cilium/proxy/go/envoy/type/v3"
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

// [#next-free-field: 6]
type RouteConfiguration struct {
	// The name of the route configuration. Reserved for future use in asynchronous route discovery.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The interface name of the service.
	Interface string `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	// Which group does the interface belong to.
	Group string `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	// The version number of the interface.
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// The list of routes that will be matched, in order, against incoming requests. The first route
	// that matches will be used.
	Routes               []*Route `protobuf:"bytes,5,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_96150343ae7318d9, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *RouteConfiguration) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *RouteConfiguration) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RouteConfiguration) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Route struct {
	// Route matching parameters.
	Match *RouteMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	// Route request to some upstream cluster.
	Route                *RouteAction `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_96150343ae7318d9, []int{1}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetMatch() *RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Route) GetRoute() *RouteAction {
	if m != nil {
		return m.Route
	}
	return nil
}

type RouteMatch struct {
	// Method level routing matching.
	Method *MethodMatch `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	// Specifies a set of headers that the route should match on. The router will check the request’s
	// headers against all the specified headers in the route config. A match will happen if all the
	// headers in the route are present in the request with the same values (or based on presence if
	// the value field is not in the config).
	Headers              []*v3.HeaderMatcher `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RouteMatch) Reset()         { *m = RouteMatch{} }
func (m *RouteMatch) String() string { return proto.CompactTextString(m) }
func (*RouteMatch) ProtoMessage()    {}
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_96150343ae7318d9, []int{2}
}

func (m *RouteMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteMatch.Unmarshal(m, b)
}
func (m *RouteMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteMatch.Marshal(b, m, deterministic)
}
func (m *RouteMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteMatch.Merge(m, src)
}
func (m *RouteMatch) XXX_Size() int {
	return xxx_messageInfo_RouteMatch.Size(m)
}
func (m *RouteMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RouteMatch proto.InternalMessageInfo

func (m *RouteMatch) GetMethod() *MethodMatch {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *RouteMatch) GetHeaders() []*v3.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

type RouteAction struct {
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	ClusterSpecifier     isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RouteAction) Reset()         { *m = RouteAction{} }
func (m *RouteAction) String() string { return proto.CompactTextString(m) }
func (*RouteAction) ProtoMessage()    {}
func (*RouteAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_96150343ae7318d9, []int{3}
}

func (m *RouteAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAction.Unmarshal(m, b)
}
func (m *RouteAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAction.Marshal(b, m, deterministic)
}
func (m *RouteAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAction.Merge(m, src)
}
func (m *RouteAction) XXX_Size() int {
	return xxx_messageInfo_RouteAction.Size(m)
}
func (m *RouteAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAction.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAction proto.InternalMessageInfo

type isRouteAction_ClusterSpecifier interface {
	isRouteAction_ClusterSpecifier()
}

type RouteAction_Cluster struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type RouteAction_WeightedClusters struct {
	WeightedClusters *v3.WeightedCluster `protobuf:"bytes,2,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

func (m *RouteAction) GetClusterSpecifier() isRouteAction_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *RouteAction) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *RouteAction) GetWeightedClusters() *v3.WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
	}
}

type MethodMatch struct {
	// The name of the method.
	Name *v31.StringMatcher `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Method parameter definition.
	// The key is the parameter index, starting from 0.
	// The value is the parameter matching type.
	ParamsMatch          map[uint32]*MethodMatch_ParameterMatchSpecifier `protobuf:"bytes,2,rep,name=params_match,json=paramsMatch,proto3" json:"params_match,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *MethodMatch) Reset()         { *m = MethodMatch{} }
func (m *MethodMatch) String() string { return proto.CompactTextString(m) }
func (*MethodMatch) ProtoMessage()    {}
func (*MethodMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_96150343ae7318d9, []int{4}
}

func (m *MethodMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MethodMatch.Unmarshal(m, b)
}
func (m *MethodMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MethodMatch.Marshal(b, m, deterministic)
}
func (m *MethodMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MethodMatch.Merge(m, src)
}
func (m *MethodMatch) XXX_Size() int {
	return xxx_messageInfo_MethodMatch.Size(m)
}
func (m *MethodMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_MethodMatch.DiscardUnknown(m)
}

var xxx_messageInfo_MethodMatch proto.InternalMessageInfo

func (m *MethodMatch) GetName() *v31.StringMatcher {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MethodMatch) GetParamsMatch() map[uint32]*MethodMatch_ParameterMatchSpecifier {
	if m != nil {
		return m.ParamsMatch
	}
	return nil
}

// The parameter matching type.
type MethodMatch_ParameterMatchSpecifier struct {
	// Types that are valid to be assigned to ParameterMatchSpecifier:
	//	*MethodMatch_ParameterMatchSpecifier_ExactMatch
	//	*MethodMatch_ParameterMatchSpecifier_RangeMatch
	ParameterMatchSpecifier isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier `protobuf_oneof:"parameter_match_specifier"`
	XXX_NoUnkeyedLiteral    struct{}                                                      `json:"-"`
	XXX_unrecognized        []byte                                                        `json:"-"`
	XXX_sizecache           int32                                                         `json:"-"`
}

func (m *MethodMatch_ParameterMatchSpecifier) Reset()         { *m = MethodMatch_ParameterMatchSpecifier{} }
func (m *MethodMatch_ParameterMatchSpecifier) String() string { return proto.CompactTextString(m) }
func (*MethodMatch_ParameterMatchSpecifier) ProtoMessage()    {}
func (*MethodMatch_ParameterMatchSpecifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_96150343ae7318d9, []int{4, 0}
}

func (m *MethodMatch_ParameterMatchSpecifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Unmarshal(m, b)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Marshal(b, m, deterministic)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Merge(m, src)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Size() int {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Size(m)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_DiscardUnknown() {
	xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.DiscardUnknown(m)
}

var xxx_messageInfo_MethodMatch_ParameterMatchSpecifier proto.InternalMessageInfo

type isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier interface {
	isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier()
}

type MethodMatch_ParameterMatchSpecifier_ExactMatch struct {
	ExactMatch string `protobuf:"bytes,3,opt,name=exact_match,json=exactMatch,proto3,oneof"`
}

type MethodMatch_ParameterMatchSpecifier_RangeMatch struct {
	RangeMatch *v32.Int64Range `protobuf:"bytes,4,opt,name=range_match,json=rangeMatch,proto3,oneof"`
}

func (*MethodMatch_ParameterMatchSpecifier_ExactMatch) isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier() {
}

func (*MethodMatch_ParameterMatchSpecifier_RangeMatch) isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier() {
}

func (m *MethodMatch_ParameterMatchSpecifier) GetParameterMatchSpecifier() isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier {
	if m != nil {
		return m.ParameterMatchSpecifier
	}
	return nil
}

func (m *MethodMatch_ParameterMatchSpecifier) GetExactMatch() string {
	if x, ok := m.GetParameterMatchSpecifier().(*MethodMatch_ParameterMatchSpecifier_ExactMatch); ok {
		return x.ExactMatch
	}
	return ""
}

func (m *MethodMatch_ParameterMatchSpecifier) GetRangeMatch() *v32.Int64Range {
	if x, ok := m.GetParameterMatchSpecifier().(*MethodMatch_ParameterMatchSpecifier_RangeMatch); ok {
		return x.RangeMatch
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MethodMatch_ParameterMatchSpecifier) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MethodMatch_ParameterMatchSpecifier_ExactMatch)(nil),
		(*MethodMatch_ParameterMatchSpecifier_RangeMatch)(nil),
	}
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.extensions.filters.network.dubbo_proxy.v3.RouteConfiguration")
	proto.RegisterType((*Route)(nil), "envoy.extensions.filters.network.dubbo_proxy.v3.Route")
	proto.RegisterType((*RouteMatch)(nil), "envoy.extensions.filters.network.dubbo_proxy.v3.RouteMatch")
	proto.RegisterType((*RouteAction)(nil), "envoy.extensions.filters.network.dubbo_proxy.v3.RouteAction")
	proto.RegisterType((*MethodMatch)(nil), "envoy.extensions.filters.network.dubbo_proxy.v3.MethodMatch")
	proto.RegisterMapType((map[uint32]*MethodMatch_ParameterMatchSpecifier)(nil), "envoy.extensions.filters.network.dubbo_proxy.v3.MethodMatch.ParamsMatchEntry")
	proto.RegisterType((*MethodMatch_ParameterMatchSpecifier)(nil), "envoy.extensions.filters.network.dubbo_proxy.v3.MethodMatch.ParameterMatchSpecifier")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/dubbo_proxy/v3/route.proto", fileDescriptor_96150343ae7318d9)
}

var fileDescriptor_96150343ae7318d9 = []byte{
	// 776 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcf, 0x6b, 0xe3, 0x46,
	0x14, 0x8e, 0xe4, 0xc8, 0x69, 0x9e, 0x5a, 0x70, 0x86, 0x42, 0x14, 0xf7, 0x07, 0x89, 0x29, 0x25,
	0x87, 0x22, 0x51, 0xbb, 0x84, 0xe0, 0xfc, 0x80, 0xda, 0x04, 0xd2, 0x82, 0x4b, 0x50, 0x92, 0xf6,
	0x90, 0x82, 0x99, 0xc8, 0x63, 0x5b, 0x8d, 0x2d, 0x89, 0xd1, 0x48, 0x89, 0x6f, 0xed, 0xad, 0x94,
	0x9e, 0x4a, 0x4f, 0xfb, 0xa7, 0xec, 0x7d, 0x61, 0xaf, 0xfb, 0x17, 0xec, 0x6d, 0xff, 0x82, 0x3d,
	0x2c, 0x7b, 0x5a, 0xe6, 0xcd, 0x38, 0x56, 0xd6, 0xbb, 0x07, 0x3b, 0xb7, 0xd1, 0x9b, 0xf7, 0xbe,
	0xf9, 0xbe, 0xf7, 0xcd, 0x3c, 0xc1, 0x01, 0x8b, 0xf2, 0x78, 0xe2, 0xb1, 0x3b, 0xc1, 0xa2, 0x34,
	0x8c, 0xa3, 0xd4, 0xeb, 0x87, 0x23, 0xc1, 0x78, 0xea, 0x45, 0x4c, 0xdc, 0xc6, 0xfc, 0xc6, 0xeb,
	0x65, 0xd7, 0xd7, 0x71, 0x37, 0xe1, 0xf1, 0xdd, 0xc4, 0xcb, 0x1b, 0x1e, 0x8f, 0x33, 0xc1, 0xdc,
	0x84, 0xc7, 0x22, 0x26, 0x1e, 0x16, 0xbb, 0xb3, 0x62, 0x57, 0x17, 0xbb, 0xba, 0xd8, 0x2d, 0x14,
	0xbb, 0x79, 0xa3, 0xfa, 0x9d, 0x3a, 0x2d, 0x88, 0xa3, 0x7e, 0x38, 0x50, 0x50, 0xf7, 0x98, 0xdd,
	0x20, 0x1e, 0x27, 0x71, 0xc4, 0x22, 0x91, 0x2a, 0xf8, 0x6a, 0x4d, 0x65, 0x8b, 0x49, 0xc2, 0xbc,
	0x31, 0x15, 0xc1, 0x90, 0x71, 0x99, 0x9d, 0x0a, 0x1e, 0x46, 0x03, 0x9d, 0xb3, 0x55, 0xc8, 0x91,
	0x48, 0x34, 0x1a, 0x68, 0x76, 0xd5, 0xaf, 0xb2, 0x5e, 0x42, 0x3d, 0x1a, 0x45, 0xb1, 0xa0, 0x02,
	0xa5, 0xa5, 0x82, 0x8a, 0x6c, 0x8a, 0xbe, 0x33, 0xb7, 0x9d, 0x33, 0x2e, 0x55, 0xcc, 0xc0, 0x37,
	0x73, 0x3a, 0x0a, 0x7b, 0x54, 0x92, 0xd4, 0x0b, 0xb5, 0x51, 0xfb, 0xcf, 0x04, 0xe2, 0x4b, 0xd2,
	0x6d, 0x54, 0x92, 0x71, 0x44, 0x20, 0x04, 0x56, 0x23, 0x3a, 0x66, 0x8e, 0xb1, 0x6d, 0xec, 0xae,
	0xfb, 0xb8, 0x26, 0x5f, 0xc2, 0x7a, 0x18, 0x09, 0xc6, 0xfb, 0x34, 0x60, 0x8e, 0x89, 0x1b, 0xb3,
	0x00, 0xf9, 0x1c, 0xac, 0x01, 0x8f, 0xb3, 0xc4, 0x29, 0xe1, 0x8e, 0xfa, 0x20, 0x0e, 0xac, 0x69,
	0x2e, 0xce, 0x2a, 0xc6, 0xa7, 0x9f, 0xe4, 0x17, 0x28, 0x63, 0xb3, 0x52, 0xc7, 0xda, 0x2e, 0xed,
	0xda, 0xf5, 0x3d, 0x77, 0x41, 0x0b, 0x5c, 0xa4, 0xed, 0x6b, 0x94, 0xe6, 0xcf, 0x4f, 0x9e, 0xfd,
	0xfd, 0xf5, 0x09, 0xb4, 0x15, 0x8a, 0xf2, 0x45, 0x23, 0x7c, 0x18, 0xa0, 0x4e, 0x47, 0xc9, 0x90,
	0x7e, 0xef, 0xce, 0xab, 0xaf, 0xfd, 0x65, 0x82, 0x85, 0x61, 0x72, 0x05, 0x16, 0xfa, 0x85, 0x8d,
	0xb0, 0xeb, 0x07, 0xcb, 0x91, 0xec, 0x48, 0x88, 0xd6, 0x27, 0x6f, 0x5b, 0xd6, 0x3f, 0x86, 0x59,
	0x31, 0x7c, 0x85, 0x49, 0x7e, 0x07, 0x0b, 0xc9, 0x63, 0x33, 0xed, 0xfa, 0xe1, 0x72, 0xe0, 0x3f,
	0x06, 0x92, 0x73, 0x11, 0x1d, 0x41, 0x9b, 0x47, 0xb2, 0x21, 0xfb, 0xb0, 0xb7, 0x5c, 0x43, 0x6a,
	0xaf, 0x0c, 0x80, 0x19, 0x79, 0x72, 0x01, 0xe5, 0x31, 0x13, 0xc3, 0xb8, 0xa7, 0x3b, 0xb1, 0x38,
	0xd9, 0x0e, 0x96, 0x23, 0x9a, 0xaf, 0xb1, 0xc8, 0x31, 0xac, 0x0d, 0x19, 0xed, 0x31, 0x9e, 0x3a,
	0x26, 0xde, 0x82, 0x6f, 0xdc, 0x07, 0x74, 0xd5, 0x13, 0xcd, 0x1b, 0xee, 0x29, 0x66, 0x75, 0xd4,
	0xd3, 0xf1, 0xa7, 0x45, 0xcd, 0x96, 0xd4, 0x78, 0xa4, 0x9f, 0xfe, 0xc2, 0x1a, 0x11, 0xae, 0xf6,
	0xd2, 0x00, 0xbb, 0xd0, 0x48, 0x52, 0x85, 0xb5, 0x60, 0x94, 0xa5, 0x82, 0x71, 0x75, 0xfb, 0x4f,
	0x57, 0xfc, 0x69, 0x80, 0x5c, 0xc2, 0xc6, 0x2d, 0x0b, 0x07, 0x43, 0xc1, 0x7a, 0x5d, 0x1d, 0x4b,
	0xb5, 0x7b, 0xdf, 0x7e, 0x84, 0xf9, 0x6f, 0x3a, 0xbf, 0xad, 0xd2, 0x4f, 0x57, 0xfc, 0xca, 0xed,
	0xc3, 0x50, 0xda, 0x6c, 0x4b, 0x19, 0xc7, 0x70, 0xb8, 0x9c, 0x0c, 0x7d, 0x01, 0x1c, 0xd8, 0xd0,
	0x94, 0xba, 0x69, 0xc2, 0x82, 0xb0, 0x1f, 0x32, 0x4e, 0x4a, 0x6f, 0x5a, 0x46, 0xed, 0x5f, 0x0b,
	0xec, 0x42, 0xf7, 0xc9, 0x7e, 0xe1, 0x71, 0xcf, 0x5a, 0x2e, 0x07, 0x8f, 0xab, 0x87, 0x93, 0x24,
	0x7e, 0x8e, 0xc3, 0x69, 0xda, 0x72, 0x35, 0x02, 0x12, 0xf8, 0x34, 0xa1, 0x9c, 0x8e, 0xd3, 0xae,
	0x7a, 0x15, 0xca, 0xb4, 0xce, 0x63, 0xee, 0x82, 0x7b, 0x86, 0x80, 0xb8, 0x3e, 0x89, 0x04, 0x9f,
	0xf8, 0x76, 0x32, 0x8b, 0x54, 0x5f, 0x1b, 0xb0, 0x89, 0x19, 0x4c, 0x68, 0xff, 0xcf, 0xef, 0xc5,
	0xed, 0x80, 0xcd, 0xee, 0x68, 0x20, 0x34, 0x99, 0x92, 0x76, 0x0b, 0x30, 0xa8, 0xa4, 0x1e, 0x82,
	0x8d, 0x83, 0x54, 0xa7, 0xac, 0xa2, 0xe2, 0xad, 0xa2, 0xe2, 0xbc, 0xe1, 0xfe, 0x14, 0x89, 0xbd,
	0x1f, 0x7c, 0x99, 0x26, 0xab, 0x31, 0x1f, 0xab, 0x9b, 0x57, 0xd2, 0x97, 0x5f, 0xe1, 0x62, 0x61,
	0x5f, 0xe6, 0xf4, 0xcd, 0xb3, 0x6f, 0x7d, 0x01, 0x5b, 0xc9, 0x74, 0x4b, 0xd1, 0x9b, 0xf9, 0x56,
	0xfd, 0xdf, 0x80, 0xca, 0xfb, 0x8d, 0x21, 0x15, 0x28, 0xdd, 0xb0, 0x09, 0xda, 0xf6, 0x99, 0x2f,
	0x97, 0xe4, 0x0f, 0xb0, 0x72, 0x3a, 0xca, 0xa6, 0x13, 0xe4, 0xe2, 0xf1, 0x46, 0xcc, 0x13, 0xf5,
	0xd5, 0x11, 0x4d, 0x73, 0xdf, 0x58, 0xfe, 0xa2, 0x16, 0xce, 0x69, 0x5d, 0x3e, 0xfd, 0xf3, 0xf9,
	0x8b, 0xb2, 0x59, 0x31, 0xe1, 0x28, 0x8c, 0x15, 0x63, 0x95, 0xbd, 0x20, 0xf9, 0x96, 0x9a, 0x4f,
	0x67, 0xf2, 0x3f, 0x76, 0x66, 0x5c, 0x97, 0xf1, 0x87, 0xd6, 0x78, 0x17, 0x00, 0x00, 0xff, 0xff,
	0x8a, 0x4a, 0x15, 0x67, 0x08, 0x08, 0x00, 0x00,
}
