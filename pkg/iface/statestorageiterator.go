package iface

type IStateStorageIterator interface {
	Seek(key []byte)
	Valid() bool
	ValidForPrefix(prefix []byte) bool
	Next()
	Close()
	StateSnapshot() IStateSnapshot
}
