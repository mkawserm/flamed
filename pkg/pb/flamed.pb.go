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
	FlameAction_APPEND FlameAction_FlameActionType = 2
	FlameAction_DELETE FlameAction_FlameActionType = 3
)

var FlameAction_FlameActionType_name = map[int32]string{
	0: "CREATE",
	1: "UPDATE",
	2: "APPEND",
	3: "DELETE",
}

var FlameAction_FlameActionType_value = map[string]int32{
	"CREATE": 0,
	"UPDATE": 1,
	"APPEND": 2,
	"DELETE": 3,
}

func (x FlameAction_FlameActionType) String() string {
	return proto.EnumName(FlameAction_FlameActionType_name, int32(x))
}

func (FlameAction_FlameActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{1, 0}
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
	return fileDescriptor_388b6a0687b80922, []int{0}
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
	return fileDescriptor_388b6a0687b80922, []int{1}
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
	return fileDescriptor_388b6a0687b80922, []int{2}
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

type FlameSnapshotEntry struct {
	Uid                  []byte   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlameSnapshotEntry) Reset()         { *m = FlameSnapshotEntry{} }
func (m *FlameSnapshotEntry) String() string { return proto.CompactTextString(m) }
func (*FlameSnapshotEntry) ProtoMessage()    {}
func (*FlameSnapshotEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{3}
}

func (m *FlameSnapshotEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameSnapshotEntry.Unmarshal(m, b)
}
func (m *FlameSnapshotEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameSnapshotEntry.Marshal(b, m, deterministic)
}
func (m *FlameSnapshotEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameSnapshotEntry.Merge(m, src)
}
func (m *FlameSnapshotEntry) XXX_Size() int {
	return xxx_messageInfo_FlameSnapshotEntry.Size(m)
}
func (m *FlameSnapshotEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameSnapshotEntry.DiscardUnknown(m)
}

var xxx_messageInfo_FlameSnapshotEntry proto.InternalMessageInfo

func (m *FlameSnapshotEntry) GetUid() []byte {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *FlameSnapshotEntry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type FlameSnapshot struct {
	Version                uint32                `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Length                 uint64                `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	FlameSnapshotEntryList []*FlameSnapshotEntry `protobuf:"bytes,3,rep,name=flameSnapshotEntryList,proto3" json:"flameSnapshotEntryList,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}              `json:"-"`
	XXX_unrecognized       []byte                `json:"-"`
	XXX_sizecache          int32                 `json:"-"`
}

func (m *FlameSnapshot) Reset()         { *m = FlameSnapshot{} }
func (m *FlameSnapshot) String() string { return proto.CompactTextString(m) }
func (*FlameSnapshot) ProtoMessage()    {}
func (*FlameSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{4}
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

func (m *FlameSnapshot) GetFlameSnapshotEntryList() []*FlameSnapshotEntry {
	if m != nil {
		return m.FlameSnapshotEntryList
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.FlameAction_FlameActionType", FlameAction_FlameActionType_name, FlameAction_FlameActionType_value)
	proto.RegisterType((*FlameEntry)(nil), "pb.FlameEntry")
	proto.RegisterType((*FlameAction)(nil), "pb.FlameAction")
	proto.RegisterType((*FlameBatch)(nil), "pb.FlameBatch")
	proto.RegisterType((*FlameSnapshotEntry)(nil), "pb.FlameSnapshotEntry")
	proto.RegisterType((*FlameSnapshot)(nil), "pb.FlameSnapshot")
}

func init() { proto.RegisterFile("flamed.proto", fileDescriptor_388b6a0687b80922) }

var fileDescriptor_388b6a0687b80922 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcd, 0x4e, 0xf2, 0x40,
	0x14, 0xfd, 0x4a, 0xf9, 0x30, 0x5e, 0xfe, 0x9a, 0x1b, 0x43, 0xba, 0x30, 0x91, 0xcc, 0x8a, 0x55,
	0x17, 0xb8, 0xd2, 0x5d, 0x95, 0xd1, 0x98, 0x10, 0x42, 0x46, 0x7c, 0x80, 0x01, 0xa6, 0xd2, 0x08,
	0xd3, 0x09, 0x1d, 0x48, 0x78, 0x0c, 0x9f, 0xcb, 0x97, 0x32, 0x73, 0xa9, 0xfc, 0xd4, 0xb8, 0x3b,
	0xf7, 0xf4, 0xdc, 0xd3, 0x7b, 0x4e, 0x06, 0x1a, 0xc9, 0x52, 0xae, 0xd4, 0x3c, 0x32, 0xeb, 0xcc,
	0x66, 0x58, 0x31, 0x53, 0x26, 0x00, 0x9e, 0x1c, 0xc7, 0xb5, 0x5d, 0xef, 0xf0, 0x1a, 0x2e, 0xb5,
	0x5c, 0xa9, 0xdc, 0xc8, 0x99, 0x0a, 0xbd, 0xae, 0xd7, 0x6b, 0x88, 0x23, 0x81, 0x01, 0xf8, 0x1f,
	0x6a, 0x17, 0x56, 0x88, 0x77, 0x10, 0xaf, 0xe0, 0xff, 0x56, 0x2e, 0x37, 0x2a, 0xf4, 0x89, 0xdb,
	0x0f, 0xec, 0xcb, 0x83, 0x3a, 0x99, 0xc6, 0x33, 0x9b, 0x66, 0x1a, 0x5f, 0xa0, 0x9d, 0x1c, 0xc7,
	0xc9, 0xce, 0xec, 0xbd, 0x5b, 0xfd, 0x9b, 0xc8, 0x4c, 0xa3, 0x13, 0xe5, 0x29, 0x76, 0x32, 0x51,
	0xde, 0xc3, 0x08, 0x20, 0x39, 0x9c, 0x4b, 0x97, 0xd4, 0xfb, 0xad, 0x83, 0x0b, 0xb1, 0xe2, 0x44,
	0xc1, 0x62, 0x68, 0x97, 0x3c, 0x11, 0xa0, 0xf6, 0x28, 0x78, 0x3c, 0xe1, 0xc1, 0x3f, 0x87, 0xdf,
	0xc6, 0x03, 0x87, 0x3d, 0x87, 0xe3, 0xf1, 0x98, 0x8f, 0x06, 0x41, 0xc5, 0xe1, 0x01, 0x1f, 0xf2,
	0x09, 0x0f, 0x7c, 0xf6, 0x5c, 0x34, 0xf4, 0x20, 0xed, 0x6c, 0x81, 0x77, 0x67, 0x59, 0x86, 0x69,
	0x6e, 0x43, 0xaf, 0xeb, 0xf7, 0xea, 0xfd, 0x76, 0x29, 0x8b, 0x28, 0xeb, 0xd8, 0x3d, 0x20, 0x7d,
	0x7f, 0xd5, 0xd2, 0xe4, 0x8b, 0xcc, 0xee, 0x2b, 0x0f, 0xc0, 0xdf, 0xa4, 0xf3, 0xa2, 0x6c, 0x07,
	0x11, 0xa1, 0x3a, 0x97, 0x56, 0x16, 0x3d, 0x13, 0x66, 0x9f, 0x1e, 0x34, 0xcf, 0x96, 0x31, 0x84,
	0x8b, 0xad, 0x5a, 0xe7, 0x69, 0xa6, 0x69, 0xb7, 0x29, 0x7e, 0x46, 0xec, 0x40, 0x6d, 0xa9, 0xf4,
	0xbb, 0x5d, 0x90, 0x43, 0x55, 0x14, 0x13, 0x8e, 0xa0, 0x93, 0xfc, 0xfa, 0x3f, 0x25, 0xf0, 0x29,
	0x41, 0xe7, 0x90, 0xe0, 0x4c, 0x21, 0xfe, 0xd8, 0x9a, 0xd6, 0xe8, 0x15, 0xdd, 0x7e, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xa3, 0x4d, 0xca, 0x86, 0x55, 0x02, 0x00, 0x00,
}
