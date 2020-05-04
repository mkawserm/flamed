package iface

//type IndexObjectType int
//
//const (
//	JSONMapType IndexObjectType = iota
//	BleveClassifierType
//	GolangStructType
//)

type IIndexObject interface {
	IndexObject(namespace, value []byte) interface{}
}
