package iface

import sm "github.com/lni/dragonboat/v3/statemachine"

type NewStateMachineFunc func(IStoragedConfiguration) func(uint64, uint64) sm.IOnDiskStateMachine

type IClusterConfiguration interface {
	ClusterId() uint64
	ClusterName() string
	InitialMembers() map[uint64]string
	Join() bool
	StateMachine() NewStateMachineFunc
}
