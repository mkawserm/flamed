package flamed

import (
	"github.com/mkawserm/flamed/pkg/conf"
	"github.com/mkawserm/flamed/pkg/iface"
)

//type Node internalNode.Node
//type Storage internalStorage.Storage
//type Storaged internalStoraged.Storaged

//type INodeConfiguration internalInterface.INodeConfiguration
//type IClusterConfiguration internalInterface.IClusterConfiguration
//type IStoragedConfiguration internalInterface.IStoragedConfiguration

func SimpleNodeConfiguration(nodeId uint64, nodePath string, raftAddress string) iface.INodeConfiguration {
	return &conf.NodeConfiguration{NodeConfigurationInput: conf.NodeConfigurationInput{
		NodeID:                        nodeId,
		NodePath:                      nodePath,
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
	return &conf.StoragedConfiguration{
		StoragedConfigurationInput: conf.StoragedConfigurationInput{
			StoragePath:      path,
			StorageSecretKey: secretKey,
		}}
}

func SimpleClusterConfiguration(clusterId uint64, clusterName string, initialMembers map[uint64]string, join bool) iface.IClusterConfiguration {
	return &conf.ClusterConfiguration{ClusterConfigurationInput: conf.ClusterConfigurationInput{
		ClusterId:      clusterId,
		ClusterName:    clusterName,
		InitialMembers: initialMembers,
		Join:           join,
	}}
}
