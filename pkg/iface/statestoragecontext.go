package iface

type IStateStorageContext interface {
	Get(address []byte) ([]byte, error)
	Set(address []byte, payload []byte) error
	Delete(address []byte) error
}
