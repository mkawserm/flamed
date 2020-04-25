package storage

import "github.com/mkawserm/flamed/pkg/iface"

type Storage struct {
	mSecretKey []byte

	mKVStoragePath          string
	mKVStorage              iface.IKVStorage
	mKVStorageConfiguration interface{}

	mIndexStoragePath          string
	mIndexStorage              iface.IIndexStorage
	mIndexStorageConfiguration interface{}
}

func (s *Storage) SetSecretKey(secretKey []byte) {
	s.mSecretKey = secretKey
}

func (s *Storage) SetKVStorage(kvStorage iface.IKVStorage) {
	s.mKVStorage = kvStorage
}

func (s *Storage) SetKVStoragePath(path string) {
	s.mKVStoragePath = path
}

func (s *Storage) SetKVStorageConfiguration(configuration interface{}) {
	s.mKVStorageConfiguration = configuration
}

func (s *Storage) SetIndexStorage(indexStorage iface.IIndexStorage) {
	s.mIndexStorage = indexStorage
}

func (s *Storage) SetIndexStoragePath(path string) {
	s.mIndexStoragePath = path
}

func (s *Storage) SetIndexStorageConfiguration(configuration interface{}) {
	s.mIndexStorageConfiguration = configuration
}
