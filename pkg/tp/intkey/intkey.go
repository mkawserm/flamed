package intkey

import (
	"bytes"
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/crypto"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/x"
)

type IntKey struct {
}

func (i *IntKey) FamilyName() string {
	return Name
}

func (i *IntKey) FamilyVersion() string {
	return Version
}

func (i *IntKey) IndexObject(_ []byte) interface{} {
	return nil
}

func (i *IntKey) Search(_ context.Context,
	_ iface.IStateContext,
	_ *pb.SearchInput) (interface{}, error) {
	return nil, x.ErrNotImplemented
}

func (i *IntKey) Iterate(_ context.Context,
	_ iface.IStateContext,
	_ *pb.IterateInput) (interface{}, error) {
	return nil, x.ErrNotImplemented
}

func (i *IntKey) Retrieve(_ context.Context,
	readOnlyStateContext iface.IStateContext,
	retrieveInput *pb.RetrieveInput) (interface{}, error) {
	if len(retrieveInput.Addresses) == 0 {
		return nil, nil
	}

	intKeyStates := make([]*IntKeyState, 0, len(retrieveInput.Addresses))

	for _, sa := range retrieveInput.Addresses {
		if !bytes.HasPrefix(sa, retrieveInput.Namespace) {
			return nil, x.ErrAccessViolation
		}

		entry, err := readOnlyStateContext.GetState(sa)

		if err != nil {
			intKeyStates = append(intKeyStates, nil)
			continue
		}

		a := &IntKeyState{}
		if err := proto.Unmarshal(entry.Payload, a); err != nil {
			return nil, err
		}
		intKeyStates = append(intKeyStates, a)
	}

	return intKeyStates, nil
}

func (i *IntKey) insert(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	transaction *pb.Transaction,
	address []byte,
	payload *IntKeyPayload) *pb.TransactionResponse {

	entry, _ := stateContext.GetState(address)
	if entry != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 6
		tpr.ErrorText = "state already exists so can not insert"
		return tpr
	}

	stateData := &IntKeyState{
		Name:  payload.Name,
		Value: payload.Value,
	}

	stateBytes, err := proto.Marshal(stateData)

	if err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 7
		tpr.ErrorText = err.Error()
		return tpr
	}

	entry = &pb.StateEntry{
		Payload:       stateBytes,
		Namespace:     transaction.Namespace,
		FamilyName:    transaction.FamilyName,
		FamilyVersion: transaction.FamilyVersion,
	}

	if err := stateContext.UpsertState(address, entry); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 8
		tpr.ErrorText = err.Error()
		return tpr
	} else {
		tpr.Status = pb.Status_ACCEPTED
		tpr.ErrorCode = 0
		tpr.ErrorText = ""
		return tpr
	}
}

func (i *IntKey) delete(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext, address []byte) *pb.TransactionResponse {

	entry, _ := stateContext.GetState(address)
	if entry == nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 9
		tpr.ErrorText = "state does not exists, can not delete"
		return tpr
	}

	if err := stateContext.DeleteState(address); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 10
		tpr.ErrorText = err.Error()
		return tpr
	} else {
		tpr.Status = pb.Status_ACCEPTED
		tpr.ErrorCode = 0
		tpr.ErrorText = ""
		return tpr
	}
}

func (i *IntKey) upsert(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	transaction *pb.Transaction,
	address []byte,
	payload *IntKeyPayload) *pb.TransactionResponse {

	stateData := &IntKeyState{
		Name:  payload.Name,
		Value: payload.Value,
	}

	stateBytes, err := proto.Marshal(stateData)

	if err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 11
		tpr.ErrorText = err.Error()
		return tpr
	}

	entry := &pb.StateEntry{
		Payload:       stateBytes,
		Namespace:     transaction.Namespace,
		FamilyName:    transaction.FamilyName,
		FamilyVersion: transaction.FamilyVersion,
	}

	if err := stateContext.UpsertState(address, entry); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 12
		tpr.ErrorText = err.Error()
		return tpr
	} else {
		tpr.Status = pb.Status_ACCEPTED
		tpr.ErrorCode = 0
		tpr.ErrorText = ""
		return tpr
	}
}

func (i *IntKey) increment(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	address []byte,
	payload *IntKeyPayload) *pb.TransactionResponse {

	entry, _ := stateContext.GetState(address)
	if entry == nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 13
		tpr.ErrorText = "state does not exists so can not increment"
		return tpr
	}

	stateData := &IntKeyState{}

	if err := proto.Unmarshal(entry.Payload, stateData); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 14
		tpr.ErrorText = err.Error()
		return tpr
	}

	if 18446744073709551615-payload.Value < stateData.Value {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 15
		tpr.ErrorText = "result can not be out of range"
		return tpr
	}
	stateData.Value = stateData.Value + payload.Value
	stateBytes, err := proto.Marshal(stateData)
	if err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 16
		tpr.ErrorText = err.Error()
		return tpr
	}

	entry.Payload = stateBytes

	if err := stateContext.UpsertState(address, entry); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 17
		tpr.ErrorText = err.Error()
		return tpr
	} else {
		tpr.Status = pb.Status_ACCEPTED
		tpr.ErrorCode = 0
		tpr.ErrorText = ""
		return tpr
	}
}

func (i *IntKey) decrement(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	address []byte,
	payload *IntKeyPayload) *pb.TransactionResponse {

	entry, _ := stateContext.GetState(address)
	if entry == nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 18
		tpr.ErrorText = "state does not exists can not decrement"
		return tpr
	}

	stateData := &IntKeyState{}

	if err := proto.Unmarshal(entry.Payload, stateData); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 19
		tpr.ErrorText = err.Error()
		return tpr
	}

	if payload.Value > stateData.Value {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 20
		tpr.ErrorText = "result can not be out of range"
		return tpr
	}

	stateData.Value = stateData.Value - payload.Value

	stateBytes, err := proto.Marshal(stateData)
	if err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 21
		tpr.ErrorText = err.Error()
		return tpr
	}

	entry.Payload = stateBytes

	if err := stateContext.UpsertState(address, entry); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 22
		tpr.ErrorText = err.Error()
		return tpr
	} else {
		tpr.Status = pb.Status_ACCEPTED
		tpr.ErrorCode = 0
		tpr.ErrorText = ""
		return tpr
	}
}

func (i *IntKey) Apply(_ context.Context,
	stateContext iface.IStateContext,
	transaction *pb.Transaction) *pb.TransactionResponse {

	tpr := &pb.TransactionResponse{
		Status:        pb.Status_REJECTED,
		ErrorCode:     1,
		ErrorText:     "",
		FamilyName:    Name,
		FamilyVersion: Version,
	}

	payload := &IntKeyPayload{}

	if err := proto.Unmarshal(transaction.Payload, payload); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 2
		tpr.ErrorText = err.Error()
		return tpr
	}

	if len(payload.Name) == 0 {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 3
		tpr.ErrorText = "name can not be empty"
		return tpr
	}

	if len(payload.Name) > 20 {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 4
		tpr.ErrorText = "name length can not exceed 20 characters"
		return tpr
	}

	hash := crypto.GetStateHashFromStringKey(i.FamilyName(), payload.Name)
	address := crypto.GetStateAddress(transaction.Namespace, hash)

	if payload.Verb == Verb_INSERT {
		return i.insert(tpr, stateContext, transaction, address, payload)
	} else if payload.Verb == Verb_UPSERT {
		return i.upsert(tpr, stateContext, transaction, address, payload)
	} else if payload.Verb == Verb_INCREMENT {
		return i.increment(tpr, stateContext, address, payload)
	} else if payload.Verb == Verb_DECREMENT {
		return i.decrement(tpr, stateContext, address, payload)
	} else if payload.Verb == Verb_DELETE {
		return i.delete(tpr, stateContext, address)
	} else {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 5
		tpr.ErrorText = "unknown verb"
		return tpr
	}
}
