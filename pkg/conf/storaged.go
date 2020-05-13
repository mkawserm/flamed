package conf

import (
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/plugin/storage/index/blevescorch"
	sBadger "github.com/mkawserm/flamed/pkg/plugin/storage/state/badger"
)

type StoragedConfigurationInput struct {
	CacheSize int `json:"cacheSize"`
	BatchSize int `json:"batchSize"`

	IndexEnable   bool `json:"indexEnable"`
	AutoIndexMeta bool `json:"autoIndexMeta"`

	StateStoragePath      string `json:"stateStoragePath"`
	StateStorageSecretKey []byte `json:"stateStorageSecretKey"`

	IndexStoragePath      string `json:"indexStoragePath"`
	IndexStorageSecretKey []byte `json:"indexStorageSecretKey"`

	StoragePluginIndex iface.IIndexStorage `json:"-"`
	StoragePluginState iface.IStateStorage `json:"-"`

	StateStorageCustomConfiguration interface{} `json:"-"`
	IndexStorageCustomConfiguration interface{} `json:"-"`
}

type StoragedConfiguration struct {
	StoragedConfigurationInput StoragedConfigurationInput
	TransactionProcessorMap    map[string]iface.ITransactionProcessor
}

func (s *StoragedConfiguration) AutoIndexMeta() bool {
	return s.StoragedConfigurationInput.AutoIndexMeta
}

func (s *StoragedConfiguration) IndexEnable() bool {
	return s.StoragedConfigurationInput.IndexEnable
}

func (s *StoragedConfiguration) StateStoragePath() string {
	return s.StoragedConfigurationInput.StateStoragePath
}

func (s *StoragedConfiguration) StateStorageSecretKey() []byte {
	return s.StoragedConfigurationInput.StateStorageSecretKey
}

func (s *StoragedConfiguration) IndexStoragePath() string {
	return s.StoragedConfigurationInput.IndexStoragePath
}

func (s *StoragedConfiguration) IndexStorageSecretKey() []byte {
	return s.StoragedConfigurationInput.IndexStorageSecretKey
}

func (s *StoragedConfiguration) StoragePluginState() iface.IStateStorage {
	if s.StoragedConfigurationInput.StoragePluginState == nil {
		return &sBadger.Badger{}
	} else {
		return s.StoragedConfigurationInput.StoragePluginState
	}
}

func (s *StoragedConfiguration) StoragePluginIndex() iface.IIndexStorage {
	if s.StoragedConfigurationInput.StoragePluginIndex == nil {
		return &blevescorch.BleveScorch{}
	} else {
		return nil
	}
}

func (s *StoragedConfiguration) StateStorageCustomConfiguration() interface{} {
	return s.StoragedConfigurationInput.StateStorageCustomConfiguration
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

func (s *StoragedConfiguration) GetTransactionProcessor(family, version string) iface.ITransactionProcessor {
	val, found := s.TransactionProcessorMap[family+"::"+version]

	if found {
		return val
	}

	return nil
}

func (s *StoragedConfiguration) AddTransactionProcessor(tp iface.ITransactionProcessor) {
	s.TransactionProcessorMap[tp.FamilyName()+"::"+tp.FamilyVersion()] = tp
}

func (s *StoragedConfiguration) IsTransactionProcessorExists(family, version string) bool {
	_, found := s.TransactionProcessorMap[family+"::"+version]
	return found
}
