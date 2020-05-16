package iface

import sm "github.com/lni/dragonboat/v3/statemachine"

type IOnDiskClusterConfiguration interface {
	Join() bool
	ClusterID() uint64
	ClusterName() string
	InitialMembers() map[uint64]string
	StateMachine(IStoragedConfiguration) func(uint64, uint64) sm.IOnDiskStateMachine
}
