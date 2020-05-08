package flamed

import "github.com/lni/dragonboat/v3"

type EntryManager struct {
	mClusterID          uint64
	mDragonboatNodeHost *dragonboat.NodeHost
}
