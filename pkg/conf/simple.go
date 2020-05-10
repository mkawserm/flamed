package conf

import (
	"github.com/lni/dragonboat/v3/plugin/pebble"
	"github.com/mkawserm/flamed/pkg/iface"
)

func SimpleNodeHostConfiguration(nodeID uint64, nodeHostDir string, walDir string, raftAddress string) iface.INodeConfiguration {
	return &NodeHostConfiguration{NodeHostConfigurationInput: NodeHostConfigurationInput{
		NodeID:                        nodeID,
		NodeHostDir:                   nodeHostDir,
		WALDir:                        walDir,
		CheckQuorum:                   true,
		ElectionRTT:                   5,
		HeartbeatRTT:                  1,
		SnapshotEntries:               100,
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
		LogDBFactory:                  pebble.NewLogDB,
		RaftRPCFactory:                nil,
		EnableMetrics:                 false,
		MaxSnapshotSendBytesPerSecond: 0,
		MaxSnapshotRecvBytesPerSecond: 0,
	}}
}

func SimpleStoragedConfiguration(path string, secretKey []byte) iface.IStoragedConfiguration {
	return &StoragedConfiguration{
		StoragedConfigurationInput: StoragedConfigurationInput{
			AutoIndexMeta:    true,
			StoragePath:      path,
			StorageSecretKey: secretKey,
		}}
}

func SimpleClusterConfiguration(clusterID uint64, clusterName string, initialMembers map[uint64]string, join bool) iface.IClusterConfiguration {
	return &ClusterConfiguration{ClusterConfigurationInput: ClusterConfigurationInput{
		ClusterID:      clusterID,
		ClusterName:    clusterName,
		InitialMembers: initialMembers,
		Join:           join,
	}}
}
