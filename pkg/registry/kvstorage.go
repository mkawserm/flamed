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
	registry map[string]iface.IStateMachineStorage
}

func (k *KVStorageRegistry) AddKVStorage(name string, storage iface.IStateMachineStorage) {
	k.registry[name] = storage
}

func (k *KVStorageRegistry) GetKVStorage(name string) iface.IStateMachineStorage {
	storage := k.registry[name]
	return storage
}

func GetKVStorageRegistry() *KVStorageRegistry {
	return kvStorageRegistryIns
}

func init() {
	kvStorageRegistryOnce.Do(func() {
		kvStorageRegistryIns = &KVStorageRegistry{
			registry: make(map[string]iface.IStateMachineStorage),
		}
	})
}
