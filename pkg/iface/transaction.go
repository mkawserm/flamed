package iface

type ITransaction interface {
	GetNamespace() []byte
	GetFamily() string
	GetVersion() string
	GetPayload() []byte
}
