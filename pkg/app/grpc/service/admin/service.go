package admin

import (
	"context"
	flamedContext "github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/x"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Service struct {
	mFlamedContext *flamedContext.FlamedContext
	UnimplementedAdminRPCServer
}

func (s *Service) RegisterGRPCService(server *grpc.Server) {
	RegisterAdminRPCServer(server, s)
}

func (s *Service) GetUser(ctx context.Context, req *UserRequest) (*pb.User, error) {
	if headers, ok := metadata.FromIncomingContext(ctx); ok {
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
	} else {
		return nil, x.ErrAccessDenied
	}
}

func NewAdminService(flamedContext *flamedContext.FlamedContext) *Service {
	return &Service{
		mFlamedContext: flamedContext,
	}
}
