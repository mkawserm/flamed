package indexmeta

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/constant"
	"github.com/mkawserm/flamed/pkg/crypto"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"

	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
)

type IndexMeta struct {
}

func (i *IndexMeta) FamilyName() string {
	return Name
}

func (i *IndexMeta) FamilyVersion() string {
	return Version
}

func (i *IndexMeta) IndexObject(_ []byte) interface{} {
	return nil
}

func (i *IndexMeta) Search(_ context.Context,
	_ iface.IStateContext,
	_ *pb.SearchInput) (interface{}, error) {
	return nil, x.ErrNotImplemented
}

func (i *IndexMeta) Iterate(_ context.Context,
	_ iface.IStateContext,
	_ *pb.IterateInput) (interface{}, error) {
	return nil, x.ErrNotImplemented
}

func (i *IndexMeta) Retrieve(ctx context.Context,
	readOnlyStateContext iface.IStateContext,
	retrieveInput *pb.RetrieveInput) (interface{}, error) {
	return nil, nil
}

func (i *IndexMeta) Lookup(_ context.Context,
	readOnlyStateContext iface.IStateContext,
	query interface{}) (interface{}, error) {

	var namespace []byte

	if v, ok := query.([]byte); ok {
		namespace = v
	} else {
		return nil, x.ErrInvalidLookupInput
	}
	address := crypto.GetStateAddress([]byte(constant.IndexMetaNamespace), namespace)

	entry, err := readOnlyStateContext.GetState(address)

	if err != nil {
		return nil, err
	}

	indexMeta := &pb.IndexMeta{}
	if err := proto.Unmarshal(entry.Payload, indexMeta); err != nil {
		return nil, err
	}

	return indexMeta, nil
}

func (i *IndexMeta) upsert(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	transaction *pb.Transaction,
	address []byte,
	indexMeta *pb.IndexMeta) *pb.TransactionResponse {

	payload, err := proto.Marshal(indexMeta)
	if err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	}

	entry := &pb.StateEntry{
		Payload:       payload,
		Namespace:     transaction.Namespace,
		FamilyName:    transaction.FamilyName,
		FamilyVersion: transaction.FamilyVersion,
	}

	if err := stateContext.UpsertState(address, entry); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	} else {
		tpr.Status = pb.Status_ACCEPTED
		tpr.ErrorCode = 0
		tpr.ErrorText = ""
		return tpr
	}
}

func (i *IndexMeta) delete(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	address []byte) *pb.TransactionResponse {

	if _, err := stateContext.GetState(address); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	}

	if err := stateContext.DeleteState(address); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	} else {
		tpr.Status = pb.Status_ACCEPTED
		tpr.ErrorCode = 0
		tpr.ErrorText = ""
		return tpr
	}
}

func (i *IndexMeta) Apply(_ context.Context,
	stateContext iface.IStateContext,
	transaction *pb.Transaction) *pb.TransactionResponse {

	tpr := &pb.TransactionResponse{
		Status:        pb.Status_REJECTED,
		ErrorCode:     0,
		ErrorText:     "",
		FamilyName:    Name,
		FamilyVersion: Version,
	}

	payload := &pb.IndexMetaPayload{}

	if err := proto.Unmarshal(transaction.Payload, payload); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	}

	if payload.IndexMeta == nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 0
		tpr.ErrorText = "indexmeta meta can not be nil"
		return tpr
	}

	if !utility.IsNamespaceValid(payload.IndexMeta.Namespace) {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 0
		tpr.ErrorText = "invalid namespace"
		return tpr
	}

	address := crypto.GetStateAddress([]byte(constant.IndexMetaNamespace), payload.IndexMeta.Namespace)

	if payload.Action == pb.Action_UPSERT {
		r := i.upsert(tpr, stateContext, transaction, address, payload.IndexMeta)
		if r.Status == pb.Status_ACCEPTED {
			_ = stateContext.UpsertIndexMeta(payload.IndexMeta)
		}
		return r
	} else if payload.Action == pb.Action_DELETE {
		r := i.delete(tpr, stateContext, address)
		if r.Status == pb.Status_ACCEPTED {
			_ = stateContext.DeleteIndexMeta(payload.IndexMeta)
		}
		return r
	} else {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 0
		tpr.ErrorText = "unknown action"
		return tpr
	}
}
