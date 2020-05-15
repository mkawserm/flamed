package conf

import (
	"github.com/lni/dragonboat/v3/config"
	"github.com/mkawserm/flamed/pkg/iface"
)

func SimpleRaftConfiguration(nodeID uint64) iface.IRaftConfiguration {
	return &RaftConfiguration{RaftConfigurationInput: RaftConfigurationInput{
		NodeID:                 nodeID,
		CheckQuorum:            true,
		ElectionRTT:            5,
		HeartbeatRTT:           1,
		SnapshotEntries:        100,
		CompactionOverhead:     5,
		OrderedConfigChange:    false,
		MaxInMemLogSize:        0,
		DisableAutoCompactions: false,
		IsObserver:             false,
		IsWitness:              false,
		Quiesce:                false,
	}}
}

func SimpleNodeConfiguration(nodeHostDir string, walDir string, raftAddress string) iface.INodeConfiguration {
	return &NodeConfiguration{NodeConfigurationInput: NodeConfigurationInput{
		NodeHostDir:                   nodeHostDir,
		WALDir:                        walDir,
		DeploymentID:                  1,
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

		LogDBConfig:  config.GetTinyMemLogDBConfig(),
		NotifyCommit: false,
	}}
}

func SimpleStoragedConfiguration(path string, secretKey []byte) iface.IStoragedConfiguration {
	return &StoragedConfiguration{
		StoragedConfigurationInput: StoragedConfigurationInput{
			AutoIndexMeta:         true,
			IndexEnable:           true,
			StateStoragePath:      path + "/state",
			StateStorageSecretKey: secretKey,
			IndexStoragePath:      path + "/index",
			IndexStorageSecretKey: secretKey,

			AutoBuildIndex: true,
		},
		TransactionProcessorMap: make(map[string]iface.ITransactionProcessor)}
}

func SimpleOnDiskClusterConfiguration(clusterID uint64, clusterName string, initialMembers map[uint64]string, join bool) iface.IOnDiskClusterConfiguration {
	return &OnDiskClusterConfiguration{OnDiskClusterConfigurationInput: OnDiskClusterConfigurationInput{
		ClusterID:      clusterID,
		ClusterName:    clusterName,
		InitialMembers: initialMembers,
		Join:           join,
	}}
}
