package types

import (
	"github.com/graphql-go/graphql"
	"github.com/lni/dragonboat/v3"
)

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
