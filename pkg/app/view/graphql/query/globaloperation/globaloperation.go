package globaloperation

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/app/utility"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	fContext "github.com/mkawserm/flamed/pkg/context"
)

var GQLGlobalOperationType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GlobalOperation",
	Description: "Flamed global features",
	Fields: graphql.Fields{
		"search":   Search,
		"iterate":  Iterate,
		"retrieve": Retrieve,
		"searchJ":  SearchJ,
	},
})

func GlobalOperation(flamedContext *fContext.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Name:        "GlobalOperation",
		Type:        GQLGlobalOperationType,
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

			query := flamedContext.Flamed().NewGlobalOperation(clusterID,
				namespace,
				flamedContext.GlobalRequestTimeout())

			return &Context{
				Query:         query,
				AccessControl: accessControl,
			}, nil
		},
	}
}
