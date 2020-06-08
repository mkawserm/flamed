package types

import "github.com/graphql-go/graphql"

type RaftNodeInfo struct {
	NodeID      uint64
	RaftAddress string
}

var GQLRaftNodeInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RaftNodeInfo",
	Description: "`RaftNodeInfo` provides node information of a raft cluster",
	Fields: graphql.Fields{
		"nodeID": &graphql.Field{
			Type:        GQLUInt64Type,
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
