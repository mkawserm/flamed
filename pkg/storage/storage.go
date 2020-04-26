package storage

import (
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
)

type Storage struct {
	mConfiguration iface.IFlameConfiguration

	mSecretKey                      []byte
	mKVStoragePath                  string
	mKVStorage                      iface.IKVStorage
	mKVStorageConfiguration         interface{}
	mKVStorageSnapshotConfiguration interface{}

	mIndexStoragePath          string
	mIndexStorage              iface.IIndexStorage
	mIndexStorageConfiguration interface{}
}

func (s *Storage) SetConfiguration(configuration iface.IFlameConfiguration) {
	s.mConfiguration = configuration

	s.mSecretKey = s.mConfiguration.FlameSecretKey()
	s.mKVStorage = s.mConfiguration.StoragePluginKV()
	s.mKVStoragePath = s.mConfiguration.FlamePath()
	s.mKVStorageConfiguration = s.mConfiguration.KVStorageCustomConfiguration()
	s.mKVStorageSnapshotConfiguration = s.mConfiguration.KVStorageSnapshotConfiguration()

	s.mIndexStorage = s.mConfiguration.StoragePluginIndex()
	s.mIndexStoragePath = s.mConfiguration.FlamePath()
	s.mIndexStorageConfiguration = s.mConfiguration.IndexStorageCustomConfiguration()
}

func (s *Storage) Open() (bool, error) {
	return s.mKVStorage.Open(s.mKVStoragePath, s.mSecretKey, s.mKVStorageConfiguration)
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

func (s *Storage) ApplyAction(batch *pb.FlameAction) (bool, error) {
	return s.mKVStorage.ApplyAction(batch)
}

func (s *Storage) AsyncSnapshot(snapshot chan *pb.FlameSnapshot, maxItem int) error {
	s.mKVStorage.SetSnapshotConfiguration(s.mKVStorageSnapshotConfiguration)
	return s.mKVStorage.AsyncSnapshot(snapshot, maxItem)
}

func (s *Storage) ApplyAsyncSnapshot(snapshot chan *pb.FlameSnapshot) error {
	return s.mKVStorage.ApplyAsyncSnapshot(snapshot)
}

func (s *Storage) SyncSnapshot() (*pb.FlameSnapshot, error) {
	s.mKVStorage.SetSnapshotConfiguration(s.mKVStorageSnapshotConfiguration)
	return s.mKVStorage.SyncSnapshot()
}

func (s *Storage) ApplySyncSnapshot(snapshot *pb.FlameSnapshot) (bool, error) {
	return s.mKVStorage.ApplySyncSnapshot(snapshot)
}
