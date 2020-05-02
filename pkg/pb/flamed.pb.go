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

type FlameIndexField_IndexFieldType int32

const (
	FlameIndexField_TEXT      FlameIndexField_IndexFieldType = 0
	FlameIndexField_NUMERIC   FlameIndexField_IndexFieldType = 1
	FlameIndexField_BOOLEAN   FlameIndexField_IndexFieldType = 2
	FlameIndexField_GEO_POINT FlameIndexField_IndexFieldType = 3
	FlameIndexField_DATE_TIME FlameIndexField_IndexFieldType = 4
)

var FlameIndexField_IndexFieldType_name = map[int32]string{
	0: "TEXT",
	1: "NUMERIC",
	2: "BOOLEAN",
	3: "GEO_POINT",
	4: "DATE_TIME",
}

var FlameIndexField_IndexFieldType_value = map[string]int32{
	"TEXT":      0,
	"NUMERIC":   1,
	"BOOLEAN":   2,
	"GEO_POINT": 3,
	"DATE_TIME": 4,
}

func (x FlameIndexField_IndexFieldType) String() string {
	return proto.EnumName(FlameIndexField_IndexFieldType_name, int32(x))
}

func (FlameIndexField_IndexFieldType) EnumDescriptor() ([]byte, []int) {
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
	return fileDescriptor_388b6a0687b80922, []int{8, 0}
}

type FlameProposal_FlameProposalType int32

const (
	FlameProposal_BATCH_ACTION          FlameProposal_FlameProposalType = 0
	FlameProposal_ADD_INDEX_META        FlameProposal_FlameProposalType = 1
	FlameProposal_UPDATE_INDEX_META     FlameProposal_FlameProposalType = 2
	FlameProposal_DELETE_INDEX_META     FlameProposal_FlameProposalType = 3
	FlameProposal_ADD_USER              FlameProposal_FlameProposalType = 4
	FlameProposal_UPDATE_USER           FlameProposal_FlameProposalType = 5
	FlameProposal_DELETE_USER           FlameProposal_FlameProposalType = 6
	FlameProposal_ADD_ACCESS_CONTROL    FlameProposal_FlameProposalType = 7
	FlameProposal_UPDATE_ACCESS_CONTROL FlameProposal_FlameProposalType = 8
	FlameProposal_DELETE_ACCESS_CONTROL FlameProposal_FlameProposalType = 9
)

var FlameProposal_FlameProposalType_name = map[int32]string{
	0: "BATCH_ACTION",
	1: "ADD_INDEX_META",
	2: "UPDATE_INDEX_META",
	3: "DELETE_INDEX_META",
	4: "ADD_USER",
	5: "UPDATE_USER",
	6: "DELETE_USER",
	7: "ADD_ACCESS_CONTROL",
	8: "UPDATE_ACCESS_CONTROL",
	9: "DELETE_ACCESS_CONTROL",
}

var FlameProposal_FlameProposalType_value = map[string]int32{
	"BATCH_ACTION":          0,
	"ADD_INDEX_META":        1,
	"UPDATE_INDEX_META":     2,
	"DELETE_INDEX_META":     3,
	"ADD_USER":              4,
	"UPDATE_USER":           5,
	"DELETE_USER":           6,
	"ADD_ACCESS_CONTROL":    7,
	"UPDATE_ACCESS_CONTROL": 8,
	"DELETE_ACCESS_CONTROL": 9,
}

func (x FlameProposal_FlameProposalType) String() string {
	return proto.EnumName(FlameProposal_FlameProposalType_name, int32(x))
}

func (FlameProposal_FlameProposalType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{10, 0}
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
	IndexFieldType       FlameIndexField_IndexFieldType `protobuf:"varint,1,opt,name=indexFieldType,proto3,enum=pb.FlameIndexField_IndexFieldType" json:"indexFieldType,omitempty"`
	Name                 string                         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Analyzer             string                         `protobuf:"bytes,3,opt,name=analyzer,proto3" json:"analyzer,omitempty"`
	Enabled              bool                           `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Index                bool                           `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
	Store                bool                           `protobuf:"varint,6,opt,name=store,proto3" json:"store,omitempty"`
	IncludeTermVectors   bool                           `protobuf:"varint,7,opt,name=includeTermVectors,proto3" json:"includeTermVectors,omitempty"`
	IncludeInAll         bool                           `protobuf:"varint,8,opt,name=includeInAll,proto3" json:"includeInAll,omitempty"`
	DocValues            bool                           `protobuf:"varint,9,opt,name=docValues,proto3" json:"docValues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
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

func (m *FlameIndexField) GetIndexFieldType() FlameIndexField_IndexFieldType {
	if m != nil {
		return m.IndexFieldType
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

type FlameIndexMeta struct {
	Namespace            []byte             `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Version              uint32             `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	CreatedAt            uint64             `protobuf:"varint,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            uint64             `protobuf:"varint,4,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	FlameIndexFieldList  []*FlameIndexField `protobuf:"bytes,5,rep,name=flameIndexFieldList,proto3" json:"flameIndexFieldList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FlameIndexMeta) Reset()         { *m = FlameIndexMeta{} }
func (m *FlameIndexMeta) String() string { return proto.CompactTextString(m) }
func (*FlameIndexMeta) ProtoMessage()    {}
func (*FlameIndexMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{7}
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

func (m *FlameIndexMeta) GetFlameIndexFieldList() []*FlameIndexField {
	if m != nil {
		return m.FlameIndexFieldList
	}
	return nil
}

type FlameUser struct {
	FlameUserType        FlameUser_FlameUserType `protobuf:"varint,1,opt,name=flameUserType,proto3,enum=pb.FlameUser_FlameUserType" json:"flameUserType,omitempty"`
	Username             string                  `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string                  `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	CreatedAt            uint64                  `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            uint64                  `protobuf:"varint,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	UserData             []byte                  `protobuf:"bytes,6,opt,name=userData,proto3" json:"userData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FlameUser) Reset()         { *m = FlameUser{} }
func (m *FlameUser) String() string { return proto.CompactTextString(m) }
func (*FlameUser) ProtoMessage()    {}
func (*FlameUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_388b6a0687b80922, []int{8}
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
	return fileDescriptor_388b6a0687b80922, []int{9}
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
	return fileDescriptor_388b6a0687b80922, []int{10}
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

func init() {
	proto.RegisterEnum("pb.FlameAction_FlameActionType", FlameAction_FlameActionType_name, FlameAction_FlameActionType_value)
	proto.RegisterEnum("pb.FlameIndexField_IndexFieldType", FlameIndexField_IndexFieldType_name, FlameIndexField_IndexFieldType_value)
	proto.RegisterEnum("pb.FlameUser_FlameUserType", FlameUser_FlameUserType_name, FlameUser_FlameUserType_value)
	proto.RegisterEnum("pb.FlameProposal_FlameProposalType", FlameProposal_FlameProposalType_name, FlameProposal_FlameProposalType_value)
	proto.RegisterType((*FlameEntry)(nil), "pb.FlameEntry")
	proto.RegisterType((*FlameAction)(nil), "pb.FlameAction")
	proto.RegisterType((*FlameBatchAction)(nil), "pb.FlameBatchAction")
	proto.RegisterType((*FlameBatchRead)(nil), "pb.FlameBatchRead")
	proto.RegisterType((*FlameSnapshotEntry)(nil), "pb.FlameSnapshotEntry")
	proto.RegisterType((*FlameSnapshot)(nil), "pb.FlameSnapshot")
	proto.RegisterType((*FlameIndexField)(nil), "pb.FlameIndexField")
	proto.RegisterType((*FlameIndexMeta)(nil), "pb.FlameIndexMeta")
	proto.RegisterType((*FlameUser)(nil), "pb.FlameUser")
	proto.RegisterType((*FlameAccessControl)(nil), "pb.FlameAccessControl")
	proto.RegisterType((*FlameProposal)(nil), "pb.FlameProposal")
}

func init() { proto.RegisterFile("flamed.proto", fileDescriptor_388b6a0687b80922) }

var fileDescriptor_388b6a0687b80922 = []byte{
	// 899 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xcd, 0x6e, 0xe3, 0xd4,
	0x17, 0xaf, 0x3f, 0x92, 0x26, 0x27, 0xa9, 0xeb, 0x9e, 0xf9, 0x4f, 0xe5, 0x3f, 0x8c, 0xa0, 0x32,
	0x9b, 0x2e, 0x50, 0x84, 0x8a, 0x84, 0x04, 0x3b, 0x37, 0xb9, 0xc3, 0x18, 0x25, 0x4e, 0xb8, 0x71,
	0x46, 0xb3, 0x8b, 0x5c, 0xfb, 0x86, 0x46, 0xb8, 0xb6, 0x65, 0x3b, 0x03, 0xe5, 0x29, 0x60, 0xc9,
	0x2b, 0xf0, 0x2a, 0x3c, 0x07, 0xef, 0xc0, 0x12, 0xdd, 0x6b, 0xc7, 0x5f, 0x2d, 0xb0, 0x60, 0x77,
	0x7e, 0xbf, 0xf3, 0x91, 0x7b, 0xce, 0xf9, 0x5d, 0xdf, 0xc0, 0x78, 0x17, 0x7a, 0x0f, 0x2c, 0x98,
	0x24, 0x69, 0x9c, 0xc7, 0x28, 0x27, 0x77, 0x26, 0x05, 0x78, 0xcd, 0x39, 0x12, 0xe5, 0xe9, 0x23,
	0xbe, 0x82, 0x61, 0xe4, 0x3d, 0xb0, 0x2c, 0xf1, 0x7c, 0x66, 0x48, 0x57, 0xd2, 0xf5, 0x98, 0xd6,
	0x04, 0xea, 0xa0, 0x7c, 0xcf, 0x1e, 0x0d, 0x59, 0xf0, 0xdc, 0xc4, 0xff, 0x41, 0xef, 0xbd, 0x17,
	0x1e, 0x98, 0xa1, 0x08, 0xae, 0x00, 0xe6, 0xef, 0x12, 0x8c, 0x44, 0x51, 0xcb, 0xcf, 0xf7, 0x71,
	0x84, 0x36, 0x9c, 0xef, 0x6a, 0xe8, 0x3e, 0x26, 0x45, 0x6d, 0xed, 0xe6, 0xe3, 0x49, 0x72, 0x37,
	0x69, 0x44, 0x36, 0x6d, 0x1e, 0x46, 0xbb, 0x79, 0x38, 0x01, 0xd8, 0x55, 0xc7, 0x15, 0x27, 0x19,
	0xdd, 0x68, 0x55, 0x15, 0xc1, 0xd2, 0x46, 0x84, 0x69, 0xc1, 0x79, 0xa7, 0x26, 0x02, 0xf4, 0xa7,
	0x94, 0x58, 0x2e, 0xd1, 0x4f, 0xb8, 0xbd, 0x59, 0xcd, 0xb8, 0x2d, 0x71, 0xdb, 0x5a, 0xad, 0x88,
	0x33, 0xd3, 0x65, 0x6e, 0xcf, 0xc8, 0x9c, 0xb8, 0x44, 0x57, 0xcc, 0x05, 0xe8, 0xa2, 0xc4, 0xad,
	0x97, 0xfb, 0xf7, 0x65, 0x47, 0x5f, 0xb6, 0x3a, 0x9a, 0xef, 0xb3, 0xdc, 0x90, 0xae, 0x94, 0xeb,
	0xd1, 0xcd, 0x79, 0xa7, 0x23, 0xda, 0x8d, 0x33, 0xdf, 0x80, 0x56, 0x97, 0xa3, 0xcc, 0x0b, 0xf0,
	0x0b, 0xd0, 0xea, 0x13, 0x37, 0x6a, 0x75, 0xfb, 0xea, 0x44, 0x99, 0x5f, 0x01, 0x0a, 0xef, 0x3a,
	0xf2, 0x92, 0xec, 0x3e, 0xce, 0x8b, 0x15, 0xea, 0xa0, 0x1c, 0xf6, 0x41, 0xb9, 0x3c, 0x6e, 0x22,
	0x82, 0x1a, 0x78, 0xb9, 0x57, 0xee, 0x4d, 0xd8, 0xe6, 0x2f, 0x12, 0x9c, 0xb5, 0x92, 0xd1, 0x80,
	0xd3, 0xf7, 0x2c, 0xcd, 0xf6, 0x71, 0x24, 0x72, 0xcf, 0xe8, 0x11, 0xe2, 0x25, 0xf4, 0x43, 0x16,
	0x7d, 0x97, 0xdf, 0x8b, 0x0a, 0x2a, 0x2d, 0x11, 0x3a, 0x70, 0xb9, 0x7b, 0xf2, 0xfb, 0xe2, 0xfc,
	0x8a, 0x38, 0xff, 0x65, 0x75, 0xfe, 0x56, 0x04, 0xfd, 0x9b, 0x2c, 0xf3, 0x57, 0xa5, 0x5c, 0x96,
	0x1d, 0x05, 0xec, 0xc7, 0xd7, 0x7b, 0x16, 0x06, 0xf8, 0x0d, 0x68, 0xfb, 0x0a, 0x35, 0x94, 0x63,
	0x56, 0xb5, 0xeb, 0xe0, 0x89, 0xdd, 0x8a, 0xa4, 0x9d, 0x4c, 0x3e, 0x07, 0xae, 0x65, 0xd1, 0xc5,
	0x90, 0x0a, 0x1b, 0x3f, 0x80, 0x81, 0x17, 0x79, 0xe1, 0xe3, 0x4f, 0x2c, 0x15, 0x1a, 0x1e, 0xd2,
	0x0a, 0xf3, 0x89, 0xb0, 0xc8, 0xbb, 0x0b, 0x59, 0x60, 0xa8, 0x57, 0xd2, 0xf5, 0x80, 0x1e, 0x21,
	0x97, 0xbd, 0xa8, 0x6d, 0xf4, 0x04, 0x5f, 0x00, 0xce, 0x66, 0x79, 0x9c, 0x32, 0xa3, 0x5f, 0xb0,
	0x02, 0xe0, 0x04, 0x70, 0x1f, 0xf9, 0xe1, 0x21, 0x60, 0x2e, 0x4b, 0x1f, 0xde, 0x32, 0x3f, 0x8f,
	0xd3, 0xcc, 0x38, 0x15, 0x21, 0xcf, 0x78, 0xd0, 0x84, 0x71, 0xc9, 0xda, 0x91, 0x15, 0x86, 0xc6,
	0x40, 0x44, 0xb6, 0x38, 0x7e, 0x4d, 0x83, 0xd8, 0x7f, 0xcb, 0x2f, 0x5b, 0x66, 0x0c, 0x45, 0x40,
	0x4d, 0x98, 0x14, 0xb4, 0xf6, 0x24, 0x70, 0x00, 0xaa, 0x4b, 0xde, 0xb9, 0xfa, 0x09, 0x8e, 0xe0,
	0xd4, 0xd9, 0x2c, 0x08, 0xb5, 0xa7, 0xba, 0xc4, 0xc1, 0xed, 0x72, 0x39, 0x27, 0x96, 0xa3, 0xcb,
	0x78, 0x06, 0xc3, 0xaf, 0xc9, 0x72, 0xbb, 0x5a, 0xda, 0x8e, 0xab, 0x2b, 0x1c, 0xf2, 0x7b, 0xb1,
	0x75, 0xed, 0x05, 0xd1, 0x55, 0x7e, 0xa5, 0xb5, 0x7a, 0xdc, 0x0b, 0x96, 0x7b, 0xff, 0xf2, 0xad,
	0x68, 0xc8, 0x49, 0x6e, 0xcb, 0xe9, 0x15, 0x0c, 0xfd, 0x94, 0x79, 0x39, 0x0b, 0xac, 0x5c, 0xcc,
	0x5c, 0xa5, 0x35, 0xc1, 0xbd, 0x87, 0x24, 0x28, 0xbd, 0x6a, 0xe1, 0xad, 0x08, 0x24, 0xf0, 0x62,
	0xd7, 0x5e, 0xba, 0xd0, 0x5b, 0x4f, 0xe8, 0xed, 0xc5, 0x33, 0x9a, 0xa0, 0xcf, 0xc5, 0x9b, 0x3f,
	0xcb, 0x30, 0x14, 0x81, 0x9b, 0x8c, 0xa5, 0x68, 0xc1, 0xd9, 0xee, 0x08, 0x1a, 0x12, 0xfb, 0xb0,
	0x2a, 0xc7, 0x1d, 0xb5, 0x25, 0xb4, 0xd5, 0xce, 0xe0, 0x32, 0x3a, 0x64, 0x2c, 0x6d, 0xc8, 0xab,
	0xc2, 0xdc, 0x97, 0x78, 0x59, 0xf6, 0x43, 0x9c, 0x06, 0x47, 0x89, 0x1d, 0x71, 0x7b, 0x16, 0xea,
	0x3f, 0xce, 0xa2, 0xd7, 0x9d, 0x45, 0xf9, 0x9b, 0x33, 0x7e, 0xb5, 0xfb, 0x62, 0xfc, 0x15, 0x36,
	0x3f, 0x2b, 0x6f, 0x77, 0x75, 0x40, 0x0d, 0x60, 0xbd, 0x59, 0x11, 0xba, 0xdd, 0xac, 0x09, 0xd5,
	0x4f, 0xf0, 0x1c, 0x46, 0xce, 0x92, 0x2e, 0xac, 0x79, 0x41, 0x48, 0xe6, 0x6f, 0x52, 0xf9, 0x35,
	0xb1, 0x7c, 0x9f, 0x65, 0xd9, 0x34, 0x8e, 0xf2, 0x34, 0x0e, 0x5b, 0x8d, 0x49, 0x9d, 0xc6, 0x5a,
	0x02, 0x90, 0xbb, 0x02, 0xf8, 0x08, 0x20, 0x61, 0xe9, 0xc3, 0x3e, 0x13, 0x1a, 0x28, 0xde, 0x87,
	0x06, 0xf3, 0x5f, 0x5a, 0x37, 0xff, 0x94, 0xcb, 0xfe, 0x56, 0x69, 0x9c, 0xc4, 0x99, 0x17, 0xe2,
	0xb7, 0x70, 0xb1, 0x6b, 0x12, 0x8d, 0x3d, 0x7e, 0x52, 0xed, 0xf1, 0xe8, 0x6c, 0x23, 0xb1, 0xcf,
	0xa7, 0xd9, 0xf8, 0x69, 0xa7, 0xe4, 0xac, 0xfe, 0x86, 0x3e, 0x75, 0x98, 0x7f, 0x48, 0x70, 0xf1,
	0xa4, 0x2c, 0xea, 0x30, 0xbe, 0xb5, 0xdc, 0xe9, 0x9b, 0xad, 0x35, 0x75, 0xed, 0xa5, 0xa3, 0x9f,
	0x20, 0x82, 0x66, 0xcd, 0x66, 0x5b, 0xdb, 0x99, 0x91, 0x77, 0xdb, 0x05, 0x71, 0x2d, 0x5d, 0xc2,
	0x97, 0x70, 0x51, 0xbc, 0x42, 0x4d, 0x5a, 0xe6, 0x74, 0xf1, 0x08, 0x35, 0x69, 0x05, 0xc7, 0x30,
	0xe0, 0x15, 0xc4, 0xde, 0x54, 0xbe, 0xc8, 0x32, 0x57, 0x10, 0x3d, 0x4e, 0x94, 0x59, 0x82, 0xe8,
	0xe3, 0x25, 0x20, 0x8f, 0xb7, 0xa6, 0x53, 0xb2, 0x5e, 0x6f, 0xa7, 0x4b, 0xc7, 0xa5, 0xcb, 0xb9,
	0x7e, 0x8a, 0xff, 0x87, 0x97, 0x65, 0x66, 0xc7, 0x35, 0xe0, 0xae, 0xb2, 0x46, 0xc7, 0x35, 0xbc,
	0xeb, 0x8b, 0xbf, 0x0e, 0x9f, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x3a, 0x26, 0x85, 0x4a,
	0x08, 0x00, 0x00,
}
