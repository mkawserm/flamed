package iface

import (
	"context"
	"github.com/mkawserm/flamed/pkg/pb"
)

type ITransactionProcessor interface {
	FamilyName() string
	FamilyVersion() string
	Lookup(ctx context.Context, readOnlyStateContext IStateContext, query interface{}) (interface{}, error)
	Apply(ctx context.Context,
		stateContext IStateContext,
		transaction *pb.Transaction) *pb.TransactionProcessorResponse
}
