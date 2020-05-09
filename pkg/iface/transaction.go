package iface

type ITransaction interface {
	Count() int
	Clear() error
	Commit() error
	Destroy() error
	Delete([]byte) error
	Put([]byte, []byte) error
}
