package iface

type IKVRaftLogStorage interface {
	Open(dir1, dir2 string, secretKey []byte, config interface{}) error
	Name() string
	Close() error
	IterateValue(fk []byte, lk []byte, inc bool, op func(key []byte, data []byte) (bool, error)) error
	GetValue(key []byte, op func([]byte) error) error
	SaveValue(key []byte, value []byte) error
	DeleteValue(key []byte) error
	GetWriteBatch() ITransaction
	CommitWriteBatch(wb ITransaction) error
	CompactEntries(firstKey []byte, lastKey []byte) error
	FullCompaction() error
	RunGC()
}
