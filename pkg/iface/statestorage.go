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
	Seek(address []byte)
	StateSnapshot() *pb.StateSnapshot
	ValidForPrefix(prefix []byte) bool
}

type IStateStorageTransaction interface {
	Discard()
	Commit() error
	Delete(address []byte) error
	Get(address []byte) ([]byte, error)
	Set(address []byte, data []byte) error

	ForwardIterator() IStateStorageIterator
	ReverseIterator() IStateStorageIterator
	KeyOnlyForwardIterator() IStateStorageIterator
	KeyOnlyReverseIterator() IStateStorageIterator
}

type IStateStorage interface {
	RunGC()
	Open() error
	Close() error
	StateStorageName() string
	NewTransaction() IStateStorageTransaction
	NewReadOnlyTransaction() IStateStorageTransaction
	Setup(path string, secretKey []byte, configuration interface{})
	ChangeSecretKey(path string, oldSecretKey []byte, newSecretKey []byte,
		encryptionKeyRotationDuration time.Duration) error
}
