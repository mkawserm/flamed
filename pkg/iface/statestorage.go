package iface

import "io"

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
