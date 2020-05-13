package intkey

import (
	"context"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
)

type IntKey struct {
}

func (i *IntKey) FamilyName() string {
	return "intkey"
}

func (i *IntKey) FamilyVersion() string {
	return "1.0"
}

func (i *IntKey) Lookup(_ context.Context,
	readOnlyStateContext iface.IStateContext,
	query interface{}) (interface{}, error) {

	return nil, nil
}

func (i *IntKey) Apply(_ context.Context,
	stateContext iface.IStateContext,
	transaction *pb.Transaction) *variant.TransactionProcessorResponse {

	return nil
}
