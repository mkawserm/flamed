package storage

import (
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
)

type Storage struct {
	mConfiguration iface.IStorageConfiguration

	mSecretKey                      []byte
	mKVStoragePath                  string
	mKVStorage                      iface.IKVStorage
	mKVStorageConfiguration         interface{}
	mKVStorageSnapshotConfiguration interface{}

	mIndexStoragePath          string
	mIndexStorage              iface.IIndexStorage
	mIndexStorageConfiguration interface{}
}

func (s *Storage) SetConfiguration(configuration iface.IStorageConfiguration) bool {
	s.mConfiguration = configuration

	if s.mConfiguration.StoragePluginKV() == nil {
		return false
	}

	if s.mConfiguration.StoragePluginRaftLog() == nil {
		return false
	}

	kvStoragePath := s.mConfiguration.StoragePath() + "/kv"
	indexStoragePath := s.mConfiguration.StoragePath() + "/index"

	if !utility.MkPath(kvStoragePath) {
		return false
	}
	if !utility.MkPath(indexStoragePath) {
		return false
	}

	s.mSecretKey = s.mConfiguration.StorageSecretKey()
	s.mKVStorage = s.mConfiguration.StoragePluginKV()
	s.mKVStoragePath = kvStoragePath
	s.mKVStorageConfiguration = s.mConfiguration.KVStorageCustomConfiguration()

	s.mIndexStorage = s.mConfiguration.StoragePluginIndex()
	s.mIndexStoragePath = indexStoragePath
	s.mIndexStorageConfiguration = s.mConfiguration.IndexStorageCustomConfiguration()

	return true
}

func (s *Storage) Open() (bool, error) {
	return s.mKVStorage.Open(s.mKVStoragePath, s.mSecretKey, false, s.mKVStorageConfiguration)
}

func (s *Storage) Close() error {
	return s.mKVStorage.Close()
}

func (s *Storage) RunGC() {
	s.mKVStorage.RunGC()
}

func (s *Storage) ChangeSecretKey(oldSecretKey []byte, newSecretKey []byte) (bool, error) {
	return s.mKVStorage.ChangeSecretKey(oldSecretKey, newSecretKey)
}

func (s *Storage) IsExists(namespace []byte, key []byte) bool {
	return s.mKVStorage.IsExists(namespace, key)
}

func (s *Storage) Read(namespace []byte, key []byte) ([]byte, error) {
	return s.mKVStorage.Read(namespace, key)
}

func (s *Storage) Delete(namespace []byte, key []byte) (bool, error) {
	return s.mKVStorage.Delete(namespace, key)
}

func (s *Storage) Create(namespace []byte, key []byte, value []byte) (bool, error) {
	return s.mKVStorage.Create(namespace, key, value)
}

func (s *Storage) Update(namespace []byte, key []byte, value []byte) (bool, error) {
	return s.mKVStorage.Update(namespace, key, value)
}

func (s *Storage) ApplyBatch(batch *pb.FlameBatch) (bool, error) {
	return s.mKVStorage.ApplyBatch(batch)
}

func (s *Storage) ApplyAction(action *pb.FlameAction) (bool, error) {
	return s.mKVStorage.ApplyAction(action)
}

func (s *Storage) AsyncSnapshot(snapshot chan *pb.FlameSnapshot) error {
	return s.mKVStorage.AsyncSnapshot(snapshot)
}

func (s *Storage) ApplyAsyncSnapshot(snapshot chan *pb.FlameSnapshot) (bool, error) {
	return s.mKVStorage.ApplyAsyncSnapshot(snapshot)
}

func (s *Storage) SyncSnapshot() (*pb.FlameSnapshot, error) {
	return s.mKVStorage.SyncSnapshot()
}

func (s *Storage) ApplySyncSnapshot(snapshot *pb.FlameSnapshot) (bool, error) {
	return s.mKVStorage.ApplySyncSnapshot(snapshot)
}
