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

type IndexFieldType int32

const (
	IndexFieldType_TEXT      IndexFieldType = 0
	IndexFieldType_NUMERIC   IndexFieldType = 1
	IndexFieldType_BOOLEAN   IndexFieldType = 2
	IndexFieldType_GEO_POINT IndexFieldType = 3
	IndexFieldType_DATE_TIME IndexFieldType = 4
)

var IndexFieldType_name = map[int32]string{
	0: "TEXT",
	1: "NUMERIC",
	2: "BOOLEAN",
	3: "GEO_POINT",
	4: "DATE_TIME",
}

var IndexFieldType_value = map[string]int32{
	"TEXT":      0,
	"NUMERIC":   1,
	"BOOLEAN":   2,
	"GEO_POINT": 3,
	"DATE_TIME": 4,
}

func (x IndexFieldType) String() string {
	return proto.EnumName(IndexFieldType_name, int32(x))
}

func (IndexFieldType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{0}
}

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

type StateEntry struct {
	Namespace            []byte   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	FamilyName           string   `protobuf:"bytes,2,opt,name=familyName,proto3" json:"familyName,omitempty"`
	FamilyVersion        string   `protobuf:"bytes,3,opt,name=familyVersion,proto3" json:"familyVersion,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateEntry) Reset()         { *m = StateEntry{} }
func (m *StateEntry) String() string { return proto.CompactTextString(m) }
func (*StateEntry) ProtoMessage()    {}
func (*StateEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{1}
}

func (m *StateEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateEntry.Unmarshal(m, b)
}
func (m *StateEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateEntry.Marshal(b, m, deterministic)
}
func (m *StateEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateEntry.Merge(m, src)
}
func (m *StateEntry) XXX_Size() int {
	return xxx_messageInfo_StateEntry.Size(m)
}
func (m *StateEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_StateEntry.DiscardUnknown(m)
}

var xxx_messageInfo_StateEntry proto.InternalMessageInfo

func (m *StateEntry) GetNamespace() []byte {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *StateEntry) GetFamilyName() string {
	if m != nil {
		return m.FamilyName
	}
	return ""
}

func (m *StateEntry) GetFamilyVersion() string {
	if m != nil {
		return m.FamilyVersion
	}
	return ""
}

func (m *StateEntry) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Transaction struct {
	Namespace            []byte   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	FamilyName           string   `protobuf:"bytes,2,opt,name=familyName,proto3" json:"familyName,omitempty"`
	FamilyVersion        string   `protobuf:"bytes,3,opt,name=familyVersion,proto3" json:"familyVersion,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{2}
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

func (m *Transaction) GetFamilyName() string {
	if m != nil {
		return m.FamilyName
	}
	return ""
}

func (m *Transaction) GetFamilyVersion() string {
	if m != nil {
		return m.FamilyVersion
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
	Meta                 []byte         `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	CreatedAt            uint64         `protobuf:"fixed64,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Transactions         []*Transaction `protobuf:"bytes,5,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{3}
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

func (m *Proposal) GetMeta() []byte {
	if m != nil {
		return m.Meta
	}
	return nil
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

type IndexField struct {
	IndexFieldType       IndexFieldType `protobuf:"varint,1,opt,name=indexFieldType,proto3,enum=pb.IndexFieldType" json:"indexFieldType,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Analyzer             string         `protobuf:"bytes,3,opt,name=analyzer,proto3" json:"analyzer,omitempty"`
	Enabled              bool           `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Index                bool           `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
	Store                bool           `protobuf:"varint,6,opt,name=store,proto3" json:"store,omitempty"`
	IncludeTermVectors   bool           `protobuf:"varint,7,opt,name=includeTermVectors,proto3" json:"includeTermVectors,omitempty"`
	IncludeInAll         bool           `protobuf:"varint,8,opt,name=includeInAll,proto3" json:"includeInAll,omitempty"`
	DocValues            bool           `protobuf:"varint,9,opt,name=docValues,proto3" json:"docValues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *IndexField) Reset()         { *m = IndexField{} }
func (m *IndexField) String() string { return proto.CompactTextString(m) }
func (*IndexField) ProtoMessage()    {}
func (*IndexField) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{4}
}

func (m *IndexField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexField.Unmarshal(m, b)
}
func (m *IndexField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexField.Marshal(b, m, deterministic)
}
func (m *IndexField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexField.Merge(m, src)
}
func (m *IndexField) XXX_Size() int {
	return xxx_messageInfo_IndexField.Size(m)
}
func (m *IndexField) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexField.DiscardUnknown(m)
}

var xxx_messageInfo_IndexField proto.InternalMessageInfo

func (m *IndexField) GetIndexFieldType() IndexFieldType {
	if m != nil {
		return m.IndexFieldType
	}
	return IndexFieldType_TEXT
}

func (m *IndexField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IndexField) GetAnalyzer() string {
	if m != nil {
		return m.Analyzer
	}
	return ""
}

func (m *IndexField) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *IndexField) GetIndex() bool {
	if m != nil {
		return m.Index
	}
	return false
}

func (m *IndexField) GetStore() bool {
	if m != nil {
		return m.Store
	}
	return false
}

func (m *IndexField) GetIncludeTermVectors() bool {
	if m != nil {
		return m.IncludeTermVectors
	}
	return false
}

func (m *IndexField) GetIncludeInAll() bool {
	if m != nil {
		return m.IncludeInAll
	}
	return false
}

func (m *IndexField) GetDocValues() bool {
	if m != nil {
		return m.DocValues
	}
	return false
}

type IndexDocument struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Default              bool          `protobuf:"varint,2,opt,name=default,proto3" json:"default,omitempty"`
	IndexFieldList       []*IndexField `protobuf:"bytes,3,rep,name=indexFieldList,proto3" json:"indexFieldList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *IndexDocument) Reset()         { *m = IndexDocument{} }
func (m *IndexDocument) String() string { return proto.CompactTextString(m) }
func (*IndexDocument) ProtoMessage()    {}
func (*IndexDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{5}
}

func (m *IndexDocument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexDocument.Unmarshal(m, b)
}
func (m *IndexDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexDocument.Marshal(b, m, deterministic)
}
func (m *IndexDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexDocument.Merge(m, src)
}
func (m *IndexDocument) XXX_Size() int {
	return xxx_messageInfo_IndexDocument.Size(m)
}
func (m *IndexDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexDocument.DiscardUnknown(m)
}

var xxx_messageInfo_IndexDocument proto.InternalMessageInfo

func (m *IndexDocument) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IndexDocument) GetDefault() bool {
	if m != nil {
		return m.Default
	}
	return false
}

func (m *IndexDocument) GetIndexFieldList() []*IndexField {
	if m != nil {
		return m.IndexFieldList
	}
	return nil
}

type IndexMeta struct {
	Namespace            []byte           `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Version              uint32           `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Enabled              bool             `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Default              bool             `protobuf:"varint,4,opt,name=default,proto3" json:"default,omitempty"`
	CreatedAt            uint64           `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            uint64           `protobuf:"varint,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	IndexDocumentList    []*IndexDocument `protobuf:"bytes,7,rep,name=indexDocumentList,proto3" json:"indexDocumentList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *IndexMeta) Reset()         { *m = IndexMeta{} }
func (m *IndexMeta) String() string { return proto.CompactTextString(m) }
func (*IndexMeta) ProtoMessage()    {}
func (*IndexMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{6}
}

func (m *IndexMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexMeta.Unmarshal(m, b)
}
func (m *IndexMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexMeta.Marshal(b, m, deterministic)
}
func (m *IndexMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexMeta.Merge(m, src)
}
func (m *IndexMeta) XXX_Size() int {
	return xxx_messageInfo_IndexMeta.Size(m)
}
func (m *IndexMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexMeta.DiscardUnknown(m)
}

var xxx_messageInfo_IndexMeta proto.InternalMessageInfo

func (m *IndexMeta) GetNamespace() []byte {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *IndexMeta) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *IndexMeta) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *IndexMeta) GetDefault() bool {
	if m != nil {
		return m.Default
	}
	return false
}

func (m *IndexMeta) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *IndexMeta) GetUpdatedAt() uint64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *IndexMeta) GetIndexDocumentList() []*IndexDocument {
	if m != nil {
		return m.IndexDocumentList
	}
	return nil
}

type IndexMetaList struct {
	IndexMetaList        []*IndexMeta `protobuf:"bytes,1,rep,name=indexMetaList,proto3" json:"indexMetaList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *IndexMetaList) Reset()         { *m = IndexMetaList{} }
func (m *IndexMetaList) String() string { return proto.CompactTextString(m) }
func (*IndexMetaList) ProtoMessage()    {}
func (*IndexMetaList) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{7}
}

func (m *IndexMetaList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexMetaList.Unmarshal(m, b)
}
func (m *IndexMetaList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexMetaList.Marshal(b, m, deterministic)
}
func (m *IndexMetaList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexMetaList.Merge(m, src)
}
func (m *IndexMetaList) XXX_Size() int {
	return xxx_messageInfo_IndexMetaList.Size(m)
}
func (m *IndexMetaList) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexMetaList.DiscardUnknown(m)
}

var xxx_messageInfo_IndexMetaList proto.InternalMessageInfo

func (m *IndexMetaList) GetIndexMetaList() []*IndexMeta {
	if m != nil {
		return m.IndexMetaList
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.IndexFieldType", IndexFieldType_name, IndexFieldType_value)
	proto.RegisterType((*StateSnapshot)(nil), "pb.StateSnapshot")
	proto.RegisterType((*StateEntry)(nil), "pb.StateEntry")
	proto.RegisterType((*Transaction)(nil), "pb.Transaction")
	proto.RegisterType((*Proposal)(nil), "pb.Proposal")
	proto.RegisterType((*IndexField)(nil), "pb.IndexField")
	proto.RegisterType((*IndexDocument)(nil), "pb.IndexDocument")
	proto.RegisterType((*IndexMeta)(nil), "pb.IndexMeta")
	proto.RegisterType((*IndexMetaList)(nil), "pb.IndexMetaList")
}

func init() { proto.RegisterFile("flamed.proto", fileDescriptor_388b6a0687b80922) }

var fileDescriptor_388b6a0687b80922 = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xfd, 0x39, 0x49, 0x63, 0x7b, 0x1a, 0xe7, 0xe7, 0xae, 0x38, 0xac, 0x50, 0x85, 0x2a, 0x8b,
	0x43, 0xc5, 0x21, 0x87, 0x56, 0x70, 0xe0, 0x82, 0x02, 0x35, 0xc8, 0x52, 0x9b, 0x54, 0x5b, 0x53,
	0x71, 0xab, 0x36, 0xf6, 0x56, 0x58, 0x5a, 0xaf, 0x2d, 0x7b, 0x8d, 0x08, 0x77, 0x2e, 0xf0, 0x65,
	0xf8, 0x70, 0x7c, 0x00, 0xb4, 0xeb, 0xff, 0x05, 0x71, 0xe5, 0x36, 0xef, 0xcd, 0xf3, 0x78, 0xde,
	0xf3, 0xc8, 0xb0, 0xb8, 0xe7, 0x34, 0x65, 0xf1, 0x2a, 0x2f, 0x32, 0x99, 0xa1, 0x49, 0xbe, 0xf3,
	0x9e, 0x83, 0x73, 0x23, 0xa9, 0x64, 0x37, 0x82, 0xe6, 0xe5, 0xc7, 0x4c, 0x22, 0x17, 0xa6, 0x55,
	0x12, 0x63, 0xe3, 0xc4, 0x38, 0x5d, 0x10, 0x55, 0x22, 0x04, 0xb3, 0x98, 0x4a, 0x8a, 0x27, 0x9a,
	0xd2, 0xb5, 0xf7, 0xcd, 0x00, 0xd0, 0xcf, 0xf9, 0x42, 0x16, 0x7b, 0x74, 0x0c, 0xb6, 0xa0, 0x29,
	0x2b, 0x73, 0x1a, 0xb1, 0xe6, 0xd1, 0x9e, 0x40, 0x4f, 0x00, 0xee, 0x69, 0x9a, 0xf0, 0xfd, 0x86,
	0xa6, 0x4c, 0x8f, 0xb1, 0xc9, 0x80, 0x41, 0x4f, 0xc1, 0xa9, 0xd1, 0x2d, 0x2b, 0xca, 0x24, 0x13,
	0x78, 0xaa, 0x25, 0x63, 0x12, 0x61, 0x30, 0x73, 0xba, 0xe7, 0x19, 0x8d, 0xf1, 0x4c, 0xbf, 0xa1,
	0x85, 0xde, 0x77, 0x03, 0x0e, 0xc3, 0x82, 0x8a, 0x92, 0x46, 0x52, 0x29, 0xff, 0xed, 0x36, 0x5f,
	0x0d, 0xb0, 0xae, 0x8b, 0x2c, 0xcf, 0x4a, 0xca, 0x55, 0x76, 0x55, 0x1b, 0xa7, 0x4d, 0x74, 0xad,
	0xb8, 0x94, 0xf5, 0x79, 0xaa, 0x5a, 0xad, 0x1c, 0x15, 0x8c, 0x4a, 0x16, 0xaf, 0xa5, 0x1e, 0x38,
	0x27, 0x3d, 0x81, 0xce, 0x61, 0x21, 0x7b, 0x7f, 0x25, 0x3e, 0x38, 0x99, 0x9e, 0x1e, 0x9e, 0xfd,
	0xbf, 0xca, 0x77, 0xab, 0x81, 0x6f, 0x32, 0x12, 0x79, 0x3f, 0x26, 0x00, 0x81, 0x88, 0xd9, 0xe7,
	0xb7, 0x09, 0xe3, 0x31, 0x7a, 0x09, 0xcb, 0xa4, 0x43, 0xe1, 0x3e, 0xaf, 0x93, 0x59, 0x9e, 0x21,
	0x35, 0x25, 0x18, 0x75, 0xc8, 0x03, 0xa5, 0xda, 0x58, 0xf4, 0x61, 0xe9, 0x1a, 0x3d, 0x06, 0x8b,
	0x0a, 0xca, 0xf7, 0x5f, 0x58, 0xd1, 0x24, 0xd4, 0x61, 0x15, 0x0e, 0x13, 0x74, 0xc7, 0x59, 0x1d,
	0x8e, 0x45, 0x5a, 0x88, 0x1e, 0xc1, 0x81, 0x9e, 0x8d, 0x0f, 0x34, 0x5f, 0x03, 0xc5, 0x96, 0x32,
	0x2b, 0x18, 0x9e, 0xd7, 0xac, 0x06, 0x68, 0x05, 0x28, 0x11, 0x11, 0xaf, 0x62, 0x16, 0xb2, 0x22,
	0xbd, 0x65, 0x91, 0xcc, 0x8a, 0x12, 0x9b, 0x5a, 0xf2, 0x87, 0x0e, 0xf2, 0x60, 0xd1, 0xb0, 0x81,
	0x58, 0x73, 0x8e, 0x2d, 0xad, 0x1c, 0x71, 0x2a, 0xe7, 0x38, 0x8b, 0x6e, 0x29, 0xaf, 0x58, 0x89,
	0x6d, 0x2d, 0xe8, 0x09, 0xaf, 0x02, 0x47, 0x27, 0x71, 0x91, 0x45, 0x55, 0xca, 0x84, 0xec, 0x8c,
	0x1b, 0x03, 0xe3, 0x18, 0xcc, 0x98, 0xdd, 0xd3, 0x8a, 0x4b, 0x9d, 0x87, 0x45, 0x5a, 0x88, 0x5e,
	0x0c, 0x23, 0xbe, 0x4c, 0x4a, 0x89, 0xa7, 0xfa, 0x43, 0x2d, 0xc7, 0x11, 0x93, 0x07, 0x2a, 0xef,
	0xa7, 0x01, 0xb6, 0x6e, 0x5f, 0x35, 0xa7, 0xf0, 0x97, 0xeb, 0xc5, 0x60, 0x7e, 0x6a, 0xee, 0x52,
	0xbd, 0xdd, 0x21, 0x2d, 0x1c, 0x86, 0x3e, 0x1d, 0x87, 0x3e, 0xd8, 0x78, 0x36, 0xde, 0x78, 0x74,
	0x76, 0xea, 0x93, 0xcc, 0x86, 0x67, 0x77, 0x0c, 0x76, 0x95, 0xc7, 0x4d, 0x77, 0x5e, 0x77, 0x3b,
	0x02, 0xbd, 0x82, 0xa3, 0x64, 0x18, 0x96, 0x36, 0x6c, 0x6a, 0xc3, 0x47, 0x9d, 0xe1, 0xb6, 0x49,
	0x7e, 0xd7, 0x7a, 0x17, 0x4d, 0xda, 0xca, 0xb5, 0x22, 0xd0, 0x39, 0x38, 0xc9, 0x90, 0xc0, 0x86,
	0x9e, 0xe6, 0x74, 0xd3, 0x54, 0x83, 0x8c, 0x35, 0xcf, 0x08, 0x2c, 0xc7, 0xd7, 0x8b, 0x2c, 0x98,
	0x85, 0xfe, 0x87, 0xd0, 0xfd, 0x0f, 0x1d, 0x82, 0xb9, 0x79, 0x7f, 0xe5, 0x93, 0xe0, 0x8d, 0x6b,
	0x28, 0xf0, 0x7a, 0xbb, 0xbd, 0xf4, 0xd7, 0x1b, 0x77, 0x82, 0x1c, 0xb0, 0xdf, 0xf9, 0xdb, 0xbb,
	0xeb, 0x6d, 0xb0, 0x09, 0xdd, 0xa9, 0x82, 0x17, 0xeb, 0xd0, 0xbf, 0x0b, 0x83, 0x2b, 0xdf, 0x9d,
	0xed, 0xe6, 0xfa, 0xff, 0x78, 0xfe, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xfb, 0xb1, 0xbf, 0x2f,
	0x05, 0x00, 0x00,
}
