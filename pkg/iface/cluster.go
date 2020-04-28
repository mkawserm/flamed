package iface

type IClusterConfiguration interface {
	ClusterID() uint64
	ClusterName() string
	InitialMembers() map[uint64]string
	Join() bool
}

type ICluster interface {
}
