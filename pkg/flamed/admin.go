package flamed

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/lni/dragonboat/v3"
	sm "github.com/lni/dragonboat/v3/statemachine"
	"github.com/mkawserm/flamed/pkg/constant"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/storage"
	"github.com/mkawserm/flamed/pkg/uidutil"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"
	"time"
)

type Admin struct {
	mClusterID          uint64
	mDragonboatNodeHost *dragonboat.NodeHost
}

func (a *Admin) IterateUser(seek *pb.FlameUser, limit int, timeout time.Duration) ([]*pb.FlameUser, error) {
	if seek == nil {
		allocationLength := 100
		if limit != 0 {
			allocationLength = limit
		}
		data := make([]*pb.FlameUser, 0, allocationLength)
		itr := &storage.Iterator{
			Seek:   []byte(constant.UserNamespace),
			Prefix: []byte(constant.UserNamespace),
			Limit:  limit,
			Receiver: func(entry *pb.FlameEntry) bool {
				u := &pb.FlameUser{}
				if err := proto.Unmarshal(entry.Value, u); err == nil {
					data = append(data, u)
				}
				return true
			},
		}

		if _, err := a.managedSyncRead(a.mClusterID, itr, timeout); err != nil {
			return nil, x.ErrFailedToIterate
		} else {
			return data, nil
		}
	} else {
		allocationLength := 100
		if limit != 0 {
			allocationLength = limit
		}

		newLimit := limit
		if newLimit != 0 {
			newLimit = newLimit + 1
		}

		data := make([]*pb.FlameUser, 0, allocationLength)
		uid := uidutil.GetUid([]byte(constant.UserNamespace), []byte(seek.Username))
		firstEntry := true
		itr := &storage.Iterator{
			Seek:   uid,
			Prefix: []byte(constant.UserNamespace),
			Limit:  newLimit,
			Receiver: func(entry *pb.FlameEntry) bool {
				if firstEntry {
					firstEntry = false
					return true
				}

				u := &pb.FlameUser{}
				if err := proto.Unmarshal(entry.Value, u); err == nil {
					data = append(data, u)
				}

				return true
			},
		}

		if _, err := a.managedSyncRead(a.mClusterID, itr, timeout); err != nil {
			return nil, x.ErrFailedToIterate
		} else {
			return data, nil
		}
	}
}

func (a *Admin) IterateAccessControl(seek *pb.FlameAccessControl, limit int, timeout time.Duration) ([]*pb.FlameAccessControl, error) {
	if seek == nil {
		allocationLength := 100
		if limit != 0 {
			allocationLength = limit
		}
		data := make([]*pb.FlameAccessControl, 0, allocationLength)
		itr := &storage.Iterator{
			Seek:   []byte(constant.AccessControlNamespace),
			Prefix: []byte(constant.AccessControlNamespace),
			Limit:  limit,
			Receiver: func(entry *pb.FlameEntry) bool {
				u := &pb.FlameAccessControl{}
				if err := proto.Unmarshal(entry.Value, u); err == nil {
					data = append(data, u)
				}
				return true
			},
		}

		if _, err := a.managedSyncRead(a.mClusterID, itr, timeout); err != nil {
			return nil, x.ErrFailedToIterate
		} else {
			return data, nil
		}
	} else {
		allocationLength := 100
		if limit != 0 {
			allocationLength = limit
		}

		newLimit := limit
		if newLimit != 0 {
			newLimit = newLimit + 1
		}

		data := make([]*pb.FlameAccessControl, 0, allocationLength)
		uid := uidutil.GetUid([]byte(constant.AccessControlNamespace),
			uidutil.GetUid([]byte(seek.Username), seek.Namespace))
		firstEntry := true
		itr := &storage.Iterator{
			Seek:   uid,
			Prefix: []byte(constant.AccessControlNamespace),
			Limit:  newLimit,
			Receiver: func(entry *pb.FlameEntry) bool {
				if firstEntry {
					firstEntry = false
					return true
				}

				u := &pb.FlameAccessControl{}
				if err := proto.Unmarshal(entry.Value, u); err == nil {
					data = append(data, u)
				}

				return true
			},
		}

		if _, err := a.managedSyncRead(a.mClusterID, itr, timeout); err != nil {
			return nil, x.ErrFailedToIterate
		} else {
			return data, nil
		}
	}
}

func (a *Admin) GetUser(user *pb.FlameUser, timeout time.Duration) error {
	_, err := a.managedSyncRead(a.mClusterID, user, timeout)
	return err
}

func (a *Admin) ReadUser(user *pb.FlameUser, timeout time.Duration) error {
	_, err := a.managedSyncRead(a.mClusterID, user, timeout)
	return err
}

func (a *Admin) CreateUser(user *pb.FlameUser, timeout time.Duration) error {
	if !utility.IsFlameUserValid(user) {
		return x.ErrInvalidUser
	}

	userData, err := proto.Marshal(user)
	if err != nil {
		return x.ErrDataMarshalError
	}

	pp := &pb.FlameProposal{
		FlameProposalType: pb.FlameProposal_CREATE_USER,
		FlameProposalData: userData,
	}

	r, err := a.managedSyncApplyProposal(a.mClusterID, pp, timeout)

	if err != nil {
		return x.ErrFailedToCreateUser
	}

	if r.Value > 0 {
		return nil
	} else {
		return x.ErrFailedToCreateUser
	}
}

func (a *Admin) UpdateUser(user *pb.FlameUser, timeout time.Duration) error {
	if !utility.IsFlameUserValid(user) {
		return x.ErrInvalidUser
	}

	userData, err := proto.Marshal(user)
	if err != nil {
		return x.ErrDataMarshalError
	}

	pp := &pb.FlameProposal{
		FlameProposalType: pb.FlameProposal_UPDATE_USER,
		FlameProposalData: userData,
	}

	r, err := a.managedSyncApplyProposal(a.mClusterID, pp, timeout)

	if err != nil {
		return x.ErrFailedToCreateUser
	}

	if r.Value > 0 {
		return nil
	} else {
		return x.ErrFailedToCreateUser
	}
}

func (a *Admin) DeleteUser(user *pb.FlameUser, timeout time.Duration) error {
	if !utility.IsUsernameValid(user.Username) {
		return x.ErrInvalidUser
	}

	userData, err := proto.Marshal(user)
	if err != nil {
		return x.ErrDataMarshalError
	}

	pp := &pb.FlameProposal{
		FlameProposalType: pb.FlameProposal_DELETE_USER,
		FlameProposalData: userData,
	}

	r, err := a.managedSyncApplyProposal(a.mClusterID, pp, timeout)

	if err != nil {
		return x.ErrFailedToDeleteUser
	}

	if r.Value > 0 {
		return nil
	} else {
		return x.ErrFailedToDeleteUser
	}
}

func (a *Admin) GetAccessControl(ac *pb.FlameAccessControl, timeout time.Duration) error {
	if !utility.IsUsernameValid(ac.Username) {
		return x.ErrInvalidUser
	}
	if !utility.IsNamespaceValid(ac.Namespace) {
		return x.ErrInvalidNamespace
	}

	_, err := a.managedSyncRead(a.mClusterID, ac, timeout)
	return err
}

func (a *Admin) ReadAccessControl(ac *pb.FlameAccessControl, timeout time.Duration) error {
	return a.GetAccessControl(ac, timeout)
}

func (a *Admin) CreateAccessControl(ac *pb.FlameAccessControl, timeout time.Duration) error {
	if !utility.IsUsernameValid(ac.Username) {
		return x.ErrInvalidUser
	}
	if !utility.IsNamespaceValid(ac.Namespace) {
		return x.ErrInvalidNamespace
	}

	acData, err := proto.Marshal(ac)
	if err != nil {
		return x.ErrDataMarshalError
	}

	pp := &pb.FlameProposal{
		FlameProposalType: pb.FlameProposal_CREATE_ACCESS_CONTROL,
		FlameProposalData: acData,
	}

	r, err := a.managedSyncApplyProposal(a.mClusterID, pp, timeout)

	if err != nil {
		return x.ErrFailedToCreateAccessControl
	}

	if r.Value > 0 {
		return nil
	} else {
		return x.ErrFailedToCreateAccessControl
	}
}

func (a *Admin) UpdateAccessControl(ac *pb.FlameAccessControl, timeout time.Duration) error {
	if !utility.IsUsernameValid(ac.Username) {
		return x.ErrInvalidUser
	}
	if !utility.IsNamespaceValid(ac.Namespace) {
		return x.ErrInvalidNamespace
	}

	acData, err := proto.Marshal(ac)
	if err != nil {
		return x.ErrDataMarshalError
	}

	pp := &pb.FlameProposal{
		FlameProposalType: pb.FlameProposal_UPDATE_ACCESS_CONTROL,
		FlameProposalData: acData,
	}

	r, err := a.managedSyncApplyProposal(a.mClusterID, pp, timeout)

	if err != nil {
		return x.ErrFailedToUpdateAccessControl
	}

	if r.Value > 0 {
		return nil
	} else {
		return x.ErrFailedToUpdateAccessControl
	}
}

func (a *Admin) DeleteAccessControl(ac *pb.FlameAccessControl, timeout time.Duration) error {
	if !utility.IsUsernameValid(ac.Username) {
		return x.ErrInvalidUser
	}
	if !utility.IsNamespaceValid(ac.Namespace) {
		return x.ErrInvalidNamespace
	}

	acData, err := proto.Marshal(ac)
	if err != nil {
		return x.ErrDataMarshalError
	}

	pp := &pb.FlameProposal{
		FlameProposalType: pb.FlameProposal_DELETE_ACCESS_CONTROL,
		FlameProposalData: acData,
	}

	r, err := a.managedSyncApplyProposal(a.mClusterID, pp, timeout)

	if err != nil {
		return x.ErrFailedToDeleteAccessControl
	}

	if r.Value > 0 {
		return nil
	} else {
		return x.ErrFailedToDeleteAccessControl
	}
}

func (a *Admin) CreateIndexMeta(meta *pb.FlameIndexMeta, timeout time.Duration) error {
	if !utility.IsNamespaceValid(meta.Namespace) {
		return x.ErrInvalidNamespace
	}

	metaData, err := proto.Marshal(meta)
	if err != nil {
		return x.ErrDataMarshalError
	}

	pp := &pb.FlameProposal{
		FlameProposalType: pb.FlameProposal_CREATE_INDEX_META,
		FlameProposalData: metaData,
	}

	r, err := a.managedSyncApplyProposal(a.mClusterID, pp, timeout)

	if err != nil {
		return x.ErrFailedToCreateIndexMeta
	}

	if r.Value > 0 {
		return nil
	} else {
		return x.ErrFailedToCreateIndexMeta
	}
}

func (a *Admin) UpdateIndexMeta(meta *pb.FlameIndexMeta, timeout time.Duration) error {
	if !utility.IsNamespaceValid(meta.Namespace) {
		return x.ErrInvalidNamespace
	}

	metaData, err := proto.Marshal(meta)
	if err != nil {
		return x.ErrDataMarshalError
	}

	pp := &pb.FlameProposal{
		FlameProposalType: pb.FlameProposal_UPDATE_INDEX_META,
		FlameProposalData: metaData,
	}

	r, err := a.managedSyncApplyProposal(a.mClusterID, pp, timeout)

	if err != nil {
		return x.ErrFailedToUpdateIndexMeta
	}

	if r.Value > 0 {
		return nil
	} else {
		return x.ErrFailedToUpdateIndexMeta
	}
}

func (a *Admin) DeleteIndexMeta(meta *pb.FlameIndexMeta, timeout time.Duration) error {
	if !utility.IsNamespaceValid(meta.Namespace) {
		return x.ErrInvalidNamespace
	}

	metaData, err := proto.Marshal(meta)
	if err != nil {
		return x.ErrDataMarshalError
	}

	pp := &pb.FlameProposal{
		FlameProposalType: pb.FlameProposal_DELETE_INDEX_META,
		FlameProposalData: metaData,
	}

	r, err := a.managedSyncApplyProposal(a.mClusterID, pp, timeout)

	if err != nil {
		return x.ErrFailedToDeleteIndexMeta
	}

	if r.Value > 0 {
		return nil
	} else {
		return x.ErrFailedToDeleteIndexMeta
	}
}

func (a *Admin) managedSyncApplyProposal(clusterID uint64,
	pp *pb.FlameProposal,
	timeout time.Duration) (sm.Result, error) {
	cmd, err := proto.Marshal(pp)
	if err != nil {
		return sm.Result{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	session := a.mDragonboatNodeHost.GetNoOPSession(clusterID)
	r, err := a.mDragonboatNodeHost.SyncPropose(ctx, session, cmd)
	cancel()

	_ = a.mDragonboatNodeHost.SyncCloseSession(context.Background(), session)

	return r, err
}

func (a *Admin) managedSyncRead(clusterID uint64, query interface{}, timeout time.Duration) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	d, e := a.mDragonboatNodeHost.SyncRead(ctx, clusterID, query)
	cancel()

	return d, e
}
