package iface

type IStateContext interface {
	GetState(key []byte) ([]byte, error)
	SetState(key []byte, value []byte) error
	DeleteState(key []byte) error

	SetIndex(id string, data interface{}) error
	DeleteIndex(id string) error
}
