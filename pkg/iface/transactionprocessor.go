package iface

import (
	"context"
	"github.com/mkawserm/flamed/pkg/pb"
)

type ITransactionProcessor interface {
	FamilyName() string
	FamilyVersion() string
	IndexObject(statePayload []byte) interface{}

	Search(ctx context.Context,
		readOnlyStateContext IStateContext,
		searchInput *pb.SearchInput) (interface{}, error)
	Iterate(ctx context.Context,
		readOnlyStateContext IStateContext,
		iterateInput *pb.IterateInput) (interface{}, error)
	Retrieve(ctx context.Context,
		readOnlyStateContext IStateContext,
		retrieveInput *pb.RetrieveInput) (interface{}, error)

	Apply(ctx context.Context,
		stateContext IStateContext,
		transaction *pb.Transaction) *pb.TransactionResponse
}
