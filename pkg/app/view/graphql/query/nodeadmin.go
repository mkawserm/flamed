package query

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	flamedContext "github.com/mkawserm/flamed/pkg/context"
)

func NodeAdmin(flamedContext *flamedContext.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Name:        "NodeAdmin",
		Type:        graphql.NewNonNull(types.NodeAdminType),
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
