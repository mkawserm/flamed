package iface

import "github.com/mkawserm/flamed/pkg/pb"

type IStateStorageIterator interface {
	Next()
	Close()
	Valid() bool
	Seek(key []byte)
	StateSnapshot() *pb.StateSnapshot
	ValidForPrefix(prefix []byte) bool
}

type IStateStorageTransaction interface {
	Discard()
	Commit() error
	Delete(key []byte) error
	Get(key []byte) ([]byte, error)
	Set(key []byte, value []byte) error

	ForwardIterator() IStateStorageIterator
	ReverseIterator() IStateStorageIterator
	KeyOnlyForwardIterator() IStateStorageIterator
	KeyOnlyReverseIterator() IStateStorageIterator
}

type IStateStorage interface {
	RunGC()
	Close() error
	NewTransaction() IStateStorageTransaction
	NewReadOnlyTransaction() IStateStorageTransaction
	Open(path string, secretKey []byte, configuration interface{}) error
	ChangeSecretKey(path string, oldSecretKey []byte, newSecretKey []byte) error
}
