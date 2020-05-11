package iface

import (
	"context"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
)

type ITransactionProcessor interface {
	Family() string
	Version() string
	Apply(context *context.Context,
		stateContext IStateContext,
		transaction *pb.Transaction) variant.TransactionProcessorResponse
}
