package query

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	flamedContext "github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/flamed"
)

var NodeAdminType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "NodeAdmin",
	Description: "`NodeAdmin` provides all administrative information related to the raft node",
	Fields: graphql.Fields{
		"leaderID": &graphql.Field{
			Type:        types.UInt64Type,
			Description: "Current leader id of the cluster",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if nodeAdmin, ok := p.Source.(*flamed.NodeAdmin); ok {
					id, _, _ := nodeAdmin.GetLeaderID()
					return types.NewUInt64FromUInt64(id), nil
				}

				return nil, nil
			},
		},

		"appliedIndex": &graphql.Field{
			Type:        types.UInt64Type,
			Description: "Current applied raft index of the cluster",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if nodeAdmin, ok := p.Source.(*flamed.NodeAdmin); ok {
					id, _ := nodeAdmin.GetAppliedIndex()
					return types.NewUInt64FromUInt64(id), nil
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
					Type:        graphql.NewNonNull(types.UInt64Type),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				nodeID := p.Args["nodeID"].(*types.UInt64)
				if nodeAdmin, ok := p.Source.(*flamed.NodeAdmin); ok {
					return nodeAdmin.HasNodeInfo(nodeID.Value()), nil
				}

				return nil, nil
			},
		},

		"nodeHostInfo": &graphql.Field{
			Type:        types.NodeHostInfoType,
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

func NodeAdmin(flamedContext *flamedContext.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Name:        "NodeAdmin",
		Type:        graphql.NewNonNull(NodeAdminType),
		Description: "Get NodeAdmin by clusterID",
		Args: graphql.FieldConfigArgument{
			"clusterID": &graphql.ArgumentConfig{
				Description: "Cluster id of the NodeAdmin",
				Type:        graphql.NewNonNull(types.UInt64Type),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			//TODO: Request must be authenticated

			clusterID := p.Args["clusterID"].(*types.UInt64)
			if !flamedContext.Flamed.IsClusterIDAvailable(clusterID.Value()) {
				return nil, gqlerrors.NewFormattedError("There is no NodeAdmin with the provided clusterID")
			}
			return flamedContext.Flamed.NewNodeAdmin(clusterID.Value(), flamedContext.GlobalRequestTimeout), nil
		},
	}
}
