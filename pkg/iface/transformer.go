package iface

type ITransformer interface {
	ToObject(namespace []byte, data []byte) (interface{}, error)
	ToByteSlice(namespace []byte, object interface{}) ([]byte, error)
}
