package index

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/constant"
	"github.com/mkawserm/flamed/pkg/uidutil"
	"github.com/mkawserm/flamed/pkg/utility"

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

func (i *Index) upsert(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	transaction *pb.Transaction,
	address []byte,
	indexMeta *pb.IndexMeta) *pb.TransactionResponse {

	payload, err := proto.Marshal(indexMeta)
	if err != nil {
		tpr.Status = 0
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
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	} else {
		tpr.Status = 1
		tpr.ErrorCode = 0
		tpr.ErrorText = ""
		return tpr
	}
}

func (i *Index) delete(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	address []byte) *pb.TransactionResponse {
	if err := stateContext.DeleteState(address); err != nil {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	} else {
		tpr.Status = 1
		tpr.ErrorCode = 0
		tpr.ErrorText = ""
		return tpr
	}
}

func (i *Index) Apply(_ context.Context,
	stateContext iface.IStateContext,
	transaction *pb.Transaction) *pb.TransactionResponse {

	tpr := &pb.TransactionResponse{
		Status:        0,
		ErrorCode:     0,
		ErrorText:     "",
		FamilyName:    Name,
		FamilyVersion: Version,
	}

	payload := &IndexPayload{}

	if err := proto.Unmarshal(transaction.Payload, payload); err != nil {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	}

	if payload.IndexMeta == nil {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = "index meta can not be nil"
		return tpr
	}

	if !utility.IsNamespaceValid(payload.IndexMeta.Namespace) {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = "invalid namespace"
		return tpr
	}

	address := uidutil.GetUid([]byte(constant.IndexMetaNamespace), payload.IndexMeta.Namespace)

	if payload.Action == Action_UPSERT {
		r := i.upsert(tpr, stateContext, transaction, address, payload.IndexMeta)
		if r.Status == 1 {
			_ = stateContext.UpsertIndexMeta(payload.IndexMeta)
		}
		return r
	} else if payload.Action == Action_DELETE {
		r := i.delete(tpr, stateContext, address)
		if r.Status == 1 {
			_ = stateContext.DeleteIndexMeta(payload.IndexMeta)
		}
		return r
	} else {
		tpr.Status = 0
		tpr.ErrorCode = 0
		tpr.ErrorText = "unknown action"
		return tpr
	}
}
