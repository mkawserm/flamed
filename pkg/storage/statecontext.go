package storage

import (
	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
)

type StateContext struct {
	mReadOnly            bool
	mStorage             *Storage
	mIndexDataList       []*variant.IndexData
	mIndexMetaActionList []*variant.IndexMetaAction
	mTxn                 iface.IStateStorageTransaction
}

func (s *StateContext) GetForwardIterator() iface.IStateIterator {
	return s.mTxn.ForwardIterator()
}

func (s *StateContext) GetReverseIterator() iface.IStateIterator {
	return s.mTxn.ReverseIterator()
}

func (s *StateContext) GetKeyOnlyForwardIterator() iface.IStateIterator {
	return s.mTxn.KeyOnlyForwardIterator()
}

func (s *StateContext) GetKeyOnlyReverseIterator() iface.IStateIterator {
	return s.mTxn.KeyOnlyReverseIterator()
}

func (s *StateContext) GetState(key []byte) (*pb.StateEntry, error) {
	if data, err := s.mTxn.Get(key); err == nil {
		entry := &pb.StateEntry{}
		if err := proto.Unmarshal(data, entry); err == nil {
			return entry, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (s *StateContext) UpsertState(key []byte, entry *pb.StateEntry) error {
	if s.mReadOnly {
		return nil
	}

	if data, err := proto.Marshal(entry); err == nil {
		return s.mTxn.Set(key, data)
	} else {
		return err
	}
}

func (s *StateContext) DeleteState(key []byte) error {
	if s.mReadOnly {
		return nil
	}

	return s.mTxn.Delete(key)
}

func (s *StateContext) UpsertIndex(id string, data interface{}) error {
	if s.mReadOnly {
		return nil
	}

	if !s.mStorage.IndexEnable() {
		return nil
	}

	s.mIndexDataList = append(s.mIndexDataList, &variant.IndexData{
		ID:     id,
		Data:   data,
		Action: pb.Action_UPSERT,
	})
	return nil
}

func (s *StateContext) DeleteIndex(id string) error {
	if s.mReadOnly {
		return nil
	}

	if !s.mStorage.IndexEnable() {
		return nil
	}

	s.mIndexDataList = append(s.mIndexDataList, &variant.IndexData{
		ID:     id,
		Action: pb.Action_DELETE,
	})
	return nil
}

//func (s *StateContext) AutoIndexMeta() bool {
//	return s.mStorage.AutoIndexMeta()
//}

//func (s *StateContext) CanIndex(namespace string) bool {
//	if !s.mStorage.IndexEnable() {
//		return false
//	}
//
//	return s.mStorage.CanIndex(namespace)
//}

func (s *StateContext) UpsertIndexMeta(meta *pb.IndexMeta) error {
	if s.mReadOnly {
		return nil
	}

	if !s.mStorage.IndexEnable() {
		return nil
	}

	m := &variant.IndexMetaAction{
		Action:    pb.Action_UPSERT,
		IndexMeta: meta,
	}
	s.mIndexMetaActionList = append(s.mIndexMetaActionList, m)
	return nil
}

func (s *StateContext) DeleteIndexMeta(meta *pb.IndexMeta) error {
	if s.mReadOnly {
		return nil
	}

	if !s.mStorage.IndexEnable() {
		return nil
	}

	m := &variant.IndexMetaAction{
		Action:    pb.Action_DELETE,
		IndexMeta: meta,
	}
	s.mIndexMetaActionList = append(s.mIndexMetaActionList, m)
	return nil
}

func (s *StateContext) DefaultIndexMeta(namespace string) error {
	if s.mReadOnly {
		return nil
	}

	if !s.mStorage.IndexEnable() {
		return nil
	}

	m := &variant.IndexMetaAction{
		Action:    pb.Action_DEFAULT,
		IndexMeta: &pb.IndexMeta{Namespace: []byte(namespace)},
	}
	s.mIndexMetaActionList = append(s.mIndexMetaActionList, m)
	return nil
}

//func (s *StateContext) ApplyIndex(namespace string, data []*kind.IndexData) error {
//	if s.mReadOnly {
//		return nil
//	}
//
//	return s.mStorage.ApplyIndex(namespace, data)
//}
