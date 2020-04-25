// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flamed.proto

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

type FlameAction_FlameActionType int32

const (
	FlameAction_CREATE FlameAction_FlameActionType = 0
	FlameAction_UPDATE FlameAction_FlameActionType = 1
	FlameAction_DELETE FlameAction_FlameActionType = 2
)

var FlameAction_FlameActionType_name = map[int32]string{
	0: "CREATE",
	1: "UPDATE",
	2: "DELETE",
}

var FlameAction_FlameActionType_value = map[string]int32{
	"CREATE": 0,
	"UPDATE": 1,
	"DELETE": 2,
}

func (x FlameAction_FlameActionType) String() string {
	return proto.EnumName(FlameAction_FlameActionType_name, int32(x))
}

func (FlameAction_FlameActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{3, 0}
}

type FlameKVPair struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlameKVPair) Reset()         { *m = FlameKVPair{} }
func (m *FlameKVPair) String() string { return proto.CompactTextString(m) }
func (*FlameKVPair) ProtoMessage()    {}
func (*FlameKVPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{0}
}

func (m *FlameKVPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameKVPair.Unmarshal(m, b)
}
func (m *FlameKVPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameKVPair.Marshal(b, m, deterministic)
}
func (m *FlameKVPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameKVPair.Merge(m, src)
}
func (m *FlameKVPair) XXX_Size() int {
	return xxx_messageInfo_FlameKVPair.Size(m)
}
func (m *FlameKVPair) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameKVPair.DiscardUnknown(m)
}

var xxx_messageInfo_FlameKVPair proto.InternalMessageInfo

func (m *FlameKVPair) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *FlameKVPair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type FlameSnapshot struct {
	Version              uint32         `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Length               uint64         `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	FlameKVPairList      []*FlameKVPair `protobuf:"bytes,3,rep,name=flameKVPairList,proto3" json:"flameKVPairList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FlameSnapshot) Reset()         { *m = FlameSnapshot{} }
func (m *FlameSnapshot) String() string { return proto.CompactTextString(m) }
func (*FlameSnapshot) ProtoMessage()    {}
func (*FlameSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{1}
}

func (m *FlameSnapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameSnapshot.Unmarshal(m, b)
}
func (m *FlameSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameSnapshot.Marshal(b, m, deterministic)
}
func (m *FlameSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameSnapshot.Merge(m, src)
}
func (m *FlameSnapshot) XXX_Size() int {
	return xxx_messageInfo_FlameSnapshot.Size(m)
}
func (m *FlameSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_FlameSnapshot proto.InternalMessageInfo

func (m *FlameSnapshot) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *FlameSnapshot) GetLength() uint64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *FlameSnapshot) GetFlameKVPairList() []*FlameKVPair {
	if m != nil {
		return m.FlameKVPairList
	}
	return nil
}

type FlameEntry struct {
	Namespace            []byte   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Key                  []byte   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlameEntry) Reset()         { *m = FlameEntry{} }
func (m *FlameEntry) String() string { return proto.CompactTextString(m) }
func (*FlameEntry) ProtoMessage()    {}
func (*FlameEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{2}
}

func (m *FlameEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameEntry.Unmarshal(m, b)
}
func (m *FlameEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameEntry.Marshal(b, m, deterministic)
}
func (m *FlameEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameEntry.Merge(m, src)
}
func (m *FlameEntry) XXX_Size() int {
	return xxx_messageInfo_FlameEntry.Size(m)
}
func (m *FlameEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameEntry.DiscardUnknown(m)
}

var xxx_messageInfo_FlameEntry proto.InternalMessageInfo

func (m *FlameEntry) GetNamespace() []byte {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *FlameEntry) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *FlameEntry) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type FlameAction struct {
	FlameActionType      FlameAction_FlameActionType `protobuf:"varint,1,opt,name=flameActionType,proto3,enum=pb.FlameAction_FlameActionType" json:"flameActionType,omitempty"`
	FlameEntry           *FlameEntry                 `protobuf:"bytes,2,opt,name=flameEntry,proto3" json:"flameEntry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *FlameAction) Reset()         { *m = FlameAction{} }
func (m *FlameAction) String() string { return proto.CompactTextString(m) }
func (*FlameAction) ProtoMessage()    {}
func (*FlameAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{3}
}

func (m *FlameAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameAction.Unmarshal(m, b)
}
func (m *FlameAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameAction.Marshal(b, m, deterministic)
}
func (m *FlameAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameAction.Merge(m, src)
}
func (m *FlameAction) XXX_Size() int {
	return xxx_messageInfo_FlameAction.Size(m)
}
func (m *FlameAction) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameAction.DiscardUnknown(m)
}

var xxx_messageInfo_FlameAction proto.InternalMessageInfo

func (m *FlameAction) GetFlameActionType() FlameAction_FlameActionType {
	if m != nil {
		return m.FlameActionType
	}
	return FlameAction_CREATE
}

func (m *FlameAction) GetFlameEntry() *FlameEntry {
	if m != nil {
		return m.FlameEntry
	}
	return nil
}

type FlameBatch struct {
	FlameActionList      []*FlameAction `protobuf:"bytes,1,rep,name=flameActionList,proto3" json:"flameActionList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FlameBatch) Reset()         { *m = FlameBatch{} }
func (m *FlameBatch) String() string { return proto.CompactTextString(m) }
func (*FlameBatch) ProtoMessage()    {}
func (*FlameBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{4}
}

func (m *FlameBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameBatch.Unmarshal(m, b)
}
func (m *FlameBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameBatch.Marshal(b, m, deterministic)
}
func (m *FlameBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameBatch.Merge(m, src)
}
func (m *FlameBatch) XXX_Size() int {
	return xxx_messageInfo_FlameBatch.Size(m)
}
func (m *FlameBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameBatch.DiscardUnknown(m)
}

var xxx_messageInfo_FlameBatch proto.InternalMessageInfo

func (m *FlameBatch) GetFlameActionList() []*FlameAction {
	if m != nil {
		return m.FlameActionList
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.FlameAction_FlameActionType", FlameAction_FlameActionType_name, FlameAction_FlameActionType_value)
	proto.RegisterType((*FlameKVPair)(nil), "pb.FlameKVPair")
	proto.RegisterType((*FlameSnapshot)(nil), "pb.FlameSnapshot")
	proto.RegisterType((*FlameEntry)(nil), "pb.FlameEntry")
	proto.RegisterType((*FlameAction)(nil), "pb.FlameAction")
	proto.RegisterType((*FlameBatch)(nil), "pb.FlameBatch")
}

func init() { proto.RegisterFile("flamed.proto", fileDescriptor_388b6a0687b80922) }

var fileDescriptor_388b6a0687b80922 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x4f, 0x83, 0x40,
	0x10, 0x15, 0xd0, 0x1a, 0xa7, 0x1f, 0x90, 0x8d, 0x31, 0x1c, 0x4c, 0x6c, 0x38, 0xf5, 0xc4, 0x01,
	0xd3, 0x83, 0xc7, 0x6a, 0x57, 0x63, 0xec, 0xa1, 0x59, 0xd1, 0xfb, 0x82, 0x8b, 0x25, 0x52, 0xd8,
	0xc0, 0xda, 0x84, 0xc4, 0xff, 0xe6, 0x5f, 0x33, 0xbb, 0x4b, 0x81, 0x10, 0x6f, 0x6f, 0x66, 0xde,
	0xbc, 0x7d, 0xf3, 0xb2, 0x30, 0x49, 0x32, 0xba, 0x67, 0x1f, 0x3e, 0x2f, 0x0b, 0x51, 0x20, 0x93,
	0x47, 0xde, 0x12, 0xc6, 0x8f, 0xb2, 0xf7, 0xf2, 0xbe, 0xa5, 0x69, 0x89, 0x1c, 0xb0, 0xbe, 0x58,
	0xed, 0x1a, 0x73, 0x63, 0x31, 0x21, 0x12, 0xa2, 0x4b, 0x38, 0x3b, 0xd0, 0xec, 0x9b, 0xb9, 0xa6,
	0xea, 0xe9, 0xc2, 0xfb, 0x81, 0xa9, 0x5a, 0x7b, 0xcd, 0x29, 0xaf, 0x76, 0x85, 0x40, 0x2e, 0x9c,
	0x1f, 0x58, 0x59, 0xa5, 0x45, 0xae, 0x96, 0xa7, 0xe4, 0x58, 0xa2, 0x2b, 0x18, 0x65, 0x2c, 0xff,
	0x14, 0x3b, 0xa5, 0x70, 0x4a, 0x9a, 0x0a, 0xdd, 0x81, 0x9d, 0x74, 0x2f, 0x6f, 0xd2, 0x4a, 0xb8,
	0xd6, 0xdc, 0x5a, 0x8c, 0x03, 0xdb, 0xe7, 0x91, 0xdf, 0x33, 0x45, 0x86, 0x3c, 0x8f, 0x00, 0xa8,
	0x39, 0xce, 0x45, 0x59, 0xa3, 0x6b, 0xb8, 0xc8, 0xe9, 0x9e, 0x55, 0x9c, 0xc6, 0xac, 0x71, 0xde,
	0x35, 0x8e, 0x17, 0x99, 0xff, 0x5c, 0x64, 0xf5, 0x2f, 0xfa, 0x35, 0x9a, 0x24, 0x56, 0xb1, 0x90,
	0xb6, 0x9f, 0x1b, 0x7b, 0xba, 0x0c, 0x6b, 0xae, 0xb5, 0x67, 0xc1, 0x4d, 0x6b, 0x4f, 0x8f, 0xfa,
	0x58, 0xd2, 0xc8, 0x70, 0x0f, 0xf9, 0x00, 0x49, 0x6b, 0x57, 0x39, 0x19, 0x07, 0xb3, 0x56, 0x45,
	0x75, 0x49, 0x8f, 0xe1, 0x2d, 0xc1, 0x1e, 0x68, 0x22, 0x80, 0xd1, 0x03, 0xc1, 0xab, 0x10, 0x3b,
	0x27, 0x12, 0xbf, 0x6d, 0xd7, 0x12, 0x1b, 0x12, 0xaf, 0xf1, 0x06, 0x87, 0xd8, 0x31, 0xbd, 0xa7,
	0x26, 0x95, 0x7b, 0x2a, 0xe2, 0x2e, 0x5e, 0x2d, 0xa2, 0xe2, 0x35, 0x06, 0xf1, 0xea, 0x11, 0x19,
	0xf2, 0xa2, 0x91, 0xfa, 0x1e, 0xb7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x75, 0x2e, 0x07, 0x39,
	0x2e, 0x02, 0x00, 0x00,
}