package flamed

import internalNode "github.com/mkawserm/flamed/pkg/node"
import internalStorage "github.com/mkawserm/flamed/pkg/storage"
import internalStoraged "github.com/mkawserm/flamed/pkg/storaged"
import internalInterface "github.com/mkawserm/flamed/pkg/iface"
import internalConf "github.com/mkawserm/flamed/pkg/conf"

type Node internalNode.Node
type Storage internalStorage.Storage
type Storaged internalStoraged.Storaged

type INodeConfiguration internalInterface.INodeConfiguration
type IClusterConfiguration internalInterface.IClusterConfiguration
type IStoragedConfiguration internalInterface.IStoragedConfiguration

type NodeConfiguration internalConf.NodeConfiguration
type NodeConfigurationInput internalConf.NodeConfigurationInput
type ClusterConfiguration internalConf.ClusterConfiguration
type ClusterConfigurationInput internalConf.ClusterConfigurationInput
