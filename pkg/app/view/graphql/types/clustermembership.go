package types

import (
	"github.com/graphql-go/graphql"
	"github.com/lni/dragonboat/v3"
)

var ClusterMembershipType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ClusterMembership",
	Description: "`ClusterMembership` provides all information related to the cluster",
	Fields: graphql.Fields{
		"configChangeID": &graphql.Field{
			Type:        UInt64Type,
			Description: "Config change ID",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if membership, ok := p.Source.(*dragonboat.Membership); ok {
					return NewUInt64FromUInt64(membership.ConfigChangeID), nil
				}
				return nil, nil
			},
		},

		"nodes": &graphql.Field{
			Type:        graphql.NewList(RaftNodeInfoType),
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
			Type:        graphql.NewList(RaftNodeInfoType),
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
			Type:        graphql.NewList(RaftNodeInfoType),
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
			Type:        graphql.NewList(RaftNodeInfoType),
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
