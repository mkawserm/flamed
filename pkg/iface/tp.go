package iface

import (
	"context"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
)

type ITransactionProcessor interface {
	Family() string
	Version() string
	Lookup(ctx context.Context, readOnlyStateContext IStateContext, query interface{}) (interface{}, error)
	Apply(ctx context.Context,
		stateContext IStateContext,
		transaction *pb.Transaction) *variant.TransactionProcessorResponse
}
