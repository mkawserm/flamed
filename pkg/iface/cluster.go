package iface

import sm "github.com/lni/dragonboat/v3/statemachine"

type IClusterConfiguration interface {
	ClusterID() uint64
	ClusterName() string
	InitialMembers() map[uint64]string
	Join() bool
	StateMachine(IStoragedConfiguration) func(uint64, uint64) sm.IOnDiskStateMachine
}
