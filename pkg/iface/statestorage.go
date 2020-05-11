package iface

import "io"

type IStateSnapshot interface {
	GetUid() []byte
	GetData() []byte
}

type IStateStorageIterator interface {
	Seek(key []byte)
	Valid() bool
	ValidForPrefix(prefix []byte) bool
	Next()
	Close()
	StateSnapshot() IStateSnapshot
}

type IStateStorageTransaction interface {
	Get(key []byte) ([]byte, error)
	Set(key []byte, value []byte) error
	Delete(key []byte) error
	Discard()
	Commit() error

	ForwardIterator() IStateStorageIterator
	ReverseIterator() IStateStorageIterator
	KeyOnlyForwardIterator() IStateStorageIterator
	KeyOnlyReverseIterator() IStateStorageIterator
}

type IStateStorage interface {
	Open(path string, secretKey []byte, configuration interface{}) error
	Close() error

	RunGC()

	NewTransaction() IStateStorageTransaction

	PrepareSnapshot() (interface{}, error)
	SaveSnapshot(snapshotContext interface{}, w io.Writer) error
	RecoverFromSnapshot(r io.Reader) error

	ChangeSecretKey(path string, oldSecretKey []byte, newSecretKey []byte) error
}
