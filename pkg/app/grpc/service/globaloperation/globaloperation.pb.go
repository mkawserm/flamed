// Code generated by protoc-gen-go. DO NOT EDIT.
// source: globaloperation.proto

package globaloperation

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

type ProposalRequest struct {
	ClusterID            uint64       `protobuf:"varint,1,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
	Namespace            []byte       `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Proposal             *pb.Proposal `protobuf:"bytes,3,opt,name=proposal,proto3" json:"proposal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ProposalRequest) Reset()         { *m = ProposalRequest{} }
func (m *ProposalRequest) String() string { return proto.CompactTextString(m) }
func (*ProposalRequest) ProtoMessage()    {}
func (*ProposalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c83549202f90941, []int{0}
}

func (m *ProposalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalRequest.Unmarshal(m, b)
}
func (m *ProposalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalRequest.Marshal(b, m, deterministic)
}
func (m *ProposalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalRequest.Merge(m, src)
}
func (m *ProposalRequest) XXX_Size() int {
	return xxx_messageInfo_ProposalRequest.Size(m)
}
func (m *ProposalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalRequest proto.InternalMessageInfo

func (m *ProposalRequest) GetClusterID() uint64 {
	if m != nil {
		return m.ClusterID
	}
	return 0
}

func (m *ProposalRequest) GetNamespace() []byte {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ProposalRequest) GetProposal() *pb.Proposal {
	if m != nil {
		return m.Proposal
	}
	return nil
}

func init() {
	proto.RegisterType((*ProposalRequest)(nil), "globaloperation.ProposalRequest")
}

func init() { proto.RegisterFile("globaloperation.proto", fileDescriptor_2c83549202f90941) }

var fileDescriptor_2c83549202f90941 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x8d, 0x8a, 0x3f, 0xb1, 0x50, 0x08, 0x0a, 0x43, 0x71, 0x11, 0xba, 0xca, 0x2a, 0x81,
	0xfa, 0x08, 0x15, 0xc4, 0x95, 0x25, 0x3b, 0x97, 0xc9, 0x78, 0x1d, 0x4b, 0x93, 0xb9, 0xd7, 0x24,
	0x43, 0x5f, 0x5f, 0xec, 0x8c, 0x33, 0x38, 0xcb, 0x9c, 0x13, 0xbe, 0xfb, 0x1d, 0xfe, 0xd0, 0x04,
	0xf4, 0x2e, 0x20, 0x41, 0x72, 0x65, 0x8f, 0xad, 0xa6, 0x84, 0x05, 0xc5, 0x72, 0x16, 0xaf, 0x74,
	0xb3, 0x2f, 0x5f, 0x9d, 0xd7, 0x35, 0x46, 0x13, 0x0f, 0xee, 0x98, 0x21, 0x45, 0xf3, 0x19, 0x5c,
	0x84, 0x0f, 0x43, 0x87, 0xc6, 0x90, 0x1f, 0x5e, 0x3d, 0x60, 0x7d, 0xe4, 0xcb, 0x5d, 0x42, 0xc2,
	0xec, 0x82, 0x85, 0xef, 0x0e, 0x72, 0x11, 0x8f, 0xfc, 0xb6, 0x0e, 0x5d, 0x2e, 0x90, 0x5e, 0x9f,
	0x2b, 0x26, 0x99, 0xba, 0xb4, 0x53, 0xf0, 0xdb, 0xb6, 0x2e, 0x42, 0x26, 0x57, 0x43, 0x75, 0x2e,
	0x99, 0x5a, 0xd8, 0x29, 0x10, 0x8a, 0xdf, 0xd0, 0x80, 0xab, 0x2e, 0x24, 0x53, 0x77, 0x9b, 0x85,
	0x26, 0xaf, 0xc7, 0x13, 0x63, 0xbb, 0x79, 0xe7, 0xe2, 0xe5, 0xe4, 0xfe, 0xf6, 0xe7, 0x6e, 0x77,
	0x5b, 0xb1, 0xe5, 0xd7, 0xfd, 0x5f, 0x10, 0x52, 0xcf, 0x27, 0xcf, 0x44, 0x57, 0xf7, 0xff, 0xd0,
	0x90, 0x09, 0xdb, 0x0c, 0xeb, 0x33, 0x7f, 0x75, 0x9a, 0xf6, 0xf4, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x14, 0x59, 0x75, 0x18, 0x34, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GlobalOperationRPCClient is the client API for GlobalOperationRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GlobalOperationRPCClient interface {
	Propose(ctx context.Context, in *ProposalRequest, opts ...grpc.CallOption) (*pb.ProposalResponse, error)
}

type globalOperationRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewGlobalOperationRPCClient(cc grpc.ClientConnInterface) GlobalOperationRPCClient {
	return &globalOperationRPCClient{cc}
}

func (c *globalOperationRPCClient) Propose(ctx context.Context, in *ProposalRequest, opts ...grpc.CallOption) (*pb.ProposalResponse, error) {
	out := new(pb.ProposalResponse)
	err := c.cc.Invoke(ctx, "/globaloperation.GlobalOperationRPC/Propose", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GlobalOperationRPCServer is the server API for GlobalOperationRPC service.
type GlobalOperationRPCServer interface {
	Propose(context.Context, *ProposalRequest) (*pb.ProposalResponse, error)
}

// UnimplementedGlobalOperationRPCServer can be embedded to have forward compatible implementations.
type UnimplementedGlobalOperationRPCServer struct {
}

func (*UnimplementedGlobalOperationRPCServer) Propose(ctx context.Context, req *ProposalRequest) (*pb.ProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Propose not implemented")
}

func RegisterGlobalOperationRPCServer(s *grpc.Server, srv GlobalOperationRPCServer) {
	s.RegisterService(&_GlobalOperationRPC_serviceDesc, srv)
}

func _GlobalOperationRPC_Propose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalOperationRPCServer).Propose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/globaloperation.GlobalOperationRPC/Propose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalOperationRPCServer).Propose(ctx, req.(*ProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GlobalOperationRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "globaloperation.GlobalOperationRPC",
	HandlerType: (*GlobalOperationRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Propose",
			Handler:    _GlobalOperationRPC_Propose_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "globaloperation.proto",
}
