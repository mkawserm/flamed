package kind

import (
	"github.com/graphql-go/graphql"
	"github.com/lni/dragonboat/v3"
)

var GQLClusterMembershipType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ClusterMembership",
	Description: "`ClusterMembership` provides all information related to the cluster",
	Fields: graphql.Fields{
		"configChangeID": &graphql.Field{
			Type:        GQLUInt64Type,
			Description: "Config change ID",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if membership, ok := p.Source.(*dragonboat.Membership); ok {
					return NewUInt64FromUInt64(membership.ConfigChangeID), nil
				}
				return nil, nil
			},
		},

		"nodes": &graphql.Field{
			Type:        graphql.NewList(GQLRaftNodeInfoType),
			Description: "Node information list",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if membership, ok := p.Source.(*dragonboat.Membership); ok {
					var raftNodeInfoList []RaftNodeInfo
					for k, v := range membership.Nodes {
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

		"observers": &graphql.Field{
			Type:        graphql.NewList(GQLRaftNodeInfoType),
			Description: "Observers information list",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if membership, ok := p.Source.(*dragonboat.Membership); ok {
					var raftNodeInfoList []RaftNodeInfo
					for k, v := range membership.Observers {
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

		"witnesses": &graphql.Field{
			Type:        graphql.NewList(GQLRaftNodeInfoType),
			Description: "Witnesses information list",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if membership, ok := p.Source.(*dragonboat.Membership); ok {
					var raftNodeInfoList []RaftNodeInfo
					for k, v := range membership.Witnesses {
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

		"removed": &graphql.Field{
			Type:        graphql.NewList(GQLRaftNodeInfoType),
			Description: "Removed node information list",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if membership, ok := p.Source.(*dragonboat.Membership); ok {
					var raftNodeInfoList []RaftNodeInfo
					for k := range membership.Removed {
						raftNodeInfoList = append(raftNodeInfoList, RaftNodeInfo{
							NodeID: k,
						})
					}
					return raftNodeInfoList, nil
				}
				return nil, nil
			},
		},
	},
})
