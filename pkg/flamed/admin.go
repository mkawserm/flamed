package flamed

import (
	"bytes"
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/constant"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/tp/accesscontrol"
	"github.com/mkawserm/flamed/pkg/tp/indexmeta"
	"github.com/mkawserm/flamed/pkg/tp/user"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/variant"
	"github.com/mkawserm/flamed/pkg/x"
	"time"
)

type Admin struct {
	mClusterID uint64
	mRW        iface.IRW
	mTimeout   time.Duration
}

func (a *Admin) UpdateTimeout(timeout time.Duration) {
	a.mTimeout = timeout
}

func (a *Admin) IsUserExists(username string) bool {
	u, err := a.GetUser(username)
	if err != nil {
		return false
	}

	if u.Username == username {
		return true
	}

	return false
}

func (a *Admin) GetUser(username string) (*pb.User, error) {
	if !utility.IsUsernameValid(username) {
		return nil, x.ErrInvalidUsername
	}

	lookupRequest := variant.LookupRequest{
		Query:         username,
		Context:       context.TODO(),
		FamilyName:    user.Name,
		FamilyVersion: user.Version,
	}

	output, err := a.mRW.Read(a.mClusterID, lookupRequest, a.mTimeout)

	if err != nil {
		return nil, err
	}

	if v, ok := output.(*pb.User); ok {
		return v, nil
	} else {
		return nil, x.ErrUnknownValue
	}
}

func (a *Admin) UpsertUser(u *pb.User) (*pb.ProposalResponse, error) {
	if !utility.IsUsernameValid(u.Username) {
		return nil, x.ErrInvalidUsername
	}

	if !utility.IsPasswordValid(u.Password) {
		return nil, x.ErrInvalidPassword
	}

	payload := &pb.UserPayload{
		Action: pb.Action_UPSERT,
		User:   u,
	}

	payloadBytes, err := proto.Marshal(payload)

	if err != nil {
		return nil, err
	}

	proposal := pb.NewProposal()
	proposal.AddTransaction([]byte(constant.UserNamespace), user.Name, user.Version, payloadBytes)

	r, err := a.mRW.Write(a.mClusterID, proposal, a.mTimeout)

	if err != nil {
		return nil, err
	}

	pr := &pb.ProposalResponse{}

	if err := proto.Unmarshal(r.Data, pr); err != nil {
		return nil, err
	}

	return pr, nil
}

func (a *Admin) DeleteUser(username string) (*pb.ProposalResponse, error) {
	if !utility.IsUsernameValid(username) {
		return nil, x.ErrInvalidUsername
	}

	payload := &pb.UserPayload{
		Action: pb.Action_DELETE,
		User:   &pb.User{Username: username},
	}

	payloadBytes, err := proto.Marshal(payload)

	if err != nil {
		return nil, err
	}

	proposal := pb.NewProposal()
	proposal.AddTransaction([]byte(constant.UserNamespace), user.Name, user.Version, payloadBytes)

	r, err := a.mRW.Write(a.mClusterID, proposal, a.mTimeout)

	if err != nil {
		return nil, err
	}

	pr := &pb.ProposalResponse{}

	if err := proto.Unmarshal(r.Data, pr); err != nil {
		return nil, err
	}

	return pr, nil
}

func (a *Admin) IsAccessControlExists(username string, namespace []byte) bool {
	ac, err := a.GetAccessControl(username, namespace)
	if err != nil {
		return false
	}

	if bytes.Equal(ac.Namespace, namespace) && ac.Username == username {
		return true
	}

	return false
}

func (a *Admin) GetAccessControl(username string, namespace []byte) (*pb.AccessControl, error) {
	if !utility.IsNamespaceValid(namespace) {
		return nil, x.ErrInvalidNamespace
	}

	if !utility.IsUsernameValid(username) {
		return nil, x.ErrInvalidUsername
	}

	lookupRequest := variant.LookupRequest{
		Query: accesscontrol.Request{
			Username:  username,
			Namespace: namespace,
		},

		Context:       context.TODO(),
		FamilyName:    accesscontrol.Name,
		FamilyVersion: accesscontrol.Version,
	}

	output, err := a.mRW.Read(a.mClusterID, lookupRequest, a.mTimeout)

	if err != nil {
		return nil, err
	}

	if v, ok := output.(*pb.AccessControl); ok {
		return v, nil
	} else {
		return nil, x.ErrUnknownValue
	}
}

func (a *Admin) UpsertAccessControl(ac *pb.AccessControl) (*pb.ProposalResponse, error) {
	if !utility.IsNamespaceValid(ac.Namespace) {
		return nil, x.ErrInvalidNamespace
	}

	if !utility.IsUsernameValid(ac.Username) {
		return nil, x.ErrInvalidUsername
	}

	payload := &pb.AccessControlPayload{
		Action:        pb.Action_UPSERT,
		AccessControl: ac,
	}

	payloadBytes, err := proto.Marshal(payload)

	if err != nil {
		return nil, err
	}

	proposal := pb.NewProposal()
	proposal.AddTransaction([]byte(constant.AccessControlNamespace),
		accesscontrol.Name,
		accesscontrol.Version,
		payloadBytes)

	r, err := a.mRW.Write(a.mClusterID, proposal, a.mTimeout)

	if err != nil {
		return nil, err
	}

	pr := &pb.ProposalResponse{}

	if err := proto.Unmarshal(r.Data, pr); err != nil {
		return nil, err
	}

	return pr, nil
}

func (a *Admin) DeleteAccessControl(namespace []byte, username string) (*pb.ProposalResponse, error) {
	if !utility.IsNamespaceValid(namespace) {
		return nil, x.ErrInvalidNamespace
	}

	if !utility.IsUsernameValid(username) {
		return nil, x.ErrInvalidUsername
	}

	payload := &pb.AccessControlPayload{
		Action:        pb.Action_DELETE,
		AccessControl: &pb.AccessControl{Username: username, Namespace: namespace},
	}

	payloadBytes, err := proto.Marshal(payload)

	if err != nil {
		return nil, err
	}

	proposal := pb.NewProposal()
	proposal.AddTransaction([]byte(constant.AccessControlNamespace),
		accesscontrol.Name,
		accesscontrol.Version,
		payloadBytes)

	r, err := a.mRW.Write(a.mClusterID, proposal, a.mTimeout)

	if err != nil {
		return nil, err
	}

	pr := &pb.ProposalResponse{}

	if err := proto.Unmarshal(r.Data, pr); err != nil {
		return nil, err
	}

	return pr, nil
}

func (a *Admin) GetIndexMeta(namespace []byte) (*pb.IndexMeta, error) {
	if !utility.IsNamespaceValid(namespace) {
		return nil, x.ErrInvalidNamespace
	}

	lookupRequest := variant.LookupRequest{
		Query:         namespace,
		Context:       context.TODO(),
		FamilyName:    indexmeta.Name,
		FamilyVersion: indexmeta.Version,
	}

	output, err := a.mRW.Read(a.mClusterID, lookupRequest, a.mTimeout)

	if err != nil {
		return nil, err
	}

	if v, ok := output.(*pb.IndexMeta); ok {
		return v, nil
	} else {
		return nil, x.ErrUnknownValue
	}
}

func (a *Admin) UpsertIndexMeta(meta *pb.IndexMeta) (*pb.ProposalResponse, error) {
	if !utility.IsNamespaceValid(meta.Namespace) {
		return nil, x.ErrInvalidNamespace
	}

	payload := &pb.IndexMetaPayload{
		Action:    pb.Action_UPSERT,
		IndexMeta: meta,
	}

	payloadBytes, err := proto.Marshal(payload)

	if err != nil {
		return nil, err
	}

	proposal := pb.NewProposal()
	proposal.AddTransaction([]byte(constant.IndexMetaNamespace), indexmeta.Name, indexmeta.Version, payloadBytes)

	r, err := a.mRW.Write(a.mClusterID, proposal, a.mTimeout)

	if err != nil {
		return nil, err
	}

	pr := &pb.ProposalResponse{}

	if err := proto.Unmarshal(r.Data, pr); err != nil {
		return nil, err
	}

	return pr, nil
}

func (a *Admin) DeleteIndexMeta(namespace []byte) (*pb.ProposalResponse, error) {
	if !utility.IsNamespaceValid(namespace) {
		return nil, x.ErrInvalidNamespace
	}

	payload := &pb.IndexMetaPayload{
		Action:    pb.Action_DELETE,
		IndexMeta: &pb.IndexMeta{Namespace: namespace},
	}

	payloadBytes, err := proto.Marshal(payload)

	if err != nil {
		return nil, err
	}

	proposal := pb.NewProposal()
	proposal.AddTransaction([]byte(constant.IndexMetaNamespace), indexmeta.Name, indexmeta.Version, payloadBytes)

	r, err := a.mRW.Write(a.mClusterID, proposal, a.mTimeout)

	if err != nil {
		return nil, err
	}

	pr := &pb.ProposalResponse{}

	if err := proto.Unmarshal(r.Data, pr); err != nil {
		return nil, err
	}

	return pr, nil
}
