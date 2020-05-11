package iface

type IStateSnapshot interface {
	GetUid() []byte
	GetData() []byte
}
