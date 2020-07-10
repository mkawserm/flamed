package iface

type IOnDiskClusterConfiguration interface {
	// Join defines the join flag for on disk cluster
	// configuration
	Join() bool

	// ClusterID method contains cluster id
	ClusterID() uint64

	// ClusterName method contains cluster name
	ClusterName() string

	// InitialMembers method defines all initial members of the cluster
	InitialMembers() map[uint64]string

	// StateMachine method return a method which returns storaged state machine
	StateMachine(IStoragedConfiguration) func(uint64, uint64) IStoraged
}
