package conf

import "github.com/mkawserm/flamed/pkg/iface"

func SimpleNodeHostConfiguration(nodeId uint64, nodeHostDir string, walDir string, raftAddress string) iface.INodeConfiguration {
	return &NodeHostConfiguration{NodeHostConfigurationInput: NodeHostConfigurationInput{
		NodeID:                        nodeId,
		NodeHostDir:                   nodeHostDir,
		WALDir:                        walDir,
		CheckQuorum:                   true,
		ElectionRTT:                   5,
		HeartbeatRTT:                  1,
		SnapshotEntries:               10,
		CompactionOverhead:            5,
		OrderedConfigChange:           false,
		MaxInMemLogSize:               0,
		DisableAutoCompactions:        false,
		IsObserver:                    false,
		IsWitness:                     false,
		Quiesce:                       false,
		DeploymentID:                  0,
		RTTMillisecond:                200,
		RaftAddress:                   raftAddress,
		ListenAddress:                 "",
		MutualTLS:                     false,
		CAFile:                        "",
		CertFile:                      "",
		KeyFile:                       "",
		MaxSendQueueSize:              0,
		MaxReceiveQueueSize:           0,
		LogDBFactory:                  nil,
		RaftRPCFactory:                nil,
		EnableMetrics:                 false,
		MaxSnapshotSendBytesPerSecond: 0,
		MaxSnapshotRecvBytesPerSecond: 0,
	}}
}

func SimpleStoragedConfiguration(path string, secretKey []byte) iface.IStoragedConfiguration {
	return &StoragedConfiguration{
		StoragedConfigurationInput: StoragedConfigurationInput{
			StoragePath:      path,
			StorageSecretKey: secretKey,
		}}
}

func SimpleClusterConfiguration(clusterId uint64, clusterName string, initialMembers map[uint64]string, join bool) iface.IClusterConfiguration {
	return &ClusterConfiguration{ClusterConfigurationInput: ClusterConfigurationInput{
		ClusterId:      clusterId,
		ClusterName:    clusterName,
		InitialMembers: initialMembers,
		Join:           join,
	}}
}
