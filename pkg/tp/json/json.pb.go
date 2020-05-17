// Code generated by protoc-gen-go. DO NOT EDIT.
// source: json.proto

package json

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

type Action int32

const (
	Action_MERGE  Action = 0
	Action_INSERT Action = 1
	Action_UPDATE Action = 2
	Action_UPSERT Action = 3
	Action_DELETE Action = 4
)

var Action_name = map[int32]string{
	0: "MERGE",
	1: "INSERT",
	2: "UPDATE",
	3: "UPSERT",
	4: "DELETE",
}

var Action_value = map[string]int32{
	"MERGE":  0,
	"INSERT": 1,
	"UPDATE": 2,
	"UPSERT": 3,
	"DELETE": 4,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_93d0772b96a2f7bf, []int{0}
}

type JSONPayload struct {
	Action               Action   `protobuf:"varint,1,opt,name=action,proto3,enum=json.Action" json:"action,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JSONPayload) Reset()         { *m = JSONPayload{} }
func (m *JSONPayload) String() string { return proto.CompactTextString(m) }
func (*JSONPayload) ProtoMessage()    {}
func (*JSONPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_93d0772b96a2f7bf, []int{0}
}

func (m *JSONPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JSONPayload.Unmarshal(m, b)
}
func (m *JSONPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JSONPayload.Marshal(b, m, deterministic)
}
func (m *JSONPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JSONPayload.Merge(m, src)
}
func (m *JSONPayload) XXX_Size() int {
	return xxx_messageInfo_JSONPayload.Size(m)
}
func (m *JSONPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_JSONPayload.DiscardUnknown(m)
}

var xxx_messageInfo_JSONPayload proto.InternalMessageInfo

func (m *JSONPayload) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_MERGE
}

func (m *JSONPayload) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("json.Action", Action_name, Action_value)
	proto.RegisterType((*JSONPayload)(nil), "json.JSONPayload")
}

func init() { proto.RegisterFile("json.proto", fileDescriptor_93d0772b96a2f7bf) }

var fileDescriptor_93d0772b96a2f7bf = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x2a, 0xce, 0xcf,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x7c, 0xb9, 0xb8, 0xbd, 0x82,
	0xfd, 0xfd, 0x02, 0x12, 0x2b, 0x73, 0xf2, 0x13, 0x53, 0x84, 0x54, 0xb8, 0xd8, 0x12, 0x93, 0x4b,
	0x32, 0xf3, 0xf3, 0x24, 0x18, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0x78, 0xf4, 0xc0, 0x3a, 0x1c, 0xc1,
	0x62, 0x41, 0x50, 0x39, 0x21, 0x09, 0x2e, 0xf6, 0x02, 0x88, 0x06, 0x09, 0x26, 0x05, 0x46, 0x0d,
	0x9e, 0x20, 0x18, 0x57, 0xcb, 0x99, 0x8b, 0x0d, 0xa2, 0x56, 0x88, 0x93, 0x8b, 0xd5, 0xd7, 0x35,
	0xc8, 0xdd, 0x55, 0x80, 0x41, 0x88, 0x8b, 0x8b, 0xcd, 0xd3, 0x2f, 0xd8, 0x35, 0x28, 0x44, 0x80,
	0x11, 0xc4, 0x0e, 0x0d, 0x70, 0x71, 0x0c, 0x71, 0x15, 0x60, 0x82, 0xb0, 0xc1, 0xe2, 0xcc, 0x20,
	0xb6, 0x8b, 0xab, 0x8f, 0x6b, 0x88, 0xab, 0x00, 0x4b, 0x12, 0x1b, 0xd8, 0x81, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x81, 0x15, 0xea, 0x20, 0xae, 0x00, 0x00, 0x00,
}