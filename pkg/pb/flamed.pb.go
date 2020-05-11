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

type StateSnapshot struct {
	Uid                  []byte   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateSnapshot) Reset()         { *m = StateSnapshot{} }
func (m *StateSnapshot) String() string { return proto.CompactTextString(m) }
func (*StateSnapshot) ProtoMessage()    {}
func (*StateSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{0}
}

func (m *StateSnapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateSnapshot.Unmarshal(m, b)
}
func (m *StateSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateSnapshot.Marshal(b, m, deterministic)
}
func (m *StateSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateSnapshot.Merge(m, src)
}
func (m *StateSnapshot) XXX_Size() int {
	return xxx_messageInfo_StateSnapshot.Size(m)
}
func (m *StateSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_StateSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_StateSnapshot proto.InternalMessageInfo

func (m *StateSnapshot) GetUid() []byte {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *StateSnapshot) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Transaction struct {
	Namespace            []byte   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Family               string   `protobuf:"bytes,2,opt,name=family,proto3" json:"family,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{1}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetNamespace() []byte {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *Transaction) GetFamily() string {
	if m != nil {
		return m.Family
	}
	return ""
}

func (m *Transaction) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Transaction) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Proposal struct {
	Uuid                 string         `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	CreatedAt            uint64         `protobuf:"fixed64,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Transactions         []*Transaction `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{2}
}

func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proposal.Unmarshal(m, b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return xxx_messageInfo_Proposal.Size(m)
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

func (m *Proposal) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Proposal) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Proposal) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func init() {
	proto.RegisterType((*StateSnapshot)(nil), "pb.StateSnapshot")
	proto.RegisterType((*Transaction)(nil), "pb.Transaction")
	proto.RegisterType((*Proposal)(nil), "pb.Proposal")
}

func init() { proto.RegisterFile("flamed.proto", fileDescriptor_388b6a0687b80922) }

var fileDescriptor_388b6a0687b80922 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4e, 0xc3, 0x30,
	0x10, 0x45, 0xd5, 0xa6, 0x0a, 0x64, 0x1a, 0x04, 0xf2, 0x02, 0x79, 0xc1, 0xa2, 0xca, 0xaa, 0xab,
	0x2c, 0xa8, 0x38, 0x00, 0x37, 0x40, 0x2e, 0x17, 0x98, 0xc4, 0xae, 0x88, 0x94, 0x78, 0x8c, 0x3d,
	0x05, 0xf5, 0xf6, 0xc8, 0xa3, 0x46, 0xa1, 0xbb, 0xff, 0xff, 0xd8, 0xdf, 0xcf, 0x03, 0xf5, 0x69,
	0xc4, 0xc9, 0xd9, 0x36, 0x44, 0x62, 0x52, 0xeb, 0xd0, 0x35, 0x6f, 0xf0, 0x70, 0x64, 0x64, 0x77,
	0xf4, 0x18, 0xd2, 0x17, 0xb1, 0x7a, 0x82, 0xe2, 0x3c, 0x58, 0xbd, 0xda, 0xad, 0xf6, 0xb5, 0xc9,
	0x52, 0x29, 0xd8, 0x58, 0x64, 0xd4, 0x6b, 0x89, 0x44, 0x37, 0xbf, 0xb0, 0xfd, 0x8c, 0xe8, 0x13,
	0xf6, 0x3c, 0x90, 0x57, 0x2f, 0x50, 0x79, 0x9c, 0x5c, 0x0a, 0xd8, 0xbb, 0xeb, 0xd5, 0x25, 0x50,
	0xcf, 0x50, 0x9e, 0x70, 0x1a, 0xc6, 0x8b, 0x54, 0x54, 0xe6, 0xea, 0x94, 0x86, 0xbb, 0x1f, 0x17,
	0xd3, 0x40, 0x5e, 0x17, 0x32, 0x98, 0x6d, 0x9e, 0x04, 0xbc, 0x8c, 0x84, 0x56, 0x6f, 0xa4, 0x6d,
	0xb6, 0xcd, 0x37, 0xdc, 0x7f, 0x44, 0x0a, 0x94, 0x70, 0xcc, 0x60, 0xe7, 0x99, 0xb5, 0x32, 0xa2,
	0x33, 0x49, 0x1f, 0x1d, 0xb2, 0xb3, 0xef, 0x2c, 0xcf, 0x95, 0x66, 0x09, 0xd4, 0x01, 0x6a, 0x5e,
	0xb0, 0x93, 0x2e, 0x76, 0xc5, 0x7e, 0xfb, 0xfa, 0xd8, 0x86, 0xae, 0xfd, 0xf7, 0x1d, 0x73, 0x73,
	0xa8, 0x2b, 0x65, 0x5b, 0x87, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0x3b, 0xd5, 0x4a, 0x3d,
	0x01, 0x00, 0x00,
}
