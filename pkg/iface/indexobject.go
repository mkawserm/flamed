package iface

type IndexObjectType int

const (
	JSONMapType IndexObjectType = iota
	BleveClassifierType
	GolangStructType
)

type IndexObject interface {
	IndexObject(namespace, value []byte) (IndexObjectType, interface{})
}
