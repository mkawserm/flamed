package user

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

type User struct {
}

func (u *User) FamilyName() string {
	return Name
}

func (u *User) FamilyVersion() string {
	return Version
}

func (u *User) IndexObject(_ []byte) interface{} {
	return nil
}

func (u *User) Search(_ context.Context,
	_ iface.IStateContext,
	_ *pb.SearchInput) (interface{}, error) {
	return nil, x.ErrNotImplemented
}

func (u *User) Iterate(_ context.Context,
	_ iface.IStateContext,
	_ *pb.IterateInput) (interface{}, error) {
	return nil, x.ErrNotImplemented
}

func (u *User) Retrieve(_ context.Context,
	readOnlyStateContext iface.IStateContext,
	retrieveInput *pb.RetrieveInput) (interface{}, error) {
	if len(retrieveInput.Addresses) == 0 {
		return nil, nil
	}

	users := make([]*pb.User, 0, 1)

	for _, sa := range retrieveInput.Addresses {
		if !bytes.HasPrefix(sa, retrieveInput.Namespace) {
			return nil, x.ErrAccessViolation
		}

		entry, err := readOnlyStateContext.GetState(sa)

		if err != nil {
			users = append(users, nil)
			continue
		}

		a := &pb.User{}
		if err := proto.Unmarshal(entry.Payload, a); err != nil {
			return nil, err
		}

		users = append(users, a)
	}

	return users, nil
}

func (u *User) upsert(tpr *pb.TransactionResponse,
	stateContext iface.IStateContext,
	transaction *pb.Transaction,
	address []byte,
	user *pb.User) *pb.TransactionResponse {

	payload, err := proto.Marshal(user)
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

func (u *User) delete(tpr *pb.TransactionResponse,
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

func (u *User) Apply(_ context.Context,
	stateContext iface.IStateContext,
	transaction *pb.Transaction) *pb.TransactionResponse {

	tpr := &pb.TransactionResponse{
		Status:        pb.Status_REJECTED,
		ErrorCode:     0,
		ErrorText:     "",
		FamilyName:    Name,
		FamilyVersion: Version,
	}

	payload := &pb.UserPayload{}

	if err := proto.Unmarshal(transaction.Payload, payload); err != nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 0
		tpr.ErrorText = err.Error()
		return tpr
	}

	if payload.User == nil {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 0
		tpr.ErrorText = "user can not be nil"
		return tpr
	}

	if !utility.IsUsernameValid(payload.User.Username) {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 0
		tpr.ErrorText = "invalid username: username length must be greater than 2"
		return tpr
	}

	address := crypto.GetStateAddress([]byte(constant.UserNamespace), []byte(payload.User.Username))

	if payload.Action == pb.Action_UPSERT {
		if !utility.IsPasswordValid(payload.User.Password) {
			tpr.Status = pb.Status_REJECTED
			tpr.ErrorCode = 0
			tpr.ErrorText = "invalid password: password length must be greater than 5"
			return tpr
		}

		return u.upsert(tpr, stateContext, transaction, address, payload.User)
	} else if payload.Action == pb.Action_DELETE {
		return u.delete(tpr, stateContext, address)
	} else {
		tpr.Status = pb.Status_REJECTED
		tpr.ErrorCode = 0
		tpr.ErrorText = "unknown action"
		return tpr
	}
}
