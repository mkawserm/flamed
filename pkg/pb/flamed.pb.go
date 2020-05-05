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
	FlameAction_APPEND FlameAction_FlameActionType = 3
)

var FlameAction_FlameActionType_name = map[int32]string{
	0: "CREATE",
	1: "UPDATE",
	2: "DELETE",
	3: "APPEND",
}

var FlameAction_FlameActionType_value = map[string]int32{
	"CREATE": 0,
	"UPDATE": 1,
	"DELETE": 2,
	"APPEND": 3,
}

func (x FlameAction_FlameActionType) String() string {
	return proto.EnumName(FlameAction_FlameActionType_name, int32(x))
}

func (FlameAction_FlameActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{1, 0}
}

type FlameIndexField_FlameIndexFieldType int32

const (
	FlameIndexField_TEXT      FlameIndexField_FlameIndexFieldType = 0
	FlameIndexField_NUMERIC   FlameIndexField_FlameIndexFieldType = 1
	FlameIndexField_BOOLEAN   FlameIndexField_FlameIndexFieldType = 2
	FlameIndexField_GEO_POINT FlameIndexField_FlameIndexFieldType = 3
	FlameIndexField_DATE_TIME FlameIndexField_FlameIndexFieldType = 4
)

var FlameIndexField_FlameIndexFieldType_name = map[int32]string{
	0: "TEXT",
	1: "NUMERIC",
	2: "BOOLEAN",
	3: "GEO_POINT",
	4: "DATE_TIME",
}

var FlameIndexField_FlameIndexFieldType_value = map[string]int32{
	"TEXT":      0,
	"NUMERIC":   1,
	"BOOLEAN":   2,
	"GEO_POINT": 3,
	"DATE_TIME": 4,
}

func (x FlameIndexField_FlameIndexFieldType) String() string {
	return proto.EnumName(FlameIndexField_FlameIndexFieldType_name, int32(x))
}

func (FlameIndexField_FlameIndexFieldType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{6, 0}
}

type FlameUser_FlameUserType int32

const (
	FlameUser_SUPER_USER  FlameUser_FlameUserType = 0
	FlameUser_NORMAL_USER FlameUser_FlameUserType = 1
)

var FlameUser_FlameUserType_name = map[int32]string{
	0: "SUPER_USER",
	1: "NORMAL_USER",
}

var FlameUser_FlameUserType_value = map[string]int32{
	"SUPER_USER":  0,
	"NORMAL_USER": 1,
}

func (x FlameUser_FlameUserType) String() string {
	return proto.EnumName(FlameUser_FlameUserType_name, int32(x))
}

func (FlameUser_FlameUserType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{9, 0}
}

type FlameProposal_FlameProposalType int32

const (
	FlameProposal_BATCH_ACTION          FlameProposal_FlameProposalType = 0
	FlameProposal_CREATE_INDEX_META     FlameProposal_FlameProposalType = 1
	FlameProposal_UPDATE_INDEX_META     FlameProposal_FlameProposalType = 2
	FlameProposal_DELETE_INDEX_META     FlameProposal_FlameProposalType = 3
	FlameProposal_CREATE_USER           FlameProposal_FlameProposalType = 4
	FlameProposal_UPDATE_USER           FlameProposal_FlameProposalType = 5
	FlameProposal_DELETE_USER           FlameProposal_FlameProposalType = 6
	FlameProposal_CREATE_ACCESS_CONTROL FlameProposal_FlameProposalType = 7
	FlameProposal_UPDATE_ACCESS_CONTROL FlameProposal_FlameProposalType = 8
	FlameProposal_DELETE_ACCESS_CONTROL FlameProposal_FlameProposalType = 9
)

var FlameProposal_FlameProposalType_name = map[int32]string{
	0: "BATCH_ACTION",
	1: "CREATE_INDEX_META",
	2: "UPDATE_INDEX_META",
	3: "DELETE_INDEX_META",
	4: "CREATE_USER",
	5: "UPDATE_USER",
	6: "DELETE_USER",
	7: "CREATE_ACCESS_CONTROL",
	8: "UPDATE_ACCESS_CONTROL",
	9: "DELETE_ACCESS_CONTROL",
}

var FlameProposal_FlameProposalType_value = map[string]int32{
	"BATCH_ACTION":          0,
	"CREATE_INDEX_META":     1,
	"UPDATE_INDEX_META":     2,
	"DELETE_INDEX_META":     3,
	"CREATE_USER":           4,
	"UPDATE_USER":           5,
	"DELETE_USER":           6,
	"CREATE_ACCESS_CONTROL": 7,
	"UPDATE_ACCESS_CONTROL": 8,
	"DELETE_ACCESS_CONTROL": 9,
}

func (x FlameProposal_FlameProposalType) String() string {
	return proto.EnumName(FlameProposal_FlameProposalType_name, int32(x))
}

func (FlameProposal_FlameProposalType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{11, 0}
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

type FlameBatchAction struct {
	FlameActionList      []*FlameAction `protobuf:"bytes,1,rep,name=flameActionList,proto3" json:"flameActionList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FlameBatchAction) Reset()         { *m = FlameBatchAction{} }
func (m *FlameBatchAction) String() string { return proto.CompactTextString(m) }
func (*FlameBatchAction) ProtoMessage()    {}
func (*FlameBatchAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{2}
}

func (m *FlameBatchAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameBatchAction.Unmarshal(m, b)
}
func (m *FlameBatchAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameBatchAction.Marshal(b, m, deterministic)
}
func (m *FlameBatchAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameBatchAction.Merge(m, src)
}
func (m *FlameBatchAction) XXX_Size() int {
	return xxx_messageInfo_FlameBatchAction.Size(m)
}
func (m *FlameBatchAction) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameBatchAction.DiscardUnknown(m)
}

var xxx_messageInfo_FlameBatchAction proto.InternalMessageInfo

func (m *FlameBatchAction) GetFlameActionList() []*FlameAction {
	if m != nil {
		return m.FlameActionList
	}
	return nil
}

type FlameBatchRead struct {
	FlameEntryList       []*FlameEntry `protobuf:"bytes,1,rep,name=flameEntryList,proto3" json:"flameEntryList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FlameBatchRead) Reset()         { *m = FlameBatchRead{} }
func (m *FlameBatchRead) String() string { return proto.CompactTextString(m) }
func (*FlameBatchRead) ProtoMessage()    {}
func (*FlameBatchRead) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{3}
}

func (m *FlameBatchRead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameBatchRead.Unmarshal(m, b)
}
func (m *FlameBatchRead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameBatchRead.Marshal(b, m, deterministic)
}
func (m *FlameBatchRead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameBatchRead.Merge(m, src)
}
func (m *FlameBatchRead) XXX_Size() int {
	return xxx_messageInfo_FlameBatchRead.Size(m)
}
func (m *FlameBatchRead) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameBatchRead.DiscardUnknown(m)
}

var xxx_messageInfo_FlameBatchRead proto.InternalMessageInfo

func (m *FlameBatchRead) GetFlameEntryList() []*FlameEntry {
	if m != nil {
		return m.FlameEntryList
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
	return fileDescriptor_388b6a0687b80922, []int{4}
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
	return fileDescriptor_388b6a0687b80922, []int{5}
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

type FlameIndexField struct {
	FlameIndexFieldType  FlameIndexField_FlameIndexFieldType `protobuf:"varint,1,opt,name=flameIndexFieldType,proto3,enum=pb.FlameIndexField_FlameIndexFieldType" json:"flameIndexFieldType,omitempty"`
	Name                 string                              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Analyzer             string                              `protobuf:"bytes,3,opt,name=analyzer,proto3" json:"analyzer,omitempty"`
	Enabled              bool                                `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Index                bool                                `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
	Store                bool                                `protobuf:"varint,6,opt,name=store,proto3" json:"store,omitempty"`
	IncludeTermVectors   bool                                `protobuf:"varint,7,opt,name=includeTermVectors,proto3" json:"includeTermVectors,omitempty"`
	IncludeInAll         bool                                `protobuf:"varint,8,opt,name=includeInAll,proto3" json:"includeInAll,omitempty"`
	DocValues            bool                                `protobuf:"varint,9,opt,name=docValues,proto3" json:"docValues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *FlameIndexField) Reset()         { *m = FlameIndexField{} }
func (m *FlameIndexField) String() string { return proto.CompactTextString(m) }
func (*FlameIndexField) ProtoMessage()    {}
func (*FlameIndexField) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{6}
}

func (m *FlameIndexField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameIndexField.Unmarshal(m, b)
}
func (m *FlameIndexField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameIndexField.Marshal(b, m, deterministic)
}
func (m *FlameIndexField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameIndexField.Merge(m, src)
}
func (m *FlameIndexField) XXX_Size() int {
	return xxx_messageInfo_FlameIndexField.Size(m)
}
func (m *FlameIndexField) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameIndexField.DiscardUnknown(m)
}

var xxx_messageInfo_FlameIndexField proto.InternalMessageInfo

func (m *FlameIndexField) GetFlameIndexFieldType() FlameIndexField_FlameIndexFieldType {
	if m != nil {
		return m.FlameIndexFieldType
	}
	return FlameIndexField_TEXT
}

func (m *FlameIndexField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FlameIndexField) GetAnalyzer() string {
	if m != nil {
		return m.Analyzer
	}
	return ""
}

func (m *FlameIndexField) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *FlameIndexField) GetIndex() bool {
	if m != nil {
		return m.Index
	}
	return false
}

func (m *FlameIndexField) GetStore() bool {
	if m != nil {
		return m.Store
	}
	return false
}

func (m *FlameIndexField) GetIncludeTermVectors() bool {
	if m != nil {
		return m.IncludeTermVectors
	}
	return false
}

func (m *FlameIndexField) GetIncludeInAll() bool {
	if m != nil {
		return m.IncludeInAll
	}
	return false
}

func (m *FlameIndexField) GetDocValues() bool {
	if m != nil {
		return m.DocValues
	}
	return false
}

type FlameIndexDocument struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Default              bool               `protobuf:"varint,2,opt,name=default,proto3" json:"default,omitempty"`
	FlameIndexFieldList  []*FlameIndexField `protobuf:"bytes,3,rep,name=flameIndexFieldList,proto3" json:"flameIndexFieldList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FlameIndexDocument) Reset()         { *m = FlameIndexDocument{} }
func (m *FlameIndexDocument) String() string { return proto.CompactTextString(m) }
func (*FlameIndexDocument) ProtoMessage()    {}
func (*FlameIndexDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{7}
}

func (m *FlameIndexDocument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameIndexDocument.Unmarshal(m, b)
}
func (m *FlameIndexDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameIndexDocument.Marshal(b, m, deterministic)
}
func (m *FlameIndexDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameIndexDocument.Merge(m, src)
}
func (m *FlameIndexDocument) XXX_Size() int {
	return xxx_messageInfo_FlameIndexDocument.Size(m)
}
func (m *FlameIndexDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameIndexDocument.DiscardUnknown(m)
}

var xxx_messageInfo_FlameIndexDocument proto.InternalMessageInfo

func (m *FlameIndexDocument) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FlameIndexDocument) GetDefault() bool {
	if m != nil {
		return m.Default
	}
	return false
}

func (m *FlameIndexDocument) GetFlameIndexFieldList() []*FlameIndexField {
	if m != nil {
		return m.FlameIndexFieldList
	}
	return nil
}

type FlameIndexMeta struct {
	Namespace              []byte                `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Version                uint32                `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Enabled                bool                  `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Default                bool                  `protobuf:"varint,4,opt,name=default,proto3" json:"default,omitempty"`
	CreatedAt              uint64                `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt              uint64                `protobuf:"varint,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	FlameIndexDocumentList []*FlameIndexDocument `protobuf:"bytes,7,rep,name=flameIndexDocumentList,proto3" json:"flameIndexDocumentList,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}              `json:"-"`
	XXX_unrecognized       []byte                `json:"-"`
	XXX_sizecache          int32                 `json:"-"`
}

func (m *FlameIndexMeta) Reset()         { *m = FlameIndexMeta{} }
func (m *FlameIndexMeta) String() string { return proto.CompactTextString(m) }
func (*FlameIndexMeta) ProtoMessage()    {}
func (*FlameIndexMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{8}
}

func (m *FlameIndexMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameIndexMeta.Unmarshal(m, b)
}
func (m *FlameIndexMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameIndexMeta.Marshal(b, m, deterministic)
}
func (m *FlameIndexMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameIndexMeta.Merge(m, src)
}
func (m *FlameIndexMeta) XXX_Size() int {
	return xxx_messageInfo_FlameIndexMeta.Size(m)
}
func (m *FlameIndexMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameIndexMeta.DiscardUnknown(m)
}

var xxx_messageInfo_FlameIndexMeta proto.InternalMessageInfo

func (m *FlameIndexMeta) GetNamespace() []byte {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *FlameIndexMeta) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *FlameIndexMeta) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *FlameIndexMeta) GetDefault() bool {
	if m != nil {
		return m.Default
	}
	return false
}

func (m *FlameIndexMeta) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *FlameIndexMeta) GetUpdatedAt() uint64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *FlameIndexMeta) GetFlameIndexDocumentList() []*FlameIndexDocument {
	if m != nil {
		return m.FlameIndexDocumentList
	}
	return nil
}

type FlameUser struct {
	FlameUserType        FlameUser_FlameUserType `protobuf:"varint,1,opt,name=flameUserType,proto3,enum=pb.FlameUser_FlameUserType" json:"flameUserType,omitempty"`
	Roles                string                  `protobuf:"bytes,2,opt,name=roles,proto3" json:"roles,omitempty"`
	Username             string                  `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password             string                  `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	CreatedAt            uint64                  `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            uint64                  `protobuf:"varint,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	UserData             []byte                  `protobuf:"bytes,7,opt,name=userData,proto3" json:"userData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FlameUser) Reset()         { *m = FlameUser{} }
func (m *FlameUser) String() string { return proto.CompactTextString(m) }
func (*FlameUser) ProtoMessage()    {}
func (*FlameUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{9}
}

func (m *FlameUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameUser.Unmarshal(m, b)
}
func (m *FlameUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameUser.Marshal(b, m, deterministic)
}
func (m *FlameUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameUser.Merge(m, src)
}
func (m *FlameUser) XXX_Size() int {
	return xxx_messageInfo_FlameUser.Size(m)
}
func (m *FlameUser) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameUser.DiscardUnknown(m)
}

var xxx_messageInfo_FlameUser proto.InternalMessageInfo

func (m *FlameUser) GetFlameUserType() FlameUser_FlameUserType {
	if m != nil {
		return m.FlameUserType
	}
	return FlameUser_SUPER_USER
}

func (m *FlameUser) GetRoles() string {
	if m != nil {
		return m.Roles
	}
	return ""
}

func (m *FlameUser) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *FlameUser) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *FlameUser) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *FlameUser) GetUpdatedAt() uint64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *FlameUser) GetUserData() []byte {
	if m != nil {
		return m.UserData
	}
	return nil
}

type FlameAccessControl struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Namespace            []byte   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Permission           []byte   `protobuf:"bytes,3,opt,name=permission,proto3" json:"permission,omitempty"`
	CreatedAt            uint64   `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            uint64   `protobuf:"varint,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlameAccessControl) Reset()         { *m = FlameAccessControl{} }
func (m *FlameAccessControl) String() string { return proto.CompactTextString(m) }
func (*FlameAccessControl) ProtoMessage()    {}
func (*FlameAccessControl) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{10}
}

func (m *FlameAccessControl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameAccessControl.Unmarshal(m, b)
}
func (m *FlameAccessControl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameAccessControl.Marshal(b, m, deterministic)
}
func (m *FlameAccessControl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameAccessControl.Merge(m, src)
}
func (m *FlameAccessControl) XXX_Size() int {
	return xxx_messageInfo_FlameAccessControl.Size(m)
}
func (m *FlameAccessControl) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameAccessControl.DiscardUnknown(m)
}

var xxx_messageInfo_FlameAccessControl proto.InternalMessageInfo

func (m *FlameAccessControl) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *FlameAccessControl) GetNamespace() []byte {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *FlameAccessControl) GetPermission() []byte {
	if m != nil {
		return m.Permission
	}
	return nil
}

func (m *FlameAccessControl) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *FlameAccessControl) GetUpdatedAt() uint64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type FlameProposal struct {
	FlameProposalType    FlameProposal_FlameProposalType `protobuf:"varint,1,opt,name=flameProposalType,proto3,enum=pb.FlameProposal_FlameProposalType" json:"flameProposalType,omitempty"`
	FlameProposalData    []byte                          `protobuf:"bytes,2,opt,name=flameProposalData,proto3" json:"flameProposalData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *FlameProposal) Reset()         { *m = FlameProposal{} }
func (m *FlameProposal) String() string { return proto.CompactTextString(m) }
func (*FlameProposal) ProtoMessage()    {}
func (*FlameProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{11}
}

func (m *FlameProposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameProposal.Unmarshal(m, b)
}
func (m *FlameProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameProposal.Marshal(b, m, deterministic)
}
func (m *FlameProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameProposal.Merge(m, src)
}
func (m *FlameProposal) XXX_Size() int {
	return xxx_messageInfo_FlameProposal.Size(m)
}
func (m *FlameProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameProposal.DiscardUnknown(m)
}

var xxx_messageInfo_FlameProposal proto.InternalMessageInfo

func (m *FlameProposal) GetFlameProposalType() FlameProposal_FlameProposalType {
	if m != nil {
		return m.FlameProposalType
	}
	return FlameProposal_BATCH_ACTION
}

func (m *FlameProposal) GetFlameProposalData() []byte {
	if m != nil {
		return m.FlameProposalData
	}
	return nil
}

type FlameQuery struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlameQuery) Reset()         { *m = FlameQuery{} }
func (m *FlameQuery) String() string { return proto.CompactTextString(m) }
func (*FlameQuery) ProtoMessage()    {}
func (*FlameQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{12}
}

func (m *FlameQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameQuery.Unmarshal(m, b)
}
func (m *FlameQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameQuery.Marshal(b, m, deterministic)
}
func (m *FlameQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameQuery.Merge(m, src)
}
func (m *FlameQuery) XXX_Size() int {
	return xxx_messageInfo_FlameQuery.Size(m)
}
func (m *FlameQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameQuery.DiscardUnknown(m)
}

var xxx_messageInfo_FlameQuery proto.InternalMessageInfo

type FlameResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlameResponse) Reset()         { *m = FlameResponse{} }
func (m *FlameResponse) String() string { return proto.CompactTextString(m) }
func (*FlameResponse) ProtoMessage()    {}
func (*FlameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{13}
}

func (m *FlameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameResponse.Unmarshal(m, b)
}
func (m *FlameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameResponse.Marshal(b, m, deterministic)
}
func (m *FlameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameResponse.Merge(m, src)
}
func (m *FlameResponse) XXX_Size() int {
	return xxx_messageInfo_FlameResponse.Size(m)
}
func (m *FlameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FlameResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("pb.FlameAction_FlameActionType", FlameAction_FlameActionType_name, FlameAction_FlameActionType_value)
	proto.RegisterEnum("pb.FlameIndexField_FlameIndexFieldType", FlameIndexField_FlameIndexFieldType_name, FlameIndexField_FlameIndexFieldType_value)
	proto.RegisterEnum("pb.FlameUser_FlameUserType", FlameUser_FlameUserType_name, FlameUser_FlameUserType_value)
	proto.RegisterEnum("pb.FlameProposal_FlameProposalType", FlameProposal_FlameProposalType_name, FlameProposal_FlameProposalType_value)
	proto.RegisterType((*FlameEntry)(nil), "pb.FlameEntry")
	proto.RegisterType((*FlameAction)(nil), "pb.FlameAction")
	proto.RegisterType((*FlameBatchAction)(nil), "pb.FlameBatchAction")
	proto.RegisterType((*FlameBatchRead)(nil), "pb.FlameBatchRead")
	proto.RegisterType((*FlameSnapshotEntry)(nil), "pb.FlameSnapshotEntry")
	proto.RegisterType((*FlameSnapshot)(nil), "pb.FlameSnapshot")
	proto.RegisterType((*FlameIndexField)(nil), "pb.FlameIndexField")
	proto.RegisterType((*FlameIndexDocument)(nil), "pb.FlameIndexDocument")
	proto.RegisterType((*FlameIndexMeta)(nil), "pb.FlameIndexMeta")
	proto.RegisterType((*FlameUser)(nil), "pb.FlameUser")
	proto.RegisterType((*FlameAccessControl)(nil), "pb.FlameAccessControl")
	proto.RegisterType((*FlameProposal)(nil), "pb.FlameProposal")
	proto.RegisterType((*FlameQuery)(nil), "pb.FlameQuery")
	proto.RegisterType((*FlameResponse)(nil), "pb.FlameResponse")
}

func init() { proto.RegisterFile("flamed.proto", fileDescriptor_388b6a0687b80922) }

var fileDescriptor_388b6a0687b80922 = []byte{
	// 980 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xcb, 0x8e, 0xe3, 0x44,
	0x14, 0x6d, 0xc7, 0xe9, 0x3c, 0x6e, 0xd2, 0x89, 0xbb, 0x9a, 0x69, 0x19, 0x18, 0x41, 0xab, 0x58,
	0xd0, 0x0b, 0x14, 0xa1, 0x46, 0x42, 0x82, 0x9d, 0x3b, 0xf1, 0x30, 0x91, 0xf2, 0x9a, 0x8a, 0x33,
	0x0c, 0xab, 0xc8, 0x6d, 0x57, 0xe8, 0x08, 0xc7, 0xb6, 0x6c, 0x67, 0x20, 0xec, 0xd8, 0xc1, 0x92,
	0x9f, 0x60, 0x81, 0xc4, 0x97, 0xf0, 0x2f, 0x7c, 0x03, 0xaa, 0xeb, 0x8a, 0x63, 0x3b, 0x69, 0x58,
	0xcc, 0xee, 0x9e, 0x73, 0x1f, 0xae, 0xfb, 0xaa, 0x32, 0xb4, 0x57, 0x9e, 0xbd, 0xe1, 0x6e, 0x2f,
	0x8c, 0x82, 0x24, 0x20, 0x95, 0xf0, 0x81, 0x32, 0x80, 0x17, 0x82, 0x33, 0xfd, 0x24, 0xda, 0x91,
	0xe7, 0xd0, 0xf4, 0xed, 0x0d, 0x8f, 0x43, 0xdb, 0xe1, 0xba, 0x72, 0xa3, 0xdc, 0xb6, 0xd9, 0x81,
	0x20, 0x1a, 0xa8, 0x3f, 0xf0, 0x9d, 0x5e, 0x41, 0x5e, 0x88, 0xe4, 0x3d, 0x38, 0x7f, 0x6b, 0x7b,
	0x5b, 0xae, 0xab, 0xc8, 0xa5, 0x80, 0xfe, 0xad, 0x40, 0x0b, 0x83, 0x1a, 0x4e, 0xb2, 0x0e, 0x7c,
	0x32, 0x84, 0xee, 0xea, 0x00, 0xad, 0x5d, 0x98, 0xc6, 0xee, 0xdc, 0x7d, 0xdc, 0x0b, 0x1f, 0x7a,
	0x39, 0xcb, 0xbc, 0x2c, 0xcc, 0x58, 0xd9, 0x8f, 0xf4, 0x00, 0x56, 0xd9, 0x71, 0xf1, 0x24, 0xad,
	0xbb, 0x4e, 0x16, 0x05, 0x59, 0x96, 0xb3, 0xa0, 0x06, 0x74, 0x4b, 0x31, 0x09, 0x40, 0xad, 0xcf,
	0x4c, 0xc3, 0x32, 0xb5, 0x33, 0x21, 0x2f, 0x66, 0x03, 0x21, 0x2b, 0x42, 0x1e, 0x98, 0x23, 0xd3,
	0x32, 0xb5, 0x8a, 0x90, 0x8d, 0xd9, 0xcc, 0x9c, 0x0c, 0x34, 0x95, 0x8e, 0x41, 0xc3, 0x10, 0xf7,
	0x76, 0xe2, 0x3c, 0xca, 0x8c, 0xbe, 0x2a, 0x64, 0x34, 0x5a, 0xc7, 0x89, 0xae, 0xdc, 0xa8, 0xb7,
	0xad, 0xbb, 0x6e, 0x29, 0x23, 0x56, 0xb6, 0xa3, 0x2f, 0xa1, 0x73, 0x08, 0xc7, 0xb8, 0xed, 0x92,
	0x2f, 0xa1, 0x73, 0x38, 0x71, 0x2e, 0x56, 0x39, 0xaf, 0x92, 0x15, 0xfd, 0x1a, 0x08, 0x6a, 0xe7,
	0xbe, 0x1d, 0xc6, 0x8f, 0x41, 0x92, 0xb6, 0x50, 0x03, 0x75, 0xbb, 0x76, 0x65, 0xf3, 0x84, 0x48,
	0x08, 0x54, 0x5d, 0x3b, 0xb1, 0x65, 0xdf, 0x50, 0xa6, 0xbf, 0x2b, 0x70, 0x51, 0x70, 0x26, 0x3a,
	0xd4, 0xdf, 0xf2, 0x28, 0x5e, 0x07, 0x3e, 0xfa, 0x5e, 0xb0, 0x3d, 0x24, 0xd7, 0x50, 0xf3, 0xb8,
	0xff, 0x7d, 0xf2, 0x88, 0x11, 0xaa, 0x4c, 0x22, 0x32, 0x81, 0xeb, 0xd5, 0xd1, 0xf7, 0xf1, 0xfc,
	0x2a, 0x9e, 0xff, 0x3a, 0x3b, 0x7f, 0xc1, 0x82, 0x3d, 0xe1, 0x45, 0xff, 0x52, 0x65, 0xb3, 0x86,
	0xbe, 0xcb, 0x7f, 0x7a, 0xb1, 0xe6, 0x9e, 0x4b, 0xbe, 0x83, 0xab, 0x55, 0x91, 0xca, 0x8d, 0xcf,
	0xa7, 0xd9, 0x07, 0x0e, 0xea, 0x32, 0xc6, 0x31, 0x3a, 0x15, 0x43, 0x94, 0x45, 0x8c, 0x36, 0x26,
	0xd5, 0x64, 0x28, 0x93, 0x0f, 0xa0, 0x61, 0xfb, 0xb6, 0xb7, 0xfb, 0x99, 0x47, 0x38, 0xd2, 0x4d,
	0x96, 0x61, 0x51, 0x20, 0xee, 0xdb, 0x0f, 0x1e, 0x77, 0xf5, 0xea, 0x8d, 0x72, 0xdb, 0x60, 0x7b,
	0x28, 0xb6, 0x60, 0x2d, 0x62, 0xeb, 0xe7, 0xc8, 0xa7, 0x40, 0xb0, 0x71, 0x12, 0x44, 0x5c, 0xaf,
	0xa5, 0x2c, 0x02, 0xd2, 0x03, 0xb2, 0xf6, 0x1d, 0x6f, 0xeb, 0x72, 0x8b, 0x47, 0x9b, 0xd7, 0xdc,
	0x49, 0x82, 0x28, 0xd6, 0xeb, 0x68, 0x72, 0x42, 0x43, 0x28, 0xb4, 0x25, 0x3b, 0xf4, 0x0d, 0xcf,
	0xd3, 0x1b, 0x68, 0x59, 0xe0, 0xc4, 0xd6, 0xba, 0x81, 0xf3, 0x5a, 0xec, 0x5e, 0xac, 0x37, 0xd1,
	0xe0, 0x40, 0xd0, 0x6f, 0xe1, 0xea, 0x44, 0x4d, 0x48, 0x03, 0xaa, 0x96, 0xf9, 0xc6, 0xd2, 0xce,
	0x48, 0x0b, 0xea, 0x93, 0xc5, 0xd8, 0x64, 0xc3, 0xbe, 0xa6, 0x08, 0x70, 0x3f, 0x9d, 0x8e, 0x4c,
	0x63, 0xa2, 0x55, 0xc8, 0x05, 0x34, 0xbf, 0x31, 0xa7, 0xcb, 0xd9, 0x74, 0x38, 0xb1, 0x34, 0x55,
	0x40, 0xb1, 0x2b, 0x4b, 0x6b, 0x38, 0x36, 0xb5, 0x2a, 0xfd, 0x4d, 0x91, 0x03, 0x88, 0x91, 0x07,
	0x81, 0xb3, 0xdd, 0x70, 0x3f, 0xc9, 0xea, 0xaa, 0xe4, 0xea, 0xaa, 0x43, 0xdd, 0xe5, 0x2b, 0x7b,
	0xeb, 0x25, 0x58, 0xee, 0x06, 0xdb, 0x43, 0x62, 0x1e, 0x35, 0x38, 0x37, 0x41, 0x57, 0x27, 0x1a,
	0xcc, 0x4e, 0xd9, 0xd3, 0x5f, 0x2b, 0x72, 0xad, 0x90, 0x1f, 0xf3, 0xc4, 0xfe, 0x9f, 0xbb, 0x2c,
	0x37, 0xee, 0x95, 0xe2, 0xb8, 0xe7, 0xfa, 0xac, 0x16, 0xfb, 0x9c, 0xcb, 0xa2, 0x5a, 0xcc, 0xe2,
	0x39, 0x34, 0x9d, 0x88, 0xdb, 0x09, 0x77, 0x8d, 0x04, 0xa7, 0xa0, 0xca, 0x0e, 0x84, 0xd0, 0x6e,
	0x43, 0x57, 0x6a, 0x6b, 0xa9, 0x36, 0x23, 0xb2, 0x35, 0x2a, 0x54, 0x11, 0x8b, 0x50, 0x2f, 0xad,
	0x51, 0xc1, 0x82, 0x3d, 0xe1, 0x45, 0xff, 0xa8, 0x40, 0x13, 0xcd, 0x17, 0x31, 0x8f, 0x88, 0x01,
	0x17, 0xab, 0x3d, 0xc8, 0xad, 0xce, 0x87, 0x59, 0x50, 0xa1, 0x38, 0x48, 0xb8, 0x2e, 0x45, 0x0f,
	0x31, 0xc8, 0x51, 0xe0, 0xf1, 0x58, 0x6e, 0x4a, 0x0a, 0xc4, 0xaa, 0x6c, 0x63, 0x1e, 0x61, 0xab,
	0xe5, 0xaa, 0xec, 0xb1, 0xd0, 0x85, 0x76, 0x1c, 0xff, 0x18, 0x44, 0xe9, 0xae, 0x34, 0x59, 0x86,
	0xdf, 0xa9, 0x54, 0xf2, 0x9b, 0x03, 0x71, 0x9b, 0xd5, 0xb1, 0xa3, 0x19, 0xa6, 0x9f, 0xcb, 0x0b,
	0x2d, 0x3b, 0x76, 0x07, 0x60, 0xbe, 0x98, 0x99, 0x6c, 0xb9, 0x98, 0x9b, 0x4c, 0x3b, 0x23, 0x5d,
	0x68, 0x4d, 0xa6, 0x6c, 0x6c, 0x8c, 0x52, 0x42, 0xa1, 0x7f, 0xee, 0xe7, 0xd7, 0x70, 0x1c, 0x1e,
	0xc7, 0xfd, 0xc0, 0x4f, 0xa2, 0xc0, 0x2b, 0x24, 0xa6, 0x94, 0x12, 0x2b, 0xcc, 0x54, 0xa5, 0x3c,
	0x53, 0x1f, 0x01, 0x84, 0x3c, 0xda, 0xac, 0x63, 0x1c, 0xab, 0xf4, 0x49, 0xcc, 0x31, 0xc5, 0xd4,
	0xab, 0xff, 0x99, 0xfa, 0x79, 0x29, 0x75, 0xfa, 0x8b, 0x2a, 0xf3, 0x9b, 0x45, 0x41, 0x18, 0xc4,
	0xb6, 0x47, 0x5e, 0xc1, 0xe5, 0x2a, 0x4f, 0xe4, 0xba, 0xfb, 0x49, 0xd6, 0xdd, 0xbd, 0xb2, 0x88,
	0xb0, 0xcb, 0xc7, 0xde, 0xe4, 0xb3, 0x52, 0xc8, 0xc1, 0xe1, 0xd9, 0x38, 0x56, 0xd0, 0x7f, 0x14,
	0xb8, 0x3c, 0x0a, 0x4b, 0x34, 0x68, 0xdf, 0x1b, 0x56, 0xff, 0xe5, 0xd2, 0xe8, 0x5b, 0xc3, 0xe9,
	0x44, 0x3b, 0x23, 0xcf, 0xe0, 0x32, 0x7d, 0x70, 0x97, 0xc3, 0xc9, 0xc0, 0x7c, 0xb3, 0x1c, 0x9b,
	0x96, 0xa1, 0x29, 0x82, 0x4e, 0xdf, 0xde, 0x3c, 0x5d, 0x11, 0x74, 0xfa, 0x0c, 0xe7, 0x69, 0x55,
	0x74, 0x4f, 0x06, 0xc1, 0xee, 0x55, 0x05, 0x21, 0xdd, 0x91, 0x38, 0x17, 0x84, 0x74, 0x44, 0xa2,
	0x46, 0xde, 0x87, 0x67, 0xd2, 0xc5, 0xe8, 0xf7, 0xcd, 0xf9, 0x7c, 0xd9, 0x9f, 0x4e, 0x2c, 0x36,
	0x1d, 0x69, 0x75, 0xa1, 0x92, 0xce, 0x25, 0x55, 0x43, 0xa8, 0x64, 0x98, 0x92, 0xaa, 0x49, 0xdb,
	0xf2, 0x5f, 0xe9, 0xd5, 0x96, 0x47, 0x3b, 0xda, 0x95, 0x0d, 0x61, 0x3c, 0x0e, 0x03, 0x3f, 0xe6,
	0x0f, 0x35, 0xfc, 0xab, 0xfa, 0xe2, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x3a, 0x86, 0x8c,
	0x65, 0x09, 0x00, 0x00,
}
