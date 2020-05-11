package iface

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
