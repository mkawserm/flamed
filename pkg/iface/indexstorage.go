package iface

type IIndexStorage interface {
	Index(id string, data interface{}) error
}
