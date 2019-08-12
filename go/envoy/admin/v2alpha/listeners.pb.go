// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v2alpha/listeners.proto

package envoy_admin_v2alpha

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
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

// Admin endpoint uses this wrapper for `/listeners` to display listener status information.
// See :ref:`/listeners <operations_admin_interface_listeners>` for more information.
type Listeners struct {
	// List of listener statuses.
	ListenerStatuses     []*ListenerStatus `protobuf:"bytes,1,rep,name=listener_statuses,json=listenerStatuses,proto3" json:"listener_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Listeners) Reset()         { *m = Listeners{} }
func (m *Listeners) String() string { return proto.CompactTextString(m) }
func (*Listeners) ProtoMessage()    {}
func (*Listeners) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd3cc5a2daffd3f5, []int{0}
}

func (m *Listeners) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listeners.Unmarshal(m, b)
}
func (m *Listeners) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listeners.Marshal(b, m, deterministic)
}
func (m *Listeners) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listeners.Merge(m, src)
}
func (m *Listeners) XXX_Size() int {
	return xxx_messageInfo_Listeners.Size(m)
}
func (m *Listeners) XXX_DiscardUnknown() {
	xxx_messageInfo_Listeners.DiscardUnknown(m)
}

var xxx_messageInfo_Listeners proto.InternalMessageInfo

func (m *Listeners) GetListenerStatuses() []*ListenerStatus {
	if m != nil {
		return m.ListenerStatuses
	}
	return nil
}

// Details an individual listener's current status.
type ListenerStatus struct {
	// Name of the listener
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The actual local address that the listener is listening on. If a listener was configured
	// to listen on port 0, then this address has the port that was allocated by the OS.
	LocalAddress         *core.Address `protobuf:"bytes,2,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListenerStatus) Reset()         { *m = ListenerStatus{} }
func (m *ListenerStatus) String() string { return proto.CompactTextString(m) }
func (*ListenerStatus) ProtoMessage()    {}
func (*ListenerStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd3cc5a2daffd3f5, []int{1}
}

func (m *ListenerStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerStatus.Unmarshal(m, b)
}
func (m *ListenerStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerStatus.Marshal(b, m, deterministic)
}
func (m *ListenerStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerStatus.Merge(m, src)
}
func (m *ListenerStatus) XXX_Size() int {
	return xxx_messageInfo_ListenerStatus.Size(m)
}
func (m *ListenerStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerStatus proto.InternalMessageInfo

func (m *ListenerStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListenerStatus) GetLocalAddress() *core.Address {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*Listeners)(nil), "envoy.admin.v2alpha.Listeners")
	proto.RegisterType((*ListenerStatus)(nil), "envoy.admin.v2alpha.ListenerStatus")
}

func init() {
	proto.RegisterFile("envoy/admin/v2alpha/listeners.proto", fileDescriptor_cd3cc5a2daffd3f5)
}

var fileDescriptor_cd3cc5a2daffd3f5 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4b, 0x03, 0x31,
	0x14, 0xc6, 0x89, 0x8a, 0xd0, 0x54, 0x45, 0xe3, 0x72, 0x74, 0xf1, 0x6c, 0x97, 0x9b, 0x5e, 0x20,
	0xe2, 0x2c, 0x76, 0x76, 0x38, 0xce, 0x59, 0x4a, 0xec, 0x3d, 0x30, 0x90, 0x5e, 0x42, 0xde, 0x19,
	0xec, 0x7f, 0x2f, 0x4d, 0x93, 0x42, 0xe1, 0xb6, 0x7c, 0xf0, 0xfb, 0xbe, 0xfc, 0x12, 0xbe, 0xc2,
	0x21, 0xba, 0xbd, 0xd4, 0xfd, 0xce, 0x0c, 0x32, 0x2a, 0x6d, 0xfd, 0x8f, 0x96, 0xd6, 0xd0, 0x88,
	0x03, 0x06, 0x02, 0x1f, 0xdc, 0xe8, 0xc4, 0x63, 0x82, 0x20, 0x41, 0x90, 0xa1, 0xc5, 0x53, 0x6e,
	0x7a, 0x23, 0xa3, 0x92, 0x5b, 0x17, 0x50, 0xea, 0xbe, 0x0f, 0x48, 0xb9, 0xb5, 0xfc, 0xe2, 0xb3,
	0x8f, 0x32, 0x24, 0x5a, 0xfe, 0x50, 0x56, 0x37, 0x34, 0xea, 0xf1, 0x97, 0x90, 0x2a, 0x56, 0x5f,
	0x36, 0x73, 0xb5, 0x82, 0x89, 0x79, 0x28, 0xd5, 0xcf, 0x04, 0x77, 0xf7, 0xf6, 0x2c, 0x23, 0x2d,
	0x91, 0xdf, 0x9d, 0x33, 0x42, 0xf0, 0xab, 0x41, 0xef, 0xb0, 0x62, 0x35, 0x6b, 0x66, 0x5d, 0x3a,
	0x8b, 0x37, 0x7e, 0x6b, 0xdd, 0x56, 0xdb, 0x4d, 0x76, 0xab, 0x2e, 0x6a, 0xd6, 0xcc, 0xd5, 0xa2,
	0xdc, 0xe9, 0x0d, 0x44, 0x05, 0x07, 0x7b, 0x78, 0x3f, 0x12, 0xdd, 0x4d, 0x2a, 0xe4, 0xb4, 0x7e,
	0xe5, 0xcf, 0xc6, 0x1d, 0x69, 0x1f, 0xdc, 0xdf, 0x7e, 0x4a, 0x76, 0x7d, 0x32, 0xa1, 0xf6, 0xf0,
	0xf4, 0x96, 0x7d, 0x5f, 0xa7, 0x3f, 0x78, 0xf9, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x79, 0x29,
	0x97, 0x60, 0x01, 0x00, 0x00,
}
