// Code generated by protoc-gen-go. DO NOT EDIT.
// source: intkeytp.proto

package intkey

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

type Verb int32

const (
	Verb_INSERT    Verb = 0
	Verb_UPSERT    Verb = 1
	Verb_DELETE    Verb = 2
	Verb_INCREMENT Verb = 3
	Verb_DECREMENT Verb = 4
)

var Verb_name = map[int32]string{
	0: "INSERT",
	1: "UPSERT",
	2: "DELETE",
	3: "INCREMENT",
	4: "DECREMENT",
}

var Verb_value = map[string]int32{
	"INSERT":    0,
	"UPSERT":    1,
	"DELETE":    2,
	"INCREMENT": 3,
	"DECREMENT": 4,
}

func (x Verb) String() string {
	return proto.EnumName(Verb_name, int32(x))
}

func (Verb) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c97da48b2fb36873, []int{0}
}

type IntKeyState struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                uint64   `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntKeyState) Reset()         { *m = IntKeyState{} }
func (m *IntKeyState) String() string { return proto.CompactTextString(m) }
func (*IntKeyState) ProtoMessage()    {}
func (*IntKeyState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c97da48b2fb36873, []int{0}
}

func (m *IntKeyState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntKeyState.Unmarshal(m, b)
}
func (m *IntKeyState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntKeyState.Marshal(b, m, deterministic)
}
func (m *IntKeyState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntKeyState.Merge(m, src)
}
func (m *IntKeyState) XXX_Size() int {
	return xxx_messageInfo_IntKeyState.Size(m)
}
func (m *IntKeyState) XXX_DiscardUnknown() {
	xxx_messageInfo_IntKeyState.DiscardUnknown(m)
}

var xxx_messageInfo_IntKeyState proto.InternalMessageInfo

func (m *IntKeyState) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IntKeyState) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type IntKeyPayload struct {
	Verb                 Verb     `protobuf:"varint,1,opt,name=verb,proto3,enum=intkeytp.Verb" json:"verb,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value                uint64   `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntKeyPayload) Reset()         { *m = IntKeyPayload{} }
func (m *IntKeyPayload) String() string { return proto.CompactTextString(m) }
func (*IntKeyPayload) ProtoMessage()    {}
func (*IntKeyPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_c97da48b2fb36873, []int{1}
}

func (m *IntKeyPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntKeyPayload.Unmarshal(m, b)
}
func (m *IntKeyPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntKeyPayload.Marshal(b, m, deterministic)
}
func (m *IntKeyPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntKeyPayload.Merge(m, src)
}
func (m *IntKeyPayload) XXX_Size() int {
	return xxx_messageInfo_IntKeyPayload.Size(m)
}
func (m *IntKeyPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_IntKeyPayload.DiscardUnknown(m)
}

var xxx_messageInfo_IntKeyPayload proto.InternalMessageInfo

func (m *IntKeyPayload) GetVerb() Verb {
	if m != nil {
		return m.Verb
	}
	return Verb_INSERT
}

func (m *IntKeyPayload) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IntKeyPayload) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterEnum("intkeytp.Verb", Verb_name, Verb_value)
	proto.RegisterType((*IntKeyState)(nil), "intkeytp.IntKeyState")
	proto.RegisterType((*IntKeyPayload)(nil), "intkeytp.IntKeyPayload")
}

func init() { proto.RegisterFile("intkeytp.proto", fileDescriptor_c97da48b2fb36873) }

var fileDescriptor_c97da48b2fb36873 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0xcc, 0x2b, 0xc9,
	0x4e, 0xad, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xcc, 0xb9, 0xb8,
	0x3d, 0xf3, 0x4a, 0xbc, 0x53, 0x2b, 0x83, 0x4b, 0x12, 0x4b, 0x52, 0x85, 0x84, 0xb8, 0x58, 0xf2,
	0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6,
	0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x08, 0x47, 0x29, 0x9a,
	0x8b, 0x17, 0xa2, 0x31, 0x20, 0xb1, 0x32, 0x27, 0x3f, 0x31, 0x45, 0x48, 0x81, 0x8b, 0xa5, 0x2c,
	0xb5, 0x28, 0x09, 0xac, 0x95, 0xcf, 0x88, 0x47, 0x0f, 0x6a, 0x5d, 0x58, 0x6a, 0x51, 0x52, 0x10,
	0x58, 0x06, 0x6e, 0x38, 0x13, 0x36, 0xc3, 0x99, 0x91, 0x0c, 0xd7, 0xf2, 0xe0, 0x62, 0x01, 0xe9,
	0x13, 0xe2, 0xe2, 0x62, 0xf3, 0xf4, 0x0b, 0x76, 0x0d, 0x0a, 0x11, 0x60, 0x00, 0xb1, 0x43, 0x03,
	0xc0, 0x6c, 0x46, 0x10, 0xdb, 0xc5, 0xd5, 0xc7, 0x35, 0xc4, 0x55, 0x80, 0x49, 0x88, 0x97, 0x8b,
	0xd3, 0xd3, 0xcf, 0x39, 0xc8, 0xd5, 0xd7, 0xd5, 0x2f, 0x44, 0x80, 0x19, 0xc4, 0x75, 0x71, 0x85,
	0x71, 0x59, 0x92, 0xd8, 0xc0, 0xde, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x3f, 0x30,
	0xbd, 0xfe, 0x00, 0x00, 0x00,
}
