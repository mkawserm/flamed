package nodeadmin

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/flamed"
)

var GQLNodeAdminType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "NodeAdmin",
	Description: "`NodeAdmin` provides all administrative information related to the cluster",
	Fields: graphql.Fields{
		"leaderID": &graphql.Field{
			Type:        kind.GQLUInt64Type,
			Description: "Current leader id of the cluster",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if nodeAdmin, ok := p.Source.(*flamed.NodeAdmin); ok {
					id, _, _ := nodeAdmin.GetLeaderID()
					return kind.NewUInt64FromUInt64(id), nil
				}

				return nil, nil
			},
		},

		"appliedIndex": &graphql.Field{
			Type:        kind.GQLUInt64Type,
			Description: "Current applied raft index of the cluster",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if nodeAdmin, ok := p.Source.(*flamed.NodeAdmin); ok {
					id, _ := nodeAdmin.GetAppliedIndex()
					return kind.NewUInt64FromUInt64(id), nil
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
					Type:        graphql.NewNonNull(kind.GQLUInt64Type),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				nodeID := p.Args["nodeID"].(*kind.UInt64)
				if nodeAdmin, ok := p.Source.(*flamed.NodeAdmin); ok {
					return nodeAdmin.HasNodeInfo(nodeID.Value()), nil
				}

				return nil, nil
			},
		},

		"clusterMembership": &graphql.Field{
			Type:        kind.GQLClusterMembershipType,
			Description: "Cluster membership information",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if nodeAdmin, ok := p.Source.(*flamed.NodeAdmin); ok {
					m, err := nodeAdmin.GetClusterMembership()
					if err != nil {
						return nil, nil
					}
					return m, nil
				}

				return nil, nil
			},
		},
	},
})
