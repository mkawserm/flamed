package types

import (
	"github.com/graphql-go/graphql"
	"github.com/lni/dragonboat/v3"
	"github.com/mkawserm/flamed/pkg/flamed"
)

type RaftNodeInfo struct {
	NodeID      uint64
	RaftAddress string
}

var RaftNodeInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RaftNodeInfo",
	Description: "`RaftNodeInfo` provides node information of a raft cluster",
	Fields: graphql.Fields{
		"nodeID": &graphql.Field{
			Type:        UInt64Type,
			Description: "Node id",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if nodeInfo, ok := p.Source.(RaftNodeInfo); ok {
					return NewUInt64FromUInt64(nodeInfo.NodeID), nil
				}
				return nil, nil
			},
		},
		"raftAddress": &graphql.Field{
			Type:        graphql.String,
			Description: "Raft address",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if nodeInfo, ok := p.Source.(RaftNodeInfo); ok {
					return nodeInfo.RaftAddress, nil
				}
				return nil, nil
			},
		},
	},
})

var ClusterInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ClusterInfo",
	Description: "`ClusterInfo` provides all information related to raft cluster",
	Fields: graphql.Fields{
		"clusterID": &graphql.Field{
			Type:        UInt64Type,
			Description: "Cluster id",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if clusterInfo, ok := p.Source.(dragonboat.ClusterInfo); ok {
					return NewUInt64FromUInt64(clusterInfo.ClusterID), nil
				}
				return nil, nil
			},
		},
		"nodeID": &graphql.Field{
			Type:        UInt64Type,
			Description: "Node id",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if clusterInfo, ok := p.Source.(dragonboat.ClusterInfo); ok {
					return NewUInt64FromUInt64(clusterInfo.NodeID), nil
				}
				return nil, nil
			},
		},
		"nodes": &graphql.Field{
			Type:        graphql.NewList(RaftNodeInfoType),
			Description: "Raft node information list",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if clusterInfo, ok := p.Source.(dragonboat.ClusterInfo); ok {
					var raftNodeInfoList []RaftNodeInfo
					for k, v := range clusterInfo.Nodes {
						raftNodeInfoList = append(raftNodeInfoList, RaftNodeInfo{
							NodeID:      k,
							RaftAddress: v,
						})
					}
					return raftNodeInfoList, nil
				}
				return nil, nil
			},
		},
		"configChangeIndex": &graphql.Field{
			Type:        UInt64Type,
			Description: "Configuration change index",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if clusterInfo, ok := p.Source.(dragonboat.ClusterInfo); ok {
					return NewUInt64FromUInt64(clusterInfo.ConfigChangeIndex), nil
				}
				return nil, nil
			},
		},
		"stateMachineType": &graphql.Field{
			Type:        UInt64Type,
			Description: "State machine type",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if clusterInfo, ok := p.Source.(dragonboat.ClusterInfo); ok {
					return NewUInt64FromUInt64(uint64(clusterInfo.StateMachineType)), nil
				}
				return nil, nil
			},
		},
		"isLeader": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Is the raft node a leader?",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if clusterInfo, ok := p.Source.(dragonboat.ClusterInfo); ok {
					return clusterInfo.IsLeader, nil
				}
				return nil, nil
			},
		},
		"isObserver": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Is the raft node a observer?",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if clusterInfo, ok := p.Source.(dragonboat.ClusterInfo); ok {
					return clusterInfo.IsObserver, nil
				}
				return nil, nil
			},
		},
		"isWitness": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Is the raft node a witness?",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if clusterInfo, ok := p.Source.(dragonboat.ClusterInfo); ok {
					return clusterInfo.IsWitness, nil
				}
				return nil, nil
			},
		},
		"pending": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Is the raft node pending?",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if clusterInfo, ok := p.Source.(dragonboat.ClusterInfo); ok {
					return clusterInfo.Pending, nil
				}
				return nil, nil
			},
		},
	},
})

var NodeHostInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "NodeHostInfo",
	Description: "`NodeHostInfo` provides all information related to raft host node",
	Fields: graphql.Fields{
		"raftAddress": &graphql.Field{
			Type:        graphql.String,
			Description: "Raft address",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if nodeHostInfo, ok := p.Source.(*dragonboat.NodeHostInfo); ok {
					return nodeHostInfo.RaftAddress, nil
				}
				return nil, nil
			},
		},

		"clusterInfoList": &graphql.Field{
			Type:        graphql.NewList(ClusterInfoType),
			Description: "Cluster information list",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if nodeHostInfo, ok := p.Source.(*dragonboat.NodeHostInfo); ok {
					return nodeHostInfo.ClusterInfoList, nil
				}
				return nil, nil
			},
		},
	},
})

var NodeAdminType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "NodeAdmin",
	Description: "`NodeAdmin` provides all administrative information related to raft node",
	Fields: graphql.Fields{
		"leaderID": &graphql.Field{
			Type:        UInt64Type,
			Description: "Current leader id of the cluster",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if nodeAdmin, ok := p.Source.(*flamed.NodeAdmin); ok {
					id, _, _ := nodeAdmin.GetLeaderID()
					return NewUInt64FromUInt64(id), nil
				}

				return nil, nil
			},
		},

		"appliedIndex": &graphql.Field{
			Type:        UInt64Type,
			Description: "Current applied raft index of the cluster",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if nodeAdmin, ok := p.Source.(*flamed.NodeAdmin); ok {
					id, _ := nodeAdmin.GetAppliedIndex()
					return NewUInt64FromUInt64(id), nil
				}
				return nil, nil
			},
		},

		"hasNodeInfo": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Has node information with the provided nodeID",
			Args: graphql.FieldConfigArgument{
				"nodeID": &graphql.ArgumentConfig{
					Description: "Node id",
					Type:        graphql.NewNonNull(UInt64Type),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				nodeID := p.Args["nodeID"].(*UInt64)
				if nodeAdmin, ok := p.Source.(*flamed.NodeAdmin); ok {
					return nodeAdmin.HasNodeInfo(nodeID.Value()), nil
				}

				return nil, nil
			},
		},

		"nodeHostInfo": &graphql.Field{
			Type:        NodeHostInfoType,
			Description: "Node host information",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if nodeAdmin, ok := p.Source.(*flamed.NodeAdmin); ok {
					return nodeAdmin.GetNodeHostInfo(), nil
				}

				return nil, nil
			},
		},
	},
})
