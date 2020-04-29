package conf

import (
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/plugin/storage/index/bleve"
	"github.com/mkawserm/flamed/pkg/plugin/storage/kv/badger"
)

type StoragedConfigurationInput struct {
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

func (s *StoragedConfiguration) IndexObject(namespace []byte, fields []string, value []byte) (iface.IndexObjectType, interface{}) {
	return iface.JSONMapType, map[string]interface{}{}
}
