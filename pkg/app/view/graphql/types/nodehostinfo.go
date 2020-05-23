package types

import (
	"github.com/graphql-go/graphql"
	"github.com/lni/dragonboat/v3"
)

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
