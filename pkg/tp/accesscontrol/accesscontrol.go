package accesscontrol

import (
	"bytes"
	"context"
	"github.com/mkawserm/flamed/pkg/crypto"

	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/constant"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"
)

type AccessControl struct {
}

func (c *AccessControl) FamilyName() string {
	return Name
}

func (c *AccessControl) FamilyVersion() string {
	return Version
}

func (c *AccessControl) IndexObject(_ []byte) interface{} {
	return nil
}

func (c *AccessControl) Search(_ context.Context,
	_ iface.IStateContext,
	_ *pb.SearchInput) (interface{}, error) {
	return nil, x.ErrNotImplemented
}

func (c *AccessControl) Iterate(_ context.Context,
	_ iface.IStateContext,
	_ *pb.IterateInput) (interface{}, error) {
	return nil, x.ErrNotImplemented
}

func (c *AccessControl) Retrieve(ctx context.Context,
	readOnlyStateContext iface.IStateContext,
	retrieveInput *pb.RetrieveInput) (interface{}, error) {
	return nil, nil
}

func (c *AccessControl) Lookup(_ context.Context,
	readOnlyStateContext iface.IStateContext,
	query interface{}) (interface{}, error) {

	var request Request

	if v, ok := query.(Request); ok {
		request = v
	} else {
		return nil, x.ErrInvalidLookupInput
	}

	if !bytes.Equal(request.Namespace, []byte("*")) {
		if !utility.IsNamespaceValid(request.Namespace) {
			return nil, x.ErrInvalidNamespace
		}
	}

	if !utility.IsUsernameValid(request.Username) {
		return nil, x.ErrInvalidUsername
	}

	address := crypto.GetStateAddress([]byte(constant.AccessControlNamespace),
		crypto.GetStateAddress([]byte(request.Username), request.Namespace))

	entry, err := readOnlyStateContext.GetState(address)

	if err != nil {
		return nil, err
	}

	ac := &pb.AccessControl{}
	if err := proto.Unmarshal(entry.Payload, ac); err != nil {
		return nil, err
	}

	return ac, nil
}

func (c *AccessControl) upsert(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	transaction *pb.Transaction,
	address []byte,
	accessControl *pb.AccessControl) *pb.TransactionResponse {

	payload, err := proto.Marshal(accessControl)
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

func (c *AccessControl) delete(tpr *pb.TransactionResponse,
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

func (c *AccessControl) Apply(_ context.Context,
	stateContext iface.IStateContext,
	transaction *pb.Transaction) *pb.TransactionResponse {

	tpr := &pb.TransactionResponse{
		Status:        pb.Status_REJECTED,
		ErrorCode:     0,
		ErrorText:     "",
		FamilyName:    Name,
		FamilyVersion: Version,
	}

	payload := &pb.AccessControlPayload{}

	if err := proto.Unmarshal(transaction.Payload, payload); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	}

	if payload.AccessControl == nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 0
		tpr.ErrorText = "access control can not be nil"
		return tpr
	}

	if !bytes.Equal(payload.AccessControl.Namespace, []byte("*")) {
		if !utility.IsNamespaceValid(payload.AccessControl.Namespace) {
			tpr.Status = pb.Status_REJECTED
			tpr.ErrorCode = 0
			tpr.ErrorText = "invalid namespace"
			return tpr
		}
	}

	if !utility.IsUsernameValid(payload.AccessControl.Username) {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 0
		tpr.ErrorText = "invalid username"
		return tpr
	}

	address := crypto.GetStateAddress([]byte(constant.AccessControlNamespace),
		crypto.GetStateAddress([]byte(payload.AccessControl.Username), payload.AccessControl.Namespace))

	if payload.Action == pb.Action_UPSERT {
		return c.upsert(tpr, stateContext, transaction, address, payload.AccessControl)
	} else if payload.Action == pb.Action_DELETE {
		return c.delete(tpr, stateContext, address)
	} else {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 0
		tpr.ErrorText = "unknown action"
		return tpr
	}
}
