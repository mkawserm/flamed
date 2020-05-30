package global

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/app/utility"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	fContext "github.com/mkawserm/flamed/pkg/context"
)

var GQLGlobalType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Global",
	Description: "Flamed global features",
	Fields: graphql.Fields{
		"searchJ": SearchJ,
	},
})

func Global(flamedContext *fContext.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Name:        "Global",
		Type:        GQLGlobalType,
		Description: "Flamed global features",

		Args: graphql.FieldConfigArgument{
			"clusterID": &graphql.ArgumentConfig{
				Description: "Cluster ID",
				Type:        graphql.NewNonNull(types.GQLUInt64Type),
			},
			"namespace": &graphql.ArgumentConfig{
				Description: "Cluster ID",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			clusterID, namespace, accessControl, err := utility.AuthCheck(p, flamedContext)
			if err != nil {
				return nil, err
			}

			query := flamedContext.Flamed().NewQuery(clusterID,
				namespace,
				flamedContext.GlobalRequestTimeout())

			return &Context{
				Query:         query,
				AccessControl: accessControl,
			}, nil
		},
	}
}
