// Code generated by protoc-gen-go. DO NOT EDIT.
// source: admin.proto

package admin

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	pb "github.com/mkawserm/flamed/pkg/pb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type UserRequest struct {
	ClusterID            uint64   `protobuf:"varint,1,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetClusterID() uint64 {
	if m != nil {
		return m.ClusterID
	}
	return 0
}

func (m *UserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UpsertUserRequest struct {
	ClusterID            uint64   `protobuf:"varint,1,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
	User                 *pb.User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpsertUserRequest) Reset()         { *m = UpsertUserRequest{} }
func (m *UpsertUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpsertUserRequest) ProtoMessage()    {}
func (*UpsertUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{1}
}

func (m *UpsertUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertUserRequest.Unmarshal(m, b)
}
func (m *UpsertUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertUserRequest.Marshal(b, m, deterministic)
}
func (m *UpsertUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertUserRequest.Merge(m, src)
}
func (m *UpsertUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpsertUserRequest.Size(m)
}
func (m *UpsertUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertUserRequest proto.InternalMessageInfo

func (m *UpsertUserRequest) GetClusterID() uint64 {
	if m != nil {
		return m.ClusterID
	}
	return 0
}

func (m *UpsertUserRequest) GetUser() *pb.User {
	if m != nil {
		return m.User
	}
	return nil
}

type ChangeUserPasswordRequest struct {
	ClusterID            uint64   `protobuf:"varint,1,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeUserPasswordRequest) Reset()         { *m = ChangeUserPasswordRequest{} }
func (m *ChangeUserPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*ChangeUserPasswordRequest) ProtoMessage()    {}
func (*ChangeUserPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{2}
}

func (m *ChangeUserPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeUserPasswordRequest.Unmarshal(m, b)
}
func (m *ChangeUserPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeUserPasswordRequest.Marshal(b, m, deterministic)
}
func (m *ChangeUserPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeUserPasswordRequest.Merge(m, src)
}
func (m *ChangeUserPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_ChangeUserPasswordRequest.Size(m)
}
func (m *ChangeUserPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeUserPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeUserPasswordRequest proto.InternalMessageInfo

func (m *ChangeUserPasswordRequest) GetClusterID() uint64 {
	if m != nil {
		return m.ClusterID
	}
	return 0
}

func (m *ChangeUserPasswordRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ChangeUserPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AccessControlRequest struct {
	ClusterID            uint64   `protobuf:"varint,1,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
	Namespace            []byte   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessControlRequest) Reset()         { *m = AccessControlRequest{} }
func (m *AccessControlRequest) String() string { return proto.CompactTextString(m) }
func (*AccessControlRequest) ProtoMessage()    {}
func (*AccessControlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{3}
}

func (m *AccessControlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessControlRequest.Unmarshal(m, b)
}
func (m *AccessControlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessControlRequest.Marshal(b, m, deterministic)
}
func (m *AccessControlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessControlRequest.Merge(m, src)
}
func (m *AccessControlRequest) XXX_Size() int {
	return xxx_messageInfo_AccessControlRequest.Size(m)
}
func (m *AccessControlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessControlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccessControlRequest proto.InternalMessageInfo

func (m *AccessControlRequest) GetClusterID() uint64 {
	if m != nil {
		return m.ClusterID
	}
	return 0
}

func (m *AccessControlRequest) GetNamespace() []byte {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *AccessControlRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UpsertAccessControlRequest struct {
	ClusterID            uint64            `protobuf:"varint,1,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
	AccessControl        *pb.AccessControl `protobuf:"bytes,2,opt,name=accessControl,proto3" json:"accessControl,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UpsertAccessControlRequest) Reset()         { *m = UpsertAccessControlRequest{} }
func (m *UpsertAccessControlRequest) String() string { return proto.CompactTextString(m) }
func (*UpsertAccessControlRequest) ProtoMessage()    {}
func (*UpsertAccessControlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{4}
}

func (m *UpsertAccessControlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertAccessControlRequest.Unmarshal(m, b)
}
func (m *UpsertAccessControlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertAccessControlRequest.Marshal(b, m, deterministic)
}
func (m *UpsertAccessControlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertAccessControlRequest.Merge(m, src)
}
func (m *UpsertAccessControlRequest) XXX_Size() int {
	return xxx_messageInfo_UpsertAccessControlRequest.Size(m)
}
func (m *UpsertAccessControlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertAccessControlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertAccessControlRequest proto.InternalMessageInfo

func (m *UpsertAccessControlRequest) GetClusterID() uint64 {
	if m != nil {
		return m.ClusterID
	}
	return 0
}

func (m *UpsertAccessControlRequest) GetAccessControl() *pb.AccessControl {
	if m != nil {
		return m.AccessControl
	}
	return nil
}

type IndexMetaRequest struct {
	ClusterID            uint64   `protobuf:"varint,1,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
	Namespace            []byte   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexMetaRequest) Reset()         { *m = IndexMetaRequest{} }
func (m *IndexMetaRequest) String() string { return proto.CompactTextString(m) }
func (*IndexMetaRequest) ProtoMessage()    {}
func (*IndexMetaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{5}
}

func (m *IndexMetaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexMetaRequest.Unmarshal(m, b)
}
func (m *IndexMetaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexMetaRequest.Marshal(b, m, deterministic)
}
func (m *IndexMetaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexMetaRequest.Merge(m, src)
}
func (m *IndexMetaRequest) XXX_Size() int {
	return xxx_messageInfo_IndexMetaRequest.Size(m)
}
func (m *IndexMetaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexMetaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IndexMetaRequest proto.InternalMessageInfo

func (m *IndexMetaRequest) GetClusterID() uint64 {
	if m != nil {
		return m.ClusterID
	}
	return 0
}

func (m *IndexMetaRequest) GetNamespace() []byte {
	if m != nil {
		return m.Namespace
	}
	return nil
}

type UpsertIndexMetaRequest struct {
	ClusterID            uint64        `protobuf:"varint,1,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
	IndexMeta            *pb.IndexMeta `protobuf:"bytes,2,opt,name=indexMeta,proto3" json:"indexMeta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpsertIndexMetaRequest) Reset()         { *m = UpsertIndexMetaRequest{} }
func (m *UpsertIndexMetaRequest) String() string { return proto.CompactTextString(m) }
func (*UpsertIndexMetaRequest) ProtoMessage()    {}
func (*UpsertIndexMetaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{6}
}

func (m *UpsertIndexMetaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertIndexMetaRequest.Unmarshal(m, b)
}
func (m *UpsertIndexMetaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertIndexMetaRequest.Marshal(b, m, deterministic)
}
func (m *UpsertIndexMetaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertIndexMetaRequest.Merge(m, src)
}
func (m *UpsertIndexMetaRequest) XXX_Size() int {
	return xxx_messageInfo_UpsertIndexMetaRequest.Size(m)
}
func (m *UpsertIndexMetaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertIndexMetaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertIndexMetaRequest proto.InternalMessageInfo

func (m *UpsertIndexMetaRequest) GetClusterID() uint64 {
	if m != nil {
		return m.ClusterID
	}
	return 0
}

func (m *UpsertIndexMetaRequest) GetIndexMeta() *pb.IndexMeta {
	if m != nil {
		return m.IndexMeta
	}
	return nil
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "admin.UserRequest")
	proto.RegisterType((*UpsertUserRequest)(nil), "admin.UpsertUserRequest")
	proto.RegisterType((*ChangeUserPasswordRequest)(nil), "admin.ChangeUserPasswordRequest")
	proto.RegisterType((*AccessControlRequest)(nil), "admin.AccessControlRequest")
	proto.RegisterType((*UpsertAccessControlRequest)(nil), "admin.UpsertAccessControlRequest")
	proto.RegisterType((*IndexMetaRequest)(nil), "admin.IndexMetaRequest")
	proto.RegisterType((*UpsertIndexMetaRequest)(nil), "admin.UpsertIndexMetaRequest")
}

func init() { proto.RegisterFile("admin.proto", fileDescriptor_73a7fc70dcc2027c) }

var fileDescriptor_73a7fc70dcc2027c = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x5b, 0xb6, 0x41, 0x7b, 0xdd, 0xb4, 0xed, 0x36, 0x41, 0x09, 0x43, 0x2a, 0x7e, 0x1a,
	0x42, 0x4a, 0xa5, 0xf1, 0x00, 0x4f, 0x48, 0xa3, 0x95, 0xca, 0x90, 0x18, 0x55, 0xa4, 0x7d, 0x00,
	0x37, 0x3d, 0xba, 0x6a, 0x49, 0xec, 0xd9, 0xae, 0xc6, 0x77, 0xe2, 0x4b, 0x22, 0xc7, 0xe9, 0xb2,
	0xa4, 0xb5, 0xd4, 0xc1, 0x1e, 0xef, 0xec, 0xff, 0xdf, 0xff, 0x3b, 0xfd, 0x12, 0xe8, 0xf0, 0x69,
	0x3a, 0xcf, 0x42, 0xa9, 0x84, 0x11, 0xb8, 0x93, 0x17, 0x41, 0x38, 0x9b, 0x9b, 0xeb, 0xc5, 0x24,
	0x8c, 0x45, 0xda, 0x4f, 0x6f, 0xf8, 0x9d, 0x26, 0x95, 0xf6, 0x7f, 0x25, 0x3c, 0xa5, 0x69, 0x5f,
	0xde, 0xcc, 0xfa, 0x72, 0x52, 0x54, 0x4e, 0xc6, 0x46, 0xd0, 0xb9, 0xd2, 0xa4, 0x22, 0xba, 0x5d,
	0x90, 0x36, 0x78, 0x02, 0xed, 0x38, 0x59, 0x68, 0x43, 0xea, 0x62, 0xd8, 0x6d, 0xf6, 0x9a, 0xa7,
	0xdb, 0x51, 0xd9, 0xc0, 0x00, 0x5a, 0x0b, 0x4d, 0x2a, 0xe3, 0x29, 0x75, 0x9f, 0xf5, 0x9a, 0xa7,
	0xed, 0xe8, 0xbe, 0x66, 0x3f, 0xe1, 0xf0, 0x4a, 0x6a, 0x52, 0x66, 0x73, 0xbb, 0x13, 0xd8, 0xb6,
	0xf2, 0xdc, 0xaa, 0x73, 0xd6, 0x0a, 0xe5, 0x24, 0xcc, 0xc5, 0x79, 0x97, 0xdd, 0xc2, 0xeb, 0xc1,
	0x35, 0xcf, 0x66, 0x64, 0x7b, 0x63, 0xae, 0xf5, 0x9d, 0x50, 0xd3, 0xff, 0xce, 0x69, 0xcf, 0x64,
	0x61, 0xd6, 0xdd, 0x72, 0x67, 0xcb, 0x9a, 0x65, 0x70, 0x7c, 0x1e, 0xc7, 0xa4, 0xf5, 0x40, 0x64,
	0x46, 0x89, 0x64, 0xd3, 0x31, 0xda, 0xd6, 0x59, 0x4b, 0x1e, 0xbb, 0xe7, 0x76, 0xa3, 0xb2, 0x51,
	0xc9, 0xb2, 0x55, 0xdb, 0x99, 0x86, 0xc0, 0xed, 0xec, 0x1f, 0x5e, 0xfd, 0x04, 0x7b, 0xfc, 0xa1,
	0xaa, 0xd8, 0xe2, 0xa1, 0xdd, 0x62, 0xd5, 0xae, 0x7a, 0x8f, 0x5d, 0xc2, 0xc1, 0x45, 0x36, 0xa5,
	0xdf, 0x3f, 0xc8, 0xf0, 0x27, 0x18, 0x90, 0xc5, 0xf0, 0xd2, 0x0d, 0xf1, 0x48, 0xd7, 0x0f, 0xd0,
	0x9e, 0x2f, 0x15, 0x45, 0xf8, 0x3d, 0x1b, 0xbe, 0xb4, 0x29, 0xcf, 0xcf, 0xfe, 0xec, 0x40, 0xeb,
	0xdc, 0x02, 0x1e, 0x8d, 0x07, 0xf8, 0x1e, 0x5e, 0x8c, 0x28, 0xe7, 0x0c, 0x31, 0x74, 0xdf, 0xc0,
	0x03, 0xe8, 0x82, 0x7b, 0x90, 0x58, 0x03, 0xbf, 0x00, 0x94, 0x54, 0x62, 0x77, 0x79, 0xbb, 0x0e,
	0x6a, 0x70, 0x6c, 0x35, 0x63, 0x25, 0xa4, 0xd0, 0x3c, 0x89, 0x48, 0x4b, 0x91, 0x69, 0x62, 0x0d,
	0xbc, 0x04, 0x5c, 0x85, 0x10, 0x7b, 0x85, 0x8f, 0x97, 0x4f, 0xaf, 0xdf, 0x67, 0x80, 0x21, 0x25,
	0x64, 0xc8, 0x9b, 0xde, 0xa7, 0x1c, 0xc2, 0xc1, 0x88, 0xaa, 0xa0, 0xe0, 0x9b, 0x42, 0xbf, 0x0e,
	0x9f, 0x60, 0x95, 0x04, 0xd6, 0xc0, 0x31, 0x1c, 0xad, 0x21, 0x0e, 0xdf, 0x55, 0x16, 0xb3, 0xd6,
	0xce, 0x97, 0xeb, 0x3b, 0x1c, 0xb9, 0x89, 0x1e, 0x11, 0xcd, 0xbf, 0x9d, 0xdd, 0x11, 0x95, 0x1c,
	0xe1, 0xab, 0xc2, 0xa4, 0x4e, 0x56, 0x50, 0x05, 0x85, 0x35, 0xf0, 0x1b, 0xec, 0xd7, 0x20, 0xc4,
	0xb7, 0x95, 0x99, 0x56, 0x2c, 0x7c, 0x19, 0xbe, 0xc2, 0xbe, 0x9b, 0x67, 0x83, 0x18, 0x1e, 0x8f,
	0xc9, 0xf3, 0xfc, 0xdf, 0xfa, 0xf1, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0x53, 0x9a, 0x4d,
	0xa1, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdminRPCClient is the client API for AdminRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminRPCClient interface {
	GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*pb.User, error)
	UpsertUser(ctx context.Context, in *UpsertUserRequest, opts ...grpc.CallOption) (*pb.ProposalResponse, error)
	ChangeUserPassword(ctx context.Context, in *ChangeUserPasswordRequest, opts ...grpc.CallOption) (*pb.ProposalResponse, error)
	DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*pb.ProposalResponse, error)
	GetAccessControl(ctx context.Context, in *AccessControlRequest, opts ...grpc.CallOption) (*pb.AccessControl, error)
	UpsertAccessControl(ctx context.Context, in *UpsertAccessControlRequest, opts ...grpc.CallOption) (*pb.ProposalResponse, error)
	DeleteAccessControl(ctx context.Context, in *AccessControlRequest, opts ...grpc.CallOption) (*pb.ProposalResponse, error)
	GetIndexMeta(ctx context.Context, in *IndexMetaRequest, opts ...grpc.CallOption) (*pb.IndexMeta, error)
	UpsertIndexMeta(ctx context.Context, in *UpsertIndexMetaRequest, opts ...grpc.CallOption) (*pb.ProposalResponse, error)
	DeleteIndexMeta(ctx context.Context, in *IndexMetaRequest, opts ...grpc.CallOption) (*pb.ProposalResponse, error)
}

type adminRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminRPCClient(cc grpc.ClientConnInterface) AdminRPCClient {
	return &adminRPCClient{cc}
}

func (c *adminRPCClient) GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*pb.User, error) {
	out := new(pb.User)
	err := c.cc.Invoke(ctx, "/admin.AdminRPC/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminRPCClient) UpsertUser(ctx context.Context, in *UpsertUserRequest, opts ...grpc.CallOption) (*pb.ProposalResponse, error) {
	out := new(pb.ProposalResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminRPC/UpsertUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminRPCClient) ChangeUserPassword(ctx context.Context, in *ChangeUserPasswordRequest, opts ...grpc.CallOption) (*pb.ProposalResponse, error) {
	out := new(pb.ProposalResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminRPC/ChangeUserPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminRPCClient) DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*pb.ProposalResponse, error) {
	out := new(pb.ProposalResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminRPC/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminRPCClient) GetAccessControl(ctx context.Context, in *AccessControlRequest, opts ...grpc.CallOption) (*pb.AccessControl, error) {
	out := new(pb.AccessControl)
	err := c.cc.Invoke(ctx, "/admin.AdminRPC/GetAccessControl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminRPCClient) UpsertAccessControl(ctx context.Context, in *UpsertAccessControlRequest, opts ...grpc.CallOption) (*pb.ProposalResponse, error) {
	out := new(pb.ProposalResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminRPC/UpsertAccessControl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminRPCClient) DeleteAccessControl(ctx context.Context, in *AccessControlRequest, opts ...grpc.CallOption) (*pb.ProposalResponse, error) {
	out := new(pb.ProposalResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminRPC/DeleteAccessControl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminRPCClient) GetIndexMeta(ctx context.Context, in *IndexMetaRequest, opts ...grpc.CallOption) (*pb.IndexMeta, error) {
	out := new(pb.IndexMeta)
	err := c.cc.Invoke(ctx, "/admin.AdminRPC/GetIndexMeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminRPCClient) UpsertIndexMeta(ctx context.Context, in *UpsertIndexMetaRequest, opts ...grpc.CallOption) (*pb.ProposalResponse, error) {
	out := new(pb.ProposalResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminRPC/UpsertIndexMeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminRPCClient) DeleteIndexMeta(ctx context.Context, in *IndexMetaRequest, opts ...grpc.CallOption) (*pb.ProposalResponse, error) {
	out := new(pb.ProposalResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminRPC/DeleteIndexMeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminRPCServer is the server API for AdminRPC service.
type AdminRPCServer interface {
	GetUser(context.Context, *UserRequest) (*pb.User, error)
	UpsertUser(context.Context, *UpsertUserRequest) (*pb.ProposalResponse, error)
	ChangeUserPassword(context.Context, *ChangeUserPasswordRequest) (*pb.ProposalResponse, error)
	DeleteUser(context.Context, *UserRequest) (*pb.ProposalResponse, error)
	GetAccessControl(context.Context, *AccessControlRequest) (*pb.AccessControl, error)
	UpsertAccessControl(context.Context, *UpsertAccessControlRequest) (*pb.ProposalResponse, error)
	DeleteAccessControl(context.Context, *AccessControlRequest) (*pb.ProposalResponse, error)
	GetIndexMeta(context.Context, *IndexMetaRequest) (*pb.IndexMeta, error)
	UpsertIndexMeta(context.Context, *UpsertIndexMetaRequest) (*pb.ProposalResponse, error)
	DeleteIndexMeta(context.Context, *IndexMetaRequest) (*pb.ProposalResponse, error)
}

// UnimplementedAdminRPCServer can be embedded to have forward compatible implementations.
type UnimplementedAdminRPCServer struct {
}

func (*UnimplementedAdminRPCServer) GetUser(ctx context.Context, req *UserRequest) (*pb.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedAdminRPCServer) UpsertUser(ctx context.Context, req *UpsertUserRequest) (*pb.ProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertUser not implemented")
}
func (*UnimplementedAdminRPCServer) ChangeUserPassword(ctx context.Context, req *ChangeUserPasswordRequest) (*pb.ProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserPassword not implemented")
}
func (*UnimplementedAdminRPCServer) DeleteUser(ctx context.Context, req *UserRequest) (*pb.ProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (*UnimplementedAdminRPCServer) GetAccessControl(ctx context.Context, req *AccessControlRequest) (*pb.AccessControl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessControl not implemented")
}
func (*UnimplementedAdminRPCServer) UpsertAccessControl(ctx context.Context, req *UpsertAccessControlRequest) (*pb.ProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertAccessControl not implemented")
}
func (*UnimplementedAdminRPCServer) DeleteAccessControl(ctx context.Context, req *AccessControlRequest) (*pb.ProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccessControl not implemented")
}
func (*UnimplementedAdminRPCServer) GetIndexMeta(ctx context.Context, req *IndexMetaRequest) (*pb.IndexMeta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndexMeta not implemented")
}
func (*UnimplementedAdminRPCServer) UpsertIndexMeta(ctx context.Context, req *UpsertIndexMetaRequest) (*pb.ProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertIndexMeta not implemented")
}
func (*UnimplementedAdminRPCServer) DeleteIndexMeta(ctx context.Context, req *IndexMetaRequest) (*pb.ProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIndexMeta not implemented")
}

func RegisterAdminRPCServer(s *grpc.Server, srv AdminRPCServer) {
	s.RegisterService(&_AdminRPC_serviceDesc, srv)
}

func _AdminRPC_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminRPCServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminRPC/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminRPCServer).GetUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminRPC_UpsertUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminRPCServer).UpsertUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminRPC/UpsertUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminRPCServer).UpsertUser(ctx, req.(*UpsertUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminRPC_ChangeUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminRPCServer).ChangeUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminRPC/ChangeUserPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminRPCServer).ChangeUserPassword(ctx, req.(*ChangeUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminRPC_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminRPCServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminRPC/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminRPCServer).DeleteUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminRPC_GetAccessControl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessControlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminRPCServer).GetAccessControl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminRPC/GetAccessControl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminRPCServer).GetAccessControl(ctx, req.(*AccessControlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminRPC_UpsertAccessControl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertAccessControlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminRPCServer).UpsertAccessControl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminRPC/UpsertAccessControl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminRPCServer).UpsertAccessControl(ctx, req.(*UpsertAccessControlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminRPC_DeleteAccessControl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessControlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminRPCServer).DeleteAccessControl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminRPC/DeleteAccessControl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminRPCServer).DeleteAccessControl(ctx, req.(*AccessControlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminRPC_GetIndexMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexMetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminRPCServer).GetIndexMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminRPC/GetIndexMeta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminRPCServer).GetIndexMeta(ctx, req.(*IndexMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminRPC_UpsertIndexMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertIndexMetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminRPCServer).UpsertIndexMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminRPC/UpsertIndexMeta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminRPCServer).UpsertIndexMeta(ctx, req.(*UpsertIndexMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminRPC_DeleteIndexMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexMetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminRPCServer).DeleteIndexMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminRPC/DeleteIndexMeta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminRPCServer).DeleteIndexMeta(ctx, req.(*IndexMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "admin.AdminRPC",
	HandlerType: (*AdminRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _AdminRPC_GetUser_Handler,
		},
		{
			MethodName: "UpsertUser",
			Handler:    _AdminRPC_UpsertUser_Handler,
		},
		{
			MethodName: "ChangeUserPassword",
			Handler:    _AdminRPC_ChangeUserPassword_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _AdminRPC_DeleteUser_Handler,
		},
		{
			MethodName: "GetAccessControl",
			Handler:    _AdminRPC_GetAccessControl_Handler,
		},
		{
			MethodName: "UpsertAccessControl",
			Handler:    _AdminRPC_UpsertAccessControl_Handler,
		},
		{
			MethodName: "DeleteAccessControl",
			Handler:    _AdminRPC_DeleteAccessControl_Handler,
		},
		{
			MethodName: "GetIndexMeta",
			Handler:    _AdminRPC_GetIndexMeta_Handler,
		},
		{
			MethodName: "UpsertIndexMeta",
			Handler:    _AdminRPC_UpsertIndexMeta_Handler,
		},
		{
			MethodName: "DeleteIndexMeta",
			Handler:    _AdminRPC_DeleteIndexMeta_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}