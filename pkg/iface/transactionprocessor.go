package iface

import "context"

type ITransactionProcessor interface {
	Family() string
	Version() string
	Apply(context *context.Context,
		stateContext IStateContext,
		transaction ITransaction) ITransactionProcessorResponse
}
