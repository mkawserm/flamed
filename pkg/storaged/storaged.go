package storaged

type Storaged struct {
}

//Open(stopc <-chan struct{}) (uint64, error)
//Update([]Entry) ([]Entry, error)
//Lookup(interface{}) (interface{}, error)
//Sync() error
//PrepareSnapshot() (interface{}, error)
//SaveSnapshot(interface{}, io.Writer, <-chan struct{}) error
//RecoverFromSnapshot(io.Reader, <-chan struct{}) error
//Close() error

// NewDiskKV creates a new disk kv test state machine.
//func NewDiskKV(clusterID uint64, nodeID uint64) sm.IOnDiskStateMachine {
//	d := &DiskKV{
//		clusterID: clusterID,
//		nodeID:    nodeID,
//	}
//	return d
//}

//func init() {
//	registry.GetKVStorageRegistry().AddKVStorage(badger.Name, &badger.Badger{})
//}
