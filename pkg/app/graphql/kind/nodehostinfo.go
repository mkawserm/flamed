package kind

import (
	"github.com/graphql-go/graphql"
	"github.com/lni/dragonboat/v3"
	"github.com/lni/dragonboat/v3/raftio"
)

var GQLLogInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "LogInfo",
	Description: "`LogInfo` is a list of NodeInfo values representing all Raft logs stored on the NodeHost.",
	Fields: graphql.Fields{
		"clusterID": &graphql.Field{
			Type:        GQLUInt64Type,
			Description: "Cluster ID",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if l, ok := p.Source.(raftio.NodeInfo); ok {
					return NewUInt64FromUInt64(l.ClusterID), nil
				}
				return nil, nil
			},
		},

		"nodeID": &graphql.Field{
			Type:        GQLUInt64Type,
			Description: "Node ID",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if l, ok := p.Source.(raftio.NodeInfo); ok {
					return NewUInt64FromUInt64(l.NodeID), nil
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

		"logInfo": &graphql.Field{
			Type:        graphql.NewList(GQLLogInfoType),
			Description: "List of NodeInfo",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if nodeHostInfo, ok := p.Source.(*dragonboat.NodeHostInfo); ok {
					return nodeHostInfo.LogInfo, nil
				}
				return nil, nil
			},
		},

		"clusterInfoList": &graphql.Field{
			Type:        graphql.NewList(GQLClusterInfoType),
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
