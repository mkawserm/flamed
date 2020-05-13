package iface

import (
	"github.com/mkawserm/flamed/pkg/pb"
	"time"
)

type IStateStorageIterator interface {
	Next()
	Close()
	Valid() bool
	Rewind()
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
	Open() error
	Close() error
	NewTransaction() IStateStorageTransaction
	NewReadOnlyTransaction() IStateStorageTransaction
	Setup(path string, secretKey []byte, configuration interface{})
	ChangeSecretKey(path string, oldSecretKey []byte, newSecretKey []byte,
		encryptionKeyRotationDuration time.Duration) error
}
