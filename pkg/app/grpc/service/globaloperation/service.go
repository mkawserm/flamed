package globaloperation

import (
	"bytes"
	"context"
	"github.com/google/uuid"
	flamedContext "github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type Service struct {
	mFlamedContext *flamedContext.FlamedContext
	UnimplementedGlobalOperationRPCServer
}

func (s *Service) RegisterGRPCService(server *grpc.Server) {
	RegisterGlobalOperationRPCServer(server, s)
}

func (s *Service) Propose(ctx context.Context, req *ProposalRequest) (*pb.ProposalResponse, error) {
	if bytes.EqualFold(req.Namespace, []byte("meta")) {
		return nil, x.ErrMetaNamespaceIsReserved
	}

	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, x.ErrAccessDenied
	}

	authContext := flamedContext.AuthContext{}
	authContext.Protocol = "GRPC"
	authContext.KVPair = headers

	if !s.mFlamedContext.Flamed().IsClusterIDAvailable(req.ClusterID) {
		return nil, x.ErrClusterIsNotAvailable
	}

	admin := s.mFlamedContext.Flamed().NewAdmin(
		req.ClusterID,
		s.mFlamedContext.GlobalRequestTimeout())

	if !authContext.Authenticate(admin) {
		return nil, x.ErrAccessDenied
	}

	username := authContext.GetUsernameFromAuth()

	accessControl, err := admin.GetAccessControl(username, req.Namespace)
	if err != nil {
		accessControl, err = admin.GetAccessControl(username, []byte("*"))
		if err != nil {
			return nil, x.ErrAccessDenied
		}
	}

	if !utility.HasGlobalCRUDPermission(accessControl) {
		return nil, x.ErrAccessDenied
	}

	globalOperation := s.mFlamedContext.Flamed().NewGlobalOperation(req.ClusterID,
		req.Namespace,
		s.mFlamedContext.GlobalRequestTimeout())

	for _, t := range req.Proposal.Transactions {
		t.Namespace = req.Namespace

		if len(t.FamilyName) == 0 || len(t.FamilyVersion) == 0 || len(t.Payload) == 0 {
			return nil, x.ErrInvalidProposal
		}
	}

	if len(req.Proposal.Uuid) == 0 {
		uuidValue := uuid.New()
		req.Proposal.Uuid = uuidValue[:]
	}

	if req.Proposal.CreatedAt == 0 {
		req.Proposal.CreatedAt = uint64(time.Now().UnixNano())
	}

	return globalOperation.Propose(req.Proposal)
}

func NewGlobalOperationService(flamedContext *flamedContext.FlamedContext) *Service {
	return &Service{
		mFlamedContext: flamedContext,
	}
}
