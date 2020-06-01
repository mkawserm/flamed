package globaloperationmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/app/utility"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	fContext "github.com/mkawserm/flamed/pkg/context"
)

var GQLGlobalOperationMutatorType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GlobalOperationMutator",
	Description: "`GlobalOperationMutator`",
	Fields: graphql.Fields{
		"propose": Propose,
	},
})

func GlobalOperationMutator(flamedContext *fContext.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Name:        "GlobalOperationMutator",
		Type:        GQLGlobalOperationMutatorType,
		Description: "Mutation in the flamed global scope",

		Args: graphql.FieldConfigArgument{
			"clusterID": &graphql.ArgumentConfig{
				Description: "Cluster ID",
				Type:        graphql.NewNonNull(types.GQLUInt64Type),
			},
			"namespace": &graphql.ArgumentConfig{
				Description: "Namespace",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			clusterID, namespace, accessControl, err := utility.AuthCheck(p, flamedContext)
			if err != nil {
				return nil, err
			}

			globalOperation := flamedContext.Flamed().NewGlobalOperation(clusterID,
				namespace,
				flamedContext.GlobalRequestTimeout())

			return &fContext.GlobalOperationContext{
				GlobalOperation: globalOperation,
				AccessControl:   accessControl,
			}, nil
		},
	}
}
