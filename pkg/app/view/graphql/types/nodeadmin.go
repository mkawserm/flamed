package types

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/flamed"
)

var NodeAdminType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "NodeAdmin",
	Description: "NodeAdmin of flamed",
	Fields: graphql.Fields{
		"leaderID": &graphql.Field{
			Type:        graphql.NewNonNull(UInt64Type),
			Description: "Current leader id of the cluster",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if nodeAdmin, ok := p.Source.(*flamed.NodeAdmin); ok {
					id, _, _ := nodeAdmin.GetLeaderID()
					return NewUInt64FromUInt64(id), nil
				}
				return NewUInt64FromInt(0), nil
			},
		},

		"appliedIndex": &graphql.Field{
			Type:        graphql.NewNonNull(UInt64Type),
			Description: "Current applied raft index of the cluster",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if nodeAdmin, ok := p.Source.(*flamed.NodeAdmin); ok {
					id, _ := nodeAdmin.GetAppliedIndex()
					return NewUInt64FromUInt64(id), nil
				}
				return NewUInt64FromInt(0), nil
			},
		},

		"hasNodeInfo": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Boolean),
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
				return false, nil
			},
		},
	},
})
