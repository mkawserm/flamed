package iface

type IStateContext interface {
	GetState(address []byte) ([]byte, error)
	SetState(address []byte, payload []byte) error
	DeleteState(address []byte) error

	SetIndex(id string, data interface{}) error
	DeleteIndex(id string) error
}
