package index

import (
	"context"

	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
)

type Index struct {
}

func (i *Index) FamilyName() string {
	return Name
}

func (i *Index) FamilyVersion() string {
	return Version
}

func (i *Index) Lookup(_ context.Context,
	readOnlyStateContext iface.IStateContext,
	query interface{}) (interface{}, error) {

	return nil, nil
}

func (i *Index) Apply(_ context.Context,
	stateContext iface.IStateContext,
	transaction *pb.Transaction) *pb.TransactionResponse {

	return nil
}
