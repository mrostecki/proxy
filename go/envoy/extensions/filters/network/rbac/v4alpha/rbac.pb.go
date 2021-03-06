// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/rbac/v4alpha/rbac.proto

package envoy_extensions_filters_network_rbac_v4alpha

import (
	fmt "fmt"
	v4alpha "github.com/cilium/proxy/go/envoy/config/rbac/v4alpha"
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

type RBAC_EnforcementType int32

const (
	// Apply RBAC policies when the first byte of data arrives on the connection.
	RBAC_ONE_TIME_ON_FIRST_BYTE RBAC_EnforcementType = 0
	// Continuously apply RBAC policies as data arrives. Use this mode when
	// using RBAC with message oriented protocols such as Mongo, MySQL, Kafka,
	// etc. when the protocol decoders emit dynamic metadata such as the
	// resources being accessed and the operations on the resources.
	RBAC_CONTINUOUS RBAC_EnforcementType = 1
)

var RBAC_EnforcementType_name = map[int32]string{
	0: "ONE_TIME_ON_FIRST_BYTE",
	1: "CONTINUOUS",
}

var RBAC_EnforcementType_value = map[string]int32{
	"ONE_TIME_ON_FIRST_BYTE": 0,
	"CONTINUOUS":             1,
}

func (x RBAC_EnforcementType) String() string {
	return proto.EnumName(RBAC_EnforcementType_name, int32(x))
}

func (RBAC_EnforcementType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_88effbe9c984e28a, []int{0, 0}
}

// RBAC network filter config.
//
// Header should not be used in rules/shadow_rules in RBAC network filter as
// this information is only available in :ref:`RBAC http filter <config_http_filters_rbac>`.
type RBAC struct {
	// Specify the RBAC rules to be applied globally.
	// If absent, no enforcing RBAC policy will be applied.
	Rules *v4alpha.RBAC `protobuf:"bytes,1,opt,name=rules,proto3" json:"rules,omitempty"`
	// Shadow rules are not enforced by the filter but will emit stats and logs
	// and can be used for rule testing.
	// If absent, no shadow RBAC policy will be applied.
	ShadowRules *v4alpha.RBAC `protobuf:"bytes,2,opt,name=shadow_rules,json=shadowRules,proto3" json:"shadow_rules,omitempty"`
	// The prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,3,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// RBAC enforcement strategy. By default RBAC will be enforced only once
	// when the first byte of data arrives from the downstream. When used in
	// conjunction with filters that emit dynamic metadata after decoding
	// every payload (e.g., Mongo, MySQL, Kafka) set the enforcement type to
	// CONTINUOUS to enforce RBAC policies on every message boundary.
	EnforcementType      RBAC_EnforcementType `protobuf:"varint,4,opt,name=enforcement_type,json=enforcementType,proto3,enum=envoy.extensions.filters.network.rbac.v4alpha.RBAC_EnforcementType" json:"enforcement_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_88effbe9c984e28a, []int{0}
}

func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBAC.Unmarshal(m, b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
}
func (m *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(m, src)
}
func (m *RBAC) XXX_Size() int {
	return xxx_messageInfo_RBAC.Size(m)
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetRules() *v4alpha.RBAC {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *RBAC) GetShadowRules() *v4alpha.RBAC {
	if m != nil {
		return m.ShadowRules
	}
	return nil
}

func (m *RBAC) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RBAC) GetEnforcementType() RBAC_EnforcementType {
	if m != nil {
		return m.EnforcementType
	}
	return RBAC_ONE_TIME_ON_FIRST_BYTE
}

func init() {
	proto.RegisterEnum("envoy.extensions.filters.network.rbac.v4alpha.RBAC_EnforcementType", RBAC_EnforcementType_name, RBAC_EnforcementType_value)
	proto.RegisterType((*RBAC)(nil), "envoy.extensions.filters.network.rbac.v4alpha.RBAC")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/rbac/v4alpha/rbac.proto", fileDescriptor_88effbe9c984e28a)
}

var fileDescriptor_88effbe9c984e28a = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0xcb, 0xd3, 0x40,
	0x1c, 0xc6, 0x4d, 0x5b, 0x5f, 0x79, 0xaf, 0xf2, 0x36, 0x64, 0xd0, 0x52, 0x50, 0x63, 0x71, 0xc8,
	0xe2, 0x1d, 0xb4, 0x15, 0x44, 0x71, 0x30, 0x25, 0x42, 0x07, 0x93, 0x72, 0x4d, 0x07, 0xa7, 0x70,
	0x4d, 0x2f, 0x6d, 0x30, 0xde, 0x85, 0xcb, 0x35, 0x6d, 0x36, 0x47, 0x47, 0x67, 0x3f, 0x8a, 0xbb,
	0xe0, 0xea, 0xd7, 0x71, 0x92, 0xdc, 0x45, 0x4a, 0x95, 0x17, 0xda, 0x2d, 0xc9, 0xf3, 0xfc, 0xfe,
	0xf7, 0x3c, 0xf9, 0x1f, 0x78, 0x49, 0x59, 0xc9, 0x2b, 0x44, 0x0f, 0x92, 0xb2, 0x22, 0xe5, 0xac,
	0x40, 0x49, 0x9a, 0x49, 0x2a, 0x0a, 0xc4, 0xa8, 0xdc, 0x73, 0xf1, 0x11, 0x89, 0x15, 0x89, 0x51,
	0x39, 0x21, 0x59, 0xbe, 0x25, 0xea, 0x05, 0xe6, 0x82, 0x4b, 0x6e, 0x3d, 0x57, 0x24, 0x3c, 0x92,
	0xb0, 0x21, 0x61, 0x43, 0x42, 0x65, 0x6e, 0xc8, 0xc1, 0x33, 0x7d, 0x50, 0xcc, 0x59, 0x92, 0x6e,
	0x6e, 0x1b, 0x3a, 0x78, 0xb4, 0x5b, 0xe7, 0x04, 0x11, 0xc6, 0xb8, 0x24, 0x52, 0xc5, 0x29, 0x24,
	0x91, 0xbb, 0xa2, 0x91, 0x9f, 0xfe, 0x27, 0x97, 0x54, 0xd4, 0x87, 0xa7, 0x6c, 0xd3, 0x58, 0x1e,
	0x96, 0x24, 0x4b, 0xd7, 0x44, 0x52, 0xf4, 0xf7, 0x41, 0x0b, 0xc3, 0xaf, 0x6d, 0xd0, 0xc1, 0xee,
	0xdb, 0xa9, 0xf5, 0x02, 0xdc, 0x15, 0xbb, 0x8c, 0x16, 0x7d, 0xc3, 0x36, 0x9c, 0xee, 0xe8, 0x09,
	0xd4, 0x45, 0x74, 0xb2, 0x93, 0xd0, 0xb0, 0xf6, 0x63, 0xed, 0xb6, 0x5c, 0x70, 0xbf, 0xd8, 0x92,
	0x35, 0xdf, 0x47, 0x9a, 0x6e, 0x9d, 0x47, 0x77, 0x35, 0x84, 0xd5, 0x0c, 0x07, 0x74, 0xeb, 0x3e,
	0x51, 0x2e, 0x68, 0x92, 0x1e, 0xfa, 0x6d, 0xdb, 0x70, 0xae, 0xdd, 0x7b, 0xbf, 0xdd, 0x8e, 0x68,
	0xd9, 0x06, 0x06, 0xb5, 0x36, 0x57, 0x92, 0xc5, 0x80, 0x49, 0x59, 0xc2, 0x45, 0x4c, 0x3f, 0x51,
	0x26, 0x23, 0x59, 0xe5, 0xb4, 0xdf, 0xb1, 0x0d, 0xe7, 0x66, 0x34, 0x85, 0x17, 0xfd, 0x78, 0x95,
	0x02, 0x7a, 0xc7, 0x59, 0x61, 0x95, 0x53, 0xdc, 0xa3, 0xa7, 0x1f, 0x86, 0x6f, 0x40, 0xef, 0x1f,
	0x8f, 0x35, 0x00, 0x0f, 0x02, 0xdf, 0x8b, 0xc2, 0xd9, 0x7b, 0x2f, 0x0a, 0xfc, 0xe8, 0xdd, 0x0c,
	0x2f, 0xc2, 0xc8, 0xfd, 0x10, 0x7a, 0xe6, 0x1d, 0xeb, 0x06, 0x80, 0x69, 0xe0, 0x87, 0x33, 0x7f,
	0x19, 0x2c, 0x17, 0xa6, 0xf1, 0x6a, 0xf2, 0xed, 0xc7, 0x97, 0xc7, 0x08, 0x9c, 0x7b, 0x27, 0xc6,
	0x2a, 0x95, 0x8b, 0xbf, 0x7f, 0xfe, 0xf9, 0xeb, 0xaa, 0x65, 0xb6, 0xc1, 0xeb, 0x94, 0xeb, 0x5a,
	0xb9, 0xe0, 0x87, 0xea, 0xb2, 0x86, 0xee, 0x35, 0x5e, 0x91, 0x78, 0x5e, 0x2f, 0x79, 0x6e, 0xac,
	0xae, 0xd4, 0xb6, 0xc7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xed, 0xf0, 0xa9, 0x64, 0xd9, 0x02,
	0x00, 0x00,
}
