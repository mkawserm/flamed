package registry

import (
	"github.com/mkawserm/flamed/pkg/iface"
	"sync"
)

var (
	kvStorageRegistryOnce sync.Once
	kvStorageRegistryIns  *KVStorageRegistry
)

type KVStorageRegistry struct {
	registry map[string]iface.IKVStorage
}

func (k *KVStorageRegistry) AddKVStorage(name string, storage iface.IKVStorage) {
	k.registry[name] = storage
}

func (k *KVStorageRegistry) GetKVStorage(name string) iface.IKVStorage {
	storage := k.registry[name]
	return storage
}

func GetKVStorageRegistry() *KVStorageRegistry {
	return kvStorageRegistryIns
}

func init() {
	kvStorageRegistryOnce.Do(func() {
		kvStorageRegistryIns = &KVStorageRegistry{
			registry: make(map[string]iface.IKVStorage),
		}
	})
}
