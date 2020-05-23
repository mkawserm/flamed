package context

import (
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/iface"
)

type FlamedContext struct {
	Flamed                       *flamed.Flamed
	TransactionProcessorMap      map[string]iface.ITransactionProcessor
	PasswordHashAlgorithmFactory iface.IPasswordHashAlgorithmFactory
}

func NewFlamedContext() *FlamedContext {
	return &FlamedContext{TransactionProcessorMap: map[string]iface.ITransactionProcessor{}}
}
