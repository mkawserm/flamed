package flamed

import "github.com/lni/dragonboat/v3"

type StorageManager struct {
	mClusterID          uint64
	mDragonboatNodeHost *dragonboat.NodeHost
}
