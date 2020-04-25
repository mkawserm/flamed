package badger

import (
	"github.com/mkawserm/flamed/pkg/pb"
)

type Badger struct {
}

func (b *Badger) Open(path string, secretKey []byte, configuration interface{}) (bool, error) {

	return false, nil
}

func (b *Badger) Close() error {
	return nil
}

func (b *Badger) RunGC() {

}

func (b *Badger) ChangeSecretKey(oldSecretKey []byte, newSecretKey []byte) (bool, error) {
	return false, nil
}

func (b *Badger) Read(namespace []byte, key []byte) ([]byte, error) {
	return nil, nil
}

func (b *Badger) Delete(namespace []byte, key []byte) (bool, error) {
	return false, nil
}

func (b *Badger) Create(namespace []byte, key []byte, value []byte) (bool, error) {
	return false, nil
}

func (b *Badger) Update(namespace []byte, key []byte, value []byte) (bool, error) {
	return false, nil
}

func (b *Badger) ApplyBatch(batch *pb.FlameBatch) (bool, error) {
	return false, nil
}

func (b *Badger) ApplyAction(batch *pb.FlameAction) (bool, error) {
	return false, nil
}

func (b *Badger) SetSnapshotConfiguration(configuration interface{}) {

}

func (b *Badger) AsyncSnapshot(snapshot chan *pb.FlameSnapshot, maxItem int) error {
	return nil
}

func (b *Badger) ApplyAsyncSnapshot(snapshot chan *pb.FlameSnapshot) error {
	return nil
}

func (b *Badger) SyncSnapshot() (*pb.FlameSnapshot, error) {
	return nil, nil
}

func (b *Badger) ApplySyncSnapshot(snapshot *pb.FlameSnapshot) (bool, error) {
	return false, nil
}
