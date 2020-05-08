package conf

import (
	"encoding/json"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/plugin/storage/index/bleve"
	"github.com/mkawserm/flamed/pkg/plugin/storage/kv/badger"
)

type StoragedConfigurationInput struct {
	CacheSize        int    `json:"cacheSize"`
	BatchSize        int    `json:"batchSize"`
	StoragePath      string `json:"storagePath"`
	StorageSecretKey []byte `json:"storageSecretKey"`

	StoragePluginKV                 iface.IKVStorage    `json:"-"`
	StoragePluginIndex              iface.IIndexStorage `json:"-"`
	KVStorageCustomConfiguration    interface{}         `json:"-"`
	IndexStorageCustomConfiguration interface{}         `json:"-"`
}

type StoragedConfiguration struct {
	StoragedConfigurationInput StoragedConfigurationInput
}

func (s *StoragedConfiguration) StoragePath() string {
	return s.StoragedConfigurationInput.StoragePath
}

func (s *StoragedConfiguration) StorageSecretKey() []byte {
	return s.StoragedConfigurationInput.StorageSecretKey
}

func (s *StoragedConfiguration) StoragePluginKV() iface.IKVStorage {
	if s.StoragedConfigurationInput.StoragePluginKV == nil {
		return &badger.Badger{}
	} else {
		return s.StoragedConfigurationInput.StoragePluginKV
	}
}

func (s *StoragedConfiguration) StoragePluginIndex() iface.IIndexStorage {
	if s.StoragedConfigurationInput.StoragePluginIndex == nil {
		return &bleve.Bleve{}
	} else {
		return nil
	}
}

func (s *StoragedConfiguration) KVStorageCustomConfiguration() interface{} {
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
