package iface

import (
	"context"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
)

type ITransactionProcessor interface {
	Family() string
	Version() string
	Lookup(readOnlyStateContext IStateContext, request variant.LookupRequest) (interface{}, error)
	Apply(ctx context.Context,
		stateContext IStateContext,
		transaction *pb.Transaction) *variant.TransactionProcessorResponse
}
