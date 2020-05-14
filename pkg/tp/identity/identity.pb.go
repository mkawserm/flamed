// Code generated by protoc-gen-go. DO NOT EDIT.
// source: identity.proto

package identity

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
	return fileDescriptor_61c7956abb761639, []int{1, 0}
}

type FlameProposal_FlameProposalType int32

const (
	FlameProposal_BATCH_ACTION          FlameProposal_FlameProposalType = 0
	FlameProposal_CREATE_USER           FlameProposal_FlameProposalType = 1
	FlameProposal_UPDATE_USER           FlameProposal_FlameProposalType = 2
	FlameProposal_DELETE_USER           FlameProposal_FlameProposalType = 3
	FlameProposal_CREATE_INDEX_META     FlameProposal_FlameProposalType = 4
	FlameProposal_UPDATE_INDEX_META     FlameProposal_FlameProposalType = 5
	FlameProposal_DELETE_INDEX_META     FlameProposal_FlameProposalType = 6
	FlameProposal_CREATE_ACCESS_CONTROL FlameProposal_FlameProposalType = 7
	FlameProposal_UPDATE_ACCESS_CONTROL FlameProposal_FlameProposalType = 8
	FlameProposal_DELETE_ACCESS_CONTROL FlameProposal_FlameProposalType = 9
	FlameProposal_BATCH_USER            FlameProposal_FlameProposalType = 10
	FlameProposal_BATCH_INDEX_META      FlameProposal_FlameProposalType = 11
	FlameProposal_BATCH_ACCESS_CONTROL  FlameProposal_FlameProposalType = 12
)

var FlameProposal_FlameProposalType_name = map[int32]string{
	0:  "BATCH_ACTION",
	1:  "CREATE_USER",
	2:  "UPDATE_USER",
	3:  "DELETE_USER",
	4:  "CREATE_INDEX_META",
	5:  "UPDATE_INDEX_META",
	6:  "DELETE_INDEX_META",
	7:  "CREATE_ACCESS_CONTROL",
	8:  "UPDATE_ACCESS_CONTROL",
	9:  "DELETE_ACCESS_CONTROL",
	10: "BATCH_USER",
	11: "BATCH_INDEX_META",
	12: "BATCH_ACCESS_CONTROL",
}

var FlameProposal_FlameProposalType_value = map[string]int32{
	"BATCH_ACTION":          0,
	"CREATE_USER":           1,
	"UPDATE_USER":           2,
	"DELETE_USER":           3,
	"CREATE_INDEX_META":     4,
	"UPDATE_INDEX_META":     5,
	"DELETE_INDEX_META":     6,
	"CREATE_ACCESS_CONTROL": 7,
	"UPDATE_ACCESS_CONTROL": 8,
	"DELETE_ACCESS_CONTROL": 9,
	"BATCH_USER":            10,
	"BATCH_INDEX_META":      11,
	"BATCH_ACCESS_CONTROL":  12,
}

func (x FlameProposal_FlameProposalType) String() string {
	return proto.EnumName(FlameProposal_FlameProposalType_name, int32(x))
}

func (FlameProposal_FlameProposalType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61c7956abb761639, []int{5, 0}
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
	return fileDescriptor_61c7956abb761639, []int{6, 0}
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
	return fileDescriptor_61c7956abb761639, []int{0}
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
	FlameActionType      FlameAction_FlameActionType `protobuf:"varint,1,opt,name=flameActionType,proto3,enum=identity.FlameAction_FlameActionType" json:"flameActionType,omitempty"`
	FlameEntry           *FlameEntry                 `protobuf:"bytes,2,opt,name=flameEntry,proto3" json:"flameEntry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *FlameAction) Reset()         { *m = FlameAction{} }
func (m *FlameAction) String() string { return proto.CompactTextString(m) }
func (*FlameAction) ProtoMessage()    {}
func (*FlameAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_61c7956abb761639, []int{1}
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
	return fileDescriptor_61c7956abb761639, []int{2}
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
	return fileDescriptor_61c7956abb761639, []int{3}
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

type FlameNamespaceList struct {
	FlameNamespaceList   [][]byte `protobuf:"bytes,1,rep,name=flameNamespaceList,proto3" json:"flameNamespaceList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlameNamespaceList) Reset()         { *m = FlameNamespaceList{} }
func (m *FlameNamespaceList) String() string { return proto.CompactTextString(m) }
func (*FlameNamespaceList) ProtoMessage()    {}
func (*FlameNamespaceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_61c7956abb761639, []int{4}
}

func (m *FlameNamespaceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameNamespaceList.Unmarshal(m, b)
}
func (m *FlameNamespaceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameNamespaceList.Marshal(b, m, deterministic)
}
func (m *FlameNamespaceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameNamespaceList.Merge(m, src)
}
func (m *FlameNamespaceList) XXX_Size() int {
	return xxx_messageInfo_FlameNamespaceList.Size(m)
}
func (m *FlameNamespaceList) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameNamespaceList.DiscardUnknown(m)
}

var xxx_messageInfo_FlameNamespaceList proto.InternalMessageInfo

func (m *FlameNamespaceList) GetFlameNamespaceList() [][]byte {
	if m != nil {
		return m.FlameNamespaceList
	}
	return nil
}

type FlameProposal struct {
	FlameProposalType    FlameProposal_FlameProposalType `protobuf:"varint,1,opt,name=flameProposalType,proto3,enum=identity.FlameProposal_FlameProposalType" json:"flameProposalType,omitempty"`
	FlameProposalData    []byte                          `protobuf:"bytes,2,opt,name=flameProposalData,proto3" json:"flameProposalData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *FlameProposal) Reset()         { *m = FlameProposal{} }
func (m *FlameProposal) String() string { return proto.CompactTextString(m) }
func (*FlameProposal) ProtoMessage()    {}
func (*FlameProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_61c7956abb761639, []int{5}
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

type FlameUser struct {
	FlameUserType        FlameUser_FlameUserType `protobuf:"varint,1,opt,name=flameUserType,proto3,enum=identity.FlameUser_FlameUserType" json:"flameUserType,omitempty"`
	Roles                string                  `protobuf:"bytes,2,opt,name=roles,proto3" json:"roles,omitempty"`
	Username             string                  `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password             string                  `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	CreatedAt            uint64                  `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            uint64                  `protobuf:"varint,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Data                 []byte                  `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	Meta                 []byte                  `protobuf:"bytes,8,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FlameUser) Reset()         { *m = FlameUser{} }
func (m *FlameUser) String() string { return proto.CompactTextString(m) }
func (*FlameUser) ProtoMessage()    {}
func (*FlameUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_61c7956abb761639, []int{6}
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

func (m *FlameUser) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *FlameUser) GetMeta() []byte {
	if m != nil {
		return m.Meta
	}
	return nil
}

type FlameUserList struct {
	FlameUserList        []*FlameUser `protobuf:"bytes,1,rep,name=flameUserList,proto3" json:"flameUserList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FlameUserList) Reset()         { *m = FlameUserList{} }
func (m *FlameUserList) String() string { return proto.CompactTextString(m) }
func (*FlameUserList) ProtoMessage()    {}
func (*FlameUserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_61c7956abb761639, []int{7}
}

func (m *FlameUserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameUserList.Unmarshal(m, b)
}
func (m *FlameUserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameUserList.Marshal(b, m, deterministic)
}
func (m *FlameUserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameUserList.Merge(m, src)
}
func (m *FlameUserList) XXX_Size() int {
	return xxx_messageInfo_FlameUserList.Size(m)
}
func (m *FlameUserList) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameUserList.DiscardUnknown(m)
}

var xxx_messageInfo_FlameUserList proto.InternalMessageInfo

func (m *FlameUserList) GetFlameUserList() []*FlameUser {
	if m != nil {
		return m.FlameUserList
	}
	return nil
}

type FlameUsernameList struct {
	FlameUsernameList    []string `protobuf:"bytes,1,rep,name=flameUsernameList,proto3" json:"flameUsernameList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlameUsernameList) Reset()         { *m = FlameUsernameList{} }
func (m *FlameUsernameList) String() string { return proto.CompactTextString(m) }
func (*FlameUsernameList) ProtoMessage()    {}
func (*FlameUsernameList) Descriptor() ([]byte, []int) {
	return fileDescriptor_61c7956abb761639, []int{8}
}

func (m *FlameUsernameList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameUsernameList.Unmarshal(m, b)
}
func (m *FlameUsernameList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameUsernameList.Marshal(b, m, deterministic)
}
func (m *FlameUsernameList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameUsernameList.Merge(m, src)
}
func (m *FlameUsernameList) XXX_Size() int {
	return xxx_messageInfo_FlameUsernameList.Size(m)
}
func (m *FlameUsernameList) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameUsernameList.DiscardUnknown(m)
}

var xxx_messageInfo_FlameUsernameList proto.InternalMessageInfo

func (m *FlameUsernameList) GetFlameUsernameList() []string {
	if m != nil {
		return m.FlameUsernameList
	}
	return nil
}

type FlameAccessControl struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Namespace            []byte   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Permission           []byte   `protobuf:"bytes,3,opt,name=permission,proto3" json:"permission,omitempty"`
	CreatedAt            uint64   `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            uint64   `protobuf:"varint,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Data                 []byte   `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	Meta                 []byte   `protobuf:"bytes,7,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlameAccessControl) Reset()         { *m = FlameAccessControl{} }
func (m *FlameAccessControl) String() string { return proto.CompactTextString(m) }
func (*FlameAccessControl) ProtoMessage()    {}
func (*FlameAccessControl) Descriptor() ([]byte, []int) {
	return fileDescriptor_61c7956abb761639, []int{9}
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

func (m *FlameAccessControl) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *FlameAccessControl) GetMeta() []byte {
	if m != nil {
		return m.Meta
	}
	return nil
}

type FlameAccessControlList struct {
	FlameAccessControlList []*FlameAccessControl `protobuf:"bytes,1,rep,name=flameAccessControlList,proto3" json:"flameAccessControlList,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}              `json:"-"`
	XXX_unrecognized       []byte                `json:"-"`
	XXX_sizecache          int32                 `json:"-"`
}

func (m *FlameAccessControlList) Reset()         { *m = FlameAccessControlList{} }
func (m *FlameAccessControlList) String() string { return proto.CompactTextString(m) }
func (*FlameAccessControlList) ProtoMessage()    {}
func (*FlameAccessControlList) Descriptor() ([]byte, []int) {
	return fileDescriptor_61c7956abb761639, []int{10}
}

func (m *FlameAccessControlList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlameAccessControlList.Unmarshal(m, b)
}
func (m *FlameAccessControlList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlameAccessControlList.Marshal(b, m, deterministic)
}
func (m *FlameAccessControlList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlameAccessControlList.Merge(m, src)
}
func (m *FlameAccessControlList) XXX_Size() int {
	return xxx_messageInfo_FlameAccessControlList.Size(m)
}
func (m *FlameAccessControlList) XXX_DiscardUnknown() {
	xxx_messageInfo_FlameAccessControlList.DiscardUnknown(m)
}

var xxx_messageInfo_FlameAccessControlList proto.InternalMessageInfo

func (m *FlameAccessControlList) GetFlameAccessControlList() []*FlameAccessControl {
	if m != nil {
		return m.FlameAccessControlList
	}
	return nil
}

func init() {
	proto.RegisterEnum("identity.FlameAction_FlameActionType", FlameAction_FlameActionType_name, FlameAction_FlameActionType_value)
	proto.RegisterEnum("identity.FlameProposal_FlameProposalType", FlameProposal_FlameProposalType_name, FlameProposal_FlameProposalType_value)
	proto.RegisterEnum("identity.FlameUser_FlameUserType", FlameUser_FlameUserType_name, FlameUser_FlameUserType_value)
	proto.RegisterType((*FlameEntry)(nil), "identity.FlameEntry")
	proto.RegisterType((*FlameAction)(nil), "identity.FlameAction")
	proto.RegisterType((*FlameBatchAction)(nil), "identity.FlameBatchAction")
	proto.RegisterType((*FlameBatchRead)(nil), "identity.FlameBatchRead")
	proto.RegisterType((*FlameNamespaceList)(nil), "identity.FlameNamespaceList")
	proto.RegisterType((*FlameProposal)(nil), "identity.FlameProposal")
	proto.RegisterType((*FlameUser)(nil), "identity.FlameUser")
	proto.RegisterType((*FlameUserList)(nil), "identity.FlameUserList")
	proto.RegisterType((*FlameUsernameList)(nil), "identity.FlameUsernameList")
	proto.RegisterType((*FlameAccessControl)(nil), "identity.FlameAccessControl")
	proto.RegisterType((*FlameAccessControlList)(nil), "identity.FlameAccessControlList")
}

func init() { proto.RegisterFile("identity.proto", fileDescriptor_61c7956abb761639) }

var fileDescriptor_61c7956abb761639 = []byte{
	// 719 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xf3, 0xd5, 0x64, 0x92, 0xa6, 0xee, 0x92, 0x56, 0x06, 0x55, 0xa8, 0xac, 0x84, 0x54,
	0x24, 0x54, 0xa1, 0xc2, 0x05, 0x09, 0x09, 0xb9, 0x8e, 0x81, 0xa2, 0xd4, 0x89, 0x36, 0x8e, 0xe0,
	0x16, 0x99, 0x64, 0x23, 0x22, 0x12, 0xdb, 0xb2, 0x1d, 0x50, 0xfe, 0x14, 0x67, 0xfe, 0x01, 0x77,
	0x2e, 0xfc, 0x1d, 0xb4, 0x1f, 0x4e, 0xd6, 0x1f, 0xe5, 0x36, 0xf3, 0xde, 0xcc, 0xcb, 0xee, 0xbc,
	0xf1, 0x06, 0xba, 0xcb, 0x39, 0xf5, 0x93, 0x65, 0xb2, 0xbd, 0x0a, 0xa3, 0x20, 0x09, 0x50, 0x33,
	0xcd, 0x31, 0x01, 0x78, 0xb7, 0xf2, 0xd6, 0xd4, 0xf6, 0x93, 0x68, 0x8b, 0xce, 0xa1, 0xe5, 0x7b,
	0x6b, 0x1a, 0x87, 0xde, 0x8c, 0x1a, 0xda, 0x85, 0x76, 0xd9, 0x21, 0x7b, 0x00, 0xe9, 0x50, 0xfd,
	0x46, 0xb7, 0x46, 0x85, 0xe3, 0x2c, 0x44, 0x3d, 0xa8, 0x7f, 0xf7, 0x56, 0x1b, 0x6a, 0x54, 0x39,
	0x26, 0x12, 0xfc, 0x57, 0x83, 0x36, 0x17, 0x35, 0x67, 0xc9, 0x32, 0xf0, 0xd1, 0x10, 0x8e, 0x17,
	0xfb, 0xd4, 0xdd, 0x86, 0x42, 0xbb, 0x7b, 0xfd, 0xf4, 0x6a, 0x77, 0x2e, 0xa5, 0x5e, 0x8d, 0x59,
	0x31, 0xc9, 0x77, 0xa3, 0x57, 0x00, 0x8b, 0xdd, 0xa1, 0xf9, 0x79, 0xda, 0xd7, 0xbd, 0x9c, 0x16,
	0xe7, 0x88, 0x52, 0x87, 0x4d, 0x38, 0xce, 0x29, 0x23, 0x80, 0x86, 0x45, 0x6c, 0xd3, 0xb5, 0xf5,
	0x03, 0x16, 0x4f, 0x46, 0x7d, 0x16, 0x6b, 0x2c, 0xee, 0xdb, 0x03, 0xdb, 0xb5, 0xf5, 0x0a, 0x8b,
	0xcd, 0xd1, 0xc8, 0x76, 0xfa, 0x7a, 0x15, 0x8f, 0x41, 0xe7, 0x12, 0x37, 0x5e, 0x32, 0xfb, 0x2a,
	0x6f, 0xf7, 0x36, 0x73, 0xbb, 0xc1, 0x32, 0x4e, 0x0c, 0xed, 0xa2, 0x7a, 0xd9, 0xbe, 0x3e, 0x2d,
	0xbd, 0x1d, 0xc9, 0x57, 0x63, 0x07, 0xba, 0x7b, 0x51, 0x42, 0xbd, 0x39, 0x7a, 0x03, 0xdd, 0xfd,
	0xb9, 0x15, 0xc5, 0xf2, 0x3b, 0xe6, 0x6a, 0x71, 0x1f, 0x10, 0x67, 0x9d, 0xd4, 0x38, 0x86, 0xa2,
	0x2b, 0x40, 0x8b, 0x02, 0xca, 0x75, 0x3b, 0xa4, 0x84, 0xc1, 0xbf, 0xab, 0x70, 0xc4, 0x65, 0x46,
	0x51, 0x10, 0x06, 0xb1, 0xb7, 0x42, 0x9f, 0xe0, 0x64, 0xa1, 0x02, 0x8a, 0x91, 0xcf, 0x72, 0x07,
	0x4b, 0x4b, 0xb2, 0x19, 0x37, 0xb3, 0xa8, 0x81, 0x9e, 0xe7, 0x84, 0xfb, 0x5e, 0xe2, 0xc9, 0x2d,
	0x2b, 0x12, 0xf8, 0x67, 0x05, 0x4e, 0x0a, 0xb2, 0x48, 0x87, 0xce, 0x8d, 0xe9, 0x5a, 0x1f, 0xa6,
	0xa6, 0xe5, 0xde, 0x0e, 0x1d, 0xfd, 0x00, 0x1d, 0x43, 0x5b, 0x78, 0x3b, 0x9d, 0x8c, 0x6d, 0xa2,
	0x6b, 0x0c, 0x10, 0x06, 0x0b, 0xa0, 0xc2, 0x00, 0xe1, 0xb2, 0x00, 0xaa, 0xe8, 0x14, 0x4e, 0x64,
	0xcb, 0xad, 0xd3, 0xb7, 0x3f, 0x4f, 0xef, 0x6c, 0xd7, 0xd4, 0x6b, 0x0c, 0x96, 0x8d, 0x0a, 0x5c,
	0x67, 0xb0, 0x6c, 0x57, 0xe0, 0x06, 0x7a, 0x08, 0xa7, 0x52, 0xc4, 0xb4, 0x2c, 0x7b, 0x3c, 0x9e,
	0x5a, 0x43, 0xc7, 0x25, 0xc3, 0x81, 0x7e, 0xc8, 0x28, 0x29, 0x94, 0xa3, 0x9a, 0x8c, 0x92, 0x62,
	0x39, 0xaa, 0x85, 0xba, 0x00, 0xe2, 0x6a, 0xfc, 0x94, 0x80, 0x7a, 0xa0, 0x8b, 0x5c, 0xf9, 0xd9,
	0x36, 0x32, 0xa0, 0x97, 0x0e, 0x20, 0xd3, 0xdf, 0xc1, 0xbf, 0x2a, 0xd0, 0xe2, 0x03, 0x9b, 0xc4,
	0x34, 0x42, 0xef, 0xe1, 0x68, 0x91, 0x26, 0x8a, 0x83, 0x4f, 0x72, 0x0e, 0x32, 0x7a, 0x1f, 0x71,
	0xe7, 0xb2, 0x7d, 0xec, 0xdb, 0x8f, 0x82, 0x15, 0x8d, 0xb9, 0x53, 0x2d, 0x22, 0x12, 0xf4, 0x08,
	0x9a, 0x9b, 0x98, 0x46, 0xec, 0xd1, 0xe0, 0x8f, 0x42, 0x8b, 0xec, 0x72, 0xc6, 0x85, 0x5e, 0x1c,
	0xff, 0x08, 0xa2, 0xb9, 0x51, 0x13, 0x5c, 0x9a, 0xb3, 0x97, 0x67, 0x16, 0x51, 0x2f, 0xa1, 0x73,
	0x33, 0x31, 0xea, 0x17, 0xda, 0x65, 0x8d, 0xec, 0x01, 0xc6, 0x6e, 0xc2, 0xb9, 0x64, 0x1b, 0x82,
	0xdd, 0x01, 0x08, 0x41, 0x6d, 0xce, 0x56, 0xe6, 0x90, 0xaf, 0x0c, 0x8f, 0x19, 0xb6, 0xa6, 0x89,
	0x67, 0x34, 0x05, 0xc6, 0x62, 0xfc, 0x42, 0x6e, 0xf4, 0xee, 0x0a, 0x5d, 0x80, 0xf1, 0x64, 0x64,
	0x13, 0x31, 0x59, 0xbe, 0x32, 0xce, 0x90, 0xdc, 0x99, 0x03, 0xb9, 0x32, 0xf8, 0xa3, 0xd2, 0xc1,
	0xbf, 0xa2, 0xd7, 0xca, 0xf4, 0x94, 0x0f, 0xf3, 0x41, 0xc9, 0xf4, 0x48, 0xb6, 0x12, 0x9b, 0x72,
	0x6d, 0x27, 0x72, 0x1c, 0x5c, 0x2f, 0x5d, 0x7d, 0x15, 0xe4, 0x9a, 0x2d, 0x52, 0x24, 0xf0, 0x1f,
	0x4d, 0x7e, 0xda, 0xe6, 0x6c, 0x46, 0xe3, 0xd8, 0x0a, 0xfc, 0x24, 0x0a, 0x56, 0x99, 0x99, 0x6b,
	0xb9, 0x99, 0x67, 0x5e, 0xf4, 0x4a, 0xfe, 0x45, 0x7f, 0x0c, 0x10, 0xd2, 0x68, 0xbd, 0x8c, 0xe3,
	0x65, 0xe0, 0xcb, 0x47, 0x5c, 0x41, 0xb2, 0xae, 0xd4, 0xfe, 0xeb, 0x4a, 0xfd, 0x3e, 0x57, 0x1a,
	0x25, 0xae, 0x1c, 0x2a, 0xae, 0xf8, 0x70, 0x56, 0xbc, 0x13, 0x1f, 0x8e, 0x0b, 0x67, 0x8b, 0x52,
	0x46, 0x4e, 0xfd, 0xbc, 0xf0, 0xc0, 0x2a, 0x75, 0xe4, 0x9e, 0xde, 0x2f, 0x0d, 0xfe, 0x17, 0xf8,
	0xf2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xa5, 0x35, 0x6c, 0x14, 0x07, 0x00, 0x00,
}