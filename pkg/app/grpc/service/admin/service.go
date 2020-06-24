package admin

import (
	"context"
	"errors"
	flamedContext "github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/crypto"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variable"
	"github.com/mkawserm/flamed/pkg/x"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type Service struct {
	mFlamedContext *flamedContext.FlamedContext
	UnimplementedAdminRPCServer
}

func (s *Service) RegisterGRPCService(server *grpc.Server) {
	RegisterAdminRPCServer(server, s)
}

func (s *Service) GetUser(ctx context.Context, req *UserRequest) (*pb.User, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, x.ErrAccessDenied
	}

	authContext := flamedContext.AuthContext{}
	authContext.Protocol = "GRPC"
	authContext.KVPair = headers

	if !authContext.AuthenticateSuperUser(s.mFlamedContext.Flamed().NewAdmin(
		1,
		s.mFlamedContext.GlobalRequestTimeout())) {
		return nil, x.ErrAccessDenied
	}

	if !s.mFlamedContext.Flamed().IsClusterIDAvailable(req.ClusterID) {
		return nil, x.ErrClusterIsNotAvailable
	}

	admin := s.mFlamedContext.Flamed().NewAdmin(req.ClusterID, s.mFlamedContext.GlobalRequestTimeout())
	return admin.GetUser(req.Username)
}

func (s *Service) UpsertUser(ctx context.Context, req *UpsertUserRequest) (*pb.ProposalResponse, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, x.ErrAccessDenied
	}

	authContext := flamedContext.AuthContext{}
	authContext.Protocol = "GRPC"
	authContext.KVPair = headers

	if !authContext.AuthenticateSuperUser(s.mFlamedContext.Flamed().NewAdmin(
		1,
		s.mFlamedContext.GlobalRequestTimeout())) {
		return nil, x.ErrAccessDenied
	}

	if !s.mFlamedContext.Flamed().IsClusterIDAvailable(req.ClusterID) {
		return nil, x.ErrClusterIsNotAvailable
	}

	if req.User.Username == "admin" {
		return nil, errors.New("upsert operation is not allowed on admin user")
	}

	if len(req.User.Password) == 0 {
		return nil, errors.New("password can not be empty")
	}

	pha := variable.DefaultPasswordHashAlgorithmFactory
	if !pha.IsAlgorithmAvailable(variable.DefaultPasswordHashAlgorithm) {
		return nil, errors.New(variable.DefaultPasswordHashAlgorithm + " is not available")
	}

	encoded, err := pha.MakePassword(req.User.Password,
		crypto.GetRandomString(12),
		variable.DefaultPasswordHashAlgorithm)

	if err != nil {
		return nil, err
	}

	req.User.Password = encoded
	admin := s.mFlamedContext.Flamed().NewAdmin(req.ClusterID, s.mFlamedContext.GlobalRequestTimeout())
	return admin.UpsertUser(req.User)
}

func (s *Service) ChangeUserPassword(ctx context.Context, req *ChangeUserPasswordRequest) (*pb.ProposalResponse, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, x.ErrAccessDenied
	}

	authContext := flamedContext.AuthContext{}
	authContext.Protocol = "GRPC"
	authContext.KVPair = headers

	if !authContext.AuthenticateSuperUser(s.mFlamedContext.Flamed().NewAdmin(
		1,
		s.mFlamedContext.GlobalRequestTimeout())) {
		return nil, x.ErrAccessDenied
	}

	if !s.mFlamedContext.Flamed().IsClusterIDAvailable(req.ClusterID) {
		return nil, x.ErrClusterIsNotAvailable
	}

	admin := s.mFlamedContext.Flamed().NewAdmin(req.ClusterID, s.mFlamedContext.GlobalRequestTimeout())
	user, err := admin.GetUser(req.Username)
	if err != nil {
		return nil, err
	}

	pha := variable.DefaultPasswordHashAlgorithmFactory
	if !pha.IsAlgorithmAvailable(variable.DefaultPasswordHashAlgorithm) {
		return nil, errors.New(variable.DefaultPasswordHashAlgorithm + " is not available")
	}

	encoded, err := pha.MakePassword(req.Password,
		crypto.GetRandomString(12),
		variable.DefaultPasswordHashAlgorithm)

	if err != nil {
		return nil, err
	}

	user.Password = encoded
	user.UpdatedAt = uint64(time.Now().UnixNano())
	return admin.UpsertUser(user)
}

func (s *Service) DeleteUser(ctx context.Context, req *UserRequest) (*pb.ProposalResponse, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, x.ErrAccessDenied
	}

	authContext := flamedContext.AuthContext{}
	authContext.Protocol = "GRPC"
	authContext.KVPair = headers

	if !authContext.AuthenticateSuperUser(s.mFlamedContext.Flamed().NewAdmin(
		1,
		s.mFlamedContext.GlobalRequestTimeout())) {
		return nil, x.ErrAccessDenied
	}

	if !s.mFlamedContext.Flamed().IsClusterIDAvailable(req.ClusterID) {
		return nil, x.ErrClusterIsNotAvailable
	}

	admin := s.mFlamedContext.Flamed().NewAdmin(req.ClusterID, s.mFlamedContext.GlobalRequestTimeout())
	return admin.DeleteUser(req.Username)
}

func (s *Service) GetAccessControl(ctx context.Context, req *AccessControlRequest) (*pb.AccessControl, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, x.ErrAccessDenied
	}

	authContext := flamedContext.AuthContext{}
	authContext.Protocol = "GRPC"
	authContext.KVPair = headers

	if !authContext.AuthenticateSuperUser(s.mFlamedContext.Flamed().NewAdmin(
		1,
		s.mFlamedContext.GlobalRequestTimeout())) {
		return nil, x.ErrAccessDenied
	}

	if !s.mFlamedContext.Flamed().IsClusterIDAvailable(req.ClusterID) {
		return nil, x.ErrClusterIsNotAvailable
	}

	admin := s.mFlamedContext.Flamed().NewAdmin(req.ClusterID, s.mFlamedContext.GlobalRequestTimeout())

	return admin.GetAccessControl(req.Username, req.Namespace)
}

func (s *Service) UpsertAccessControl(ctx context.Context, req *UpsertAccessControlRequest) (*pb.ProposalResponse, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, x.ErrAccessDenied
	}

	authContext := flamedContext.AuthContext{}
	authContext.Protocol = "GRPC"
	authContext.KVPair = headers

	if !authContext.AuthenticateSuperUser(s.mFlamedContext.Flamed().NewAdmin(
		1,
		s.mFlamedContext.GlobalRequestTimeout())) {
		return nil, x.ErrAccessDenied
	}

	if !s.mFlamedContext.Flamed().IsClusterIDAvailable(req.ClusterID) {
		return nil, x.ErrClusterIsNotAvailable
	}

	admin := s.mFlamedContext.Flamed().NewAdmin(req.ClusterID, s.mFlamedContext.GlobalRequestTimeout())

	return admin.UpsertAccessControl(req.AccessControl)
}

func (s *Service) DeleteAccessControl(ctx context.Context, req *AccessControlRequest) (*pb.ProposalResponse, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, x.ErrAccessDenied
	}

	authContext := flamedContext.AuthContext{}
	authContext.Protocol = "GRPC"
	authContext.KVPair = headers

	if !authContext.AuthenticateSuperUser(s.mFlamedContext.Flamed().NewAdmin(
		1,
		s.mFlamedContext.GlobalRequestTimeout())) {
		return nil, x.ErrAccessDenied
	}

	if !s.mFlamedContext.Flamed().IsClusterIDAvailable(req.ClusterID) {
		return nil, x.ErrClusterIsNotAvailable
	}

	admin := s.mFlamedContext.Flamed().NewAdmin(req.ClusterID, s.mFlamedContext.GlobalRequestTimeout())
	return admin.DeleteAccessControl(req.Namespace, req.Username)
}

func (s *Service) GetIndexMeta(ctx context.Context, req *IndexMetaRequest) (*pb.IndexMeta, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, x.ErrAccessDenied
	}

	authContext := flamedContext.AuthContext{}
	authContext.Protocol = "GRPC"
	authContext.KVPair = headers

	if !authContext.AuthenticateSuperUser(s.mFlamedContext.Flamed().NewAdmin(
		1,
		s.mFlamedContext.GlobalRequestTimeout())) {
		return nil, x.ErrAccessDenied
	}

	if !s.mFlamedContext.Flamed().IsClusterIDAvailable(req.ClusterID) {
		return nil, x.ErrClusterIsNotAvailable
	}

	admin := s.mFlamedContext.Flamed().NewAdmin(req.ClusterID, s.mFlamedContext.GlobalRequestTimeout())

	return admin.GetIndexMeta(req.Namespace)
}

func (s *Service) UpsertIndexMeta(ctx context.Context, req *UpsertIndexMetaRequest) (*pb.ProposalResponse, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, x.ErrAccessDenied
	}

	authContext := flamedContext.AuthContext{}
	authContext.Protocol = "GRPC"
	authContext.KVPair = headers

	if !authContext.AuthenticateSuperUser(s.mFlamedContext.Flamed().NewAdmin(
		1,
		s.mFlamedContext.GlobalRequestTimeout())) {
		return nil, x.ErrAccessDenied
	}

	if !s.mFlamedContext.Flamed().IsClusterIDAvailable(req.ClusterID) {
		return nil, x.ErrClusterIsNotAvailable
	}

	admin := s.mFlamedContext.Flamed().NewAdmin(req.ClusterID, s.mFlamedContext.GlobalRequestTimeout())

	return admin.UpsertIndexMeta(req.IndexMeta)
}

func (s *Service) DeleteIndexMeta(ctx context.Context, req *IndexMetaRequest) (*pb.ProposalResponse, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, x.ErrAccessDenied
	}

	authContext := flamedContext.AuthContext{}
	authContext.Protocol = "GRPC"
	authContext.KVPair = headers

	if !authContext.AuthenticateSuperUser(s.mFlamedContext.Flamed().NewAdmin(
		1,
		s.mFlamedContext.GlobalRequestTimeout())) {
		return nil, x.ErrAccessDenied
	}

	if !s.mFlamedContext.Flamed().IsClusterIDAvailable(req.ClusterID) {
		return nil, x.ErrClusterIsNotAvailable
	}

	admin := s.mFlamedContext.Flamed().NewAdmin(req.ClusterID, s.mFlamedContext.GlobalRequestTimeout())

	return admin.DeleteIndexMeta(req.Namespace)
}

func NewAdminService(flamedContext *flamedContext.FlamedContext) *Service {
	return &Service{
		mFlamedContext: flamedContext,
	}
}
