package conf

import (
	"encoding/json"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/plugin/storage/index/bleve"
	"github.com/mkawserm/flamed/pkg/plugin/storage/statemachine/badger"
)

type StoragedConfigurationInput struct {
	AutoIndexMeta    bool   `json:"autoIndexMeta"`
	CacheSize        int    `json:"cacheSize"`
	BatchSize        int    `json:"batchSize"`
	StoragePath      string `json:"storagePath"`
	StorageSecretKey []byte `json:"storageSecretKey"`

	StoragePluginStateMachine       iface.IStateMachineStorage `json:"-"`
	StoragePluginIndex              iface.IIndexStorage        `json:"-"`
	KVStorageCustomConfiguration    interface{}                `json:"-"`
	IndexStorageCustomConfiguration interface{}                `json:"-"`
}

type StoragedConfiguration struct {
	StoragedConfigurationInput StoragedConfigurationInput
}

func (s *StoragedConfiguration) AutoIndexMeta() bool {
	return s.StoragedConfigurationInput.AutoIndexMeta
}

func (s *StoragedConfiguration) StoragePath() string {
	return s.StoragedConfigurationInput.StoragePath
}

func (s *StoragedConfiguration) StorageSecretKey() []byte {
	return s.StoragedConfigurationInput.StorageSecretKey
}

func (s *StoragedConfiguration) StoragePluginStateMachine() iface.IStateMachineStorage {
	if s.StoragedConfigurationInput.StoragePluginStateMachine == nil {
		return &badger.Badger{}
	} else {
		return s.StoragedConfigurationInput.StoragePluginStateMachine
	}
}

func (s *StoragedConfiguration) StoragePluginIndex() iface.IIndexStorage {
	if s.StoragedConfigurationInput.StoragePluginIndex == nil {
		return &bleve.Bleve{}
	} else {
		return nil
	}
}

func (s *StoragedConfiguration) StateMachineStorageCustomConfiguration() interface{} {
	return s.StoragedConfigurationInput.KVStorageCustomConfiguration
}

func (s *StoragedConfiguration) IndexStorageCustomConfiguration() interface{} {
	return s.StoragedConfigurationInput.IndexStorageCustomConfiguration
}

func (s *StoragedConfiguration) CacheSize() int {
	if s.StoragedConfigurationInput.CacheSize <= 0 {
		s.StoragedConfigurationInput.CacheSize = 100
	}
	return s.StoragedConfigurationInput.CacheSize
}

func (s *StoragedConfiguration) BatchSize() int {
	if s.StoragedConfigurationInput.BatchSize <= 0 {
		s.StoragedConfigurationInput.BatchSize = 100
	}
	return s.StoragedConfigurationInput.BatchSize
}

func (s *StoragedConfiguration) IndexObject(_, value []byte) interface{} {
	data := make(map[string]interface{})
	if err := json.Unmarshal(value, &data); err == nil {
		return data
	} else {
		//internalLogger.Error("IndexObject json unmarshal error",
		//	zap.Error(err),
		//	zap.String("namespace", string(namespace)))
		return nil
	}
}
