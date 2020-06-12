package iface

type IOnDiskClusterConfiguration interface {
	Join() bool
	ClusterID() uint64
	ClusterName() string
	InitialMembers() map[uint64]string
	StateMachine(IStoragedConfiguration) func(uint64, uint64) IStoraged
}
