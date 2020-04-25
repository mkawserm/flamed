// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flamedsnapshot.proto

package pb

import (
	fmt "fmt"
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

type FlamedKVPair struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlamedKVPair) Reset()         { *m = FlamedKVPair{} }
func (m *FlamedKVPair) String() string { return proto.CompactTextString(m) }
func (*FlamedKVPair) ProtoMessage()    {}
func (*FlamedKVPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9f66b55daa4590e, []int{0}
}

func (m *FlamedKVPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlamedKVPair.Unmarshal(m, b)
}
func (m *FlamedKVPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlamedKVPair.Marshal(b, m, deterministic)
}
func (m *FlamedKVPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlamedKVPair.Merge(m, src)
}
func (m *FlamedKVPair) XXX_Size() int {
	return xxx_messageInfo_FlamedKVPair.Size(m)
}
func (m *FlamedKVPair) XXX_DiscardUnknown() {
	xxx_messageInfo_FlamedKVPair.DiscardUnknown(m)
}

var xxx_messageInfo_FlamedKVPair proto.InternalMessageInfo

func (m *FlamedKVPair) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *FlamedKVPair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type FlamedSnapshot struct {
	Length               uint64          `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Version              uint32          `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	FlamedKVPair         []*FlamedKVPair `protobuf:"bytes,3,rep,name=flamedKVPair,proto3" json:"flamedKVPair,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FlamedSnapshot) Reset()         { *m = FlamedSnapshot{} }
func (m *FlamedSnapshot) String() string { return proto.CompactTextString(m) }
func (*FlamedSnapshot) ProtoMessage()    {}
func (*FlamedSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9f66b55daa4590e, []int{1}
}

func (m *FlamedSnapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlamedSnapshot.Unmarshal(m, b)
}
func (m *FlamedSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlamedSnapshot.Marshal(b, m, deterministic)
}
func (m *FlamedSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlamedSnapshot.Merge(m, src)
}
func (m *FlamedSnapshot) XXX_Size() int {
	return xxx_messageInfo_FlamedSnapshot.Size(m)
}
func (m *FlamedSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_FlamedSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_FlamedSnapshot proto.InternalMessageInfo

func (m *FlamedSnapshot) GetLength() uint64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *FlamedSnapshot) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *FlamedSnapshot) GetFlamedKVPair() []*FlamedKVPair {
	if m != nil {
		return m.FlamedKVPair
	}
	return nil
}

func init() {
	proto.RegisterType((*FlamedKVPair)(nil), "pb.FlamedKVPair")
	proto.RegisterType((*FlamedSnapshot)(nil), "pb.FlamedSnapshot")
}

func init() { proto.RegisterFile("flamedsnapshot.proto", fileDescriptor_f9f66b55daa4590e) }

var fileDescriptor_f9f66b55daa4590e = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xcb, 0x49, 0xcc,
	0x4d, 0x4d, 0x29, 0xce, 0x4b, 0x2c, 0x28, 0xce, 0xc8, 0x2f, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x2a, 0x48, 0x52, 0x32, 0xe3, 0xe2, 0x71, 0x03, 0xcb, 0x79, 0x87, 0x05, 0x24, 0x66,
	0x16, 0x09, 0x09, 0x70, 0x31, 0x67, 0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x81,
	0x98, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x4c, 0x60, 0x31, 0x08, 0x47,
	0xa9, 0x82, 0x8b, 0x0f, 0xa2, 0x2f, 0x18, 0x6a, 0xa6, 0x90, 0x18, 0x17, 0x5b, 0x4e, 0x6a, 0x5e,
	0x7a, 0x49, 0x06, 0x58, 0x33, 0x4b, 0x10, 0x94, 0x27, 0x24, 0xc1, 0xc5, 0x5e, 0x96, 0x5a, 0x54,
	0x9c, 0x99, 0x9f, 0x07, 0x36, 0x81, 0x37, 0x08, 0xc6, 0x15, 0x32, 0xe1, 0xe2, 0x49, 0x43, 0xb2,
	0x5b, 0x82, 0x59, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x40, 0xaf, 0x20, 0x49, 0x0f, 0xd9, 0x4d, 0x41,
	0x28, 0xaa, 0x9c, 0x38, 0xa2, 0xd8, 0x0a, 0xb2, 0xd3, 0xf5, 0x0b, 0x92, 0x92, 0xd8, 0xc0, 0xde,
	0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x86, 0xaa, 0x44, 0xde, 0x00, 0x00, 0x00,
}
