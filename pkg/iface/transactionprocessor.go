package iface

import "context"

type ITransactionProcessor interface {
	Family() string
	Version() string
	Apply(context *context.Context,
		stateStorageContext IStateStorageContext,
		transaction ITransaction) ITransactionProcessorResponse
}
