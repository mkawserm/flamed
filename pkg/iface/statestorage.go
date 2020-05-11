package iface

type IStateSnapshot interface {
	GetUid() []byte
	GetData() []byte
}

type IStateStorageIterator interface {
	Next()
	Close()
	Valid() bool
	Seek(key []byte)
	StateSnapshot() IStateSnapshot
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
