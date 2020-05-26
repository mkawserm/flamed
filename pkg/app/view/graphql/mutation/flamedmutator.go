package mutation

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/mutation/adminmutator"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/mutation/nodeadminmutator"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	flamedContext "github.com/mkawserm/flamed/pkg/context"
)

var FlamedMutatorType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "FlamedMutator",
	Description: "`FlamedMutator` gives the ability to perform any tasks related to the cluster",
	Fields: graphql.Fields{

		// NodeAdminMutator
		"nodeAdminMutator": &graphql.Field{
			Name:        "NodeAdminMutator",
			Type:        nodeadminmutator.GQLNodeAdminMutatorType,
			Description: "Perform administrative tasks using NodeAdmin by clusterID",
			Args: graphql.FieldConfigArgument{
				"clusterID": &graphql.ArgumentConfig{
					Description: "Cluster ID",
					Type:        graphql.NewNonNull(types.UInt64Type),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				clusterID := p.Args["clusterID"].(*types.UInt64)
				fc, ok := p.Source.(*flamedContext.FlamedContext)
				if !ok {
					return nil, nil
				}

				if !fc.Flamed.IsClusterIDAvailable(clusterID.Value()) {
					return nil,
						gqlerrors.NewFormattedError(
							fmt.Sprintf("clusterID [%d] is not available", clusterID.Value()))
				}
				return fc.Flamed.NewNodeAdmin(clusterID.Value(), fc.GlobalRequestTimeout), nil
			},
		},

		// AdminMutator
		"adminMutator": &graphql.Field{
			Name:        "AdminMutator",
			Type:        adminmutator.GQLAdminMutatorType,
			Description: "",
			Args: graphql.FieldConfigArgument{
				"clusterID": &graphql.ArgumentConfig{
					Description: "Cluster ID",
					Type:        graphql.NewNonNull(types.UInt64Type),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				clusterID := p.Args["clusterID"].(*types.UInt64)
				fc, ok := p.Source.(*flamedContext.FlamedContext)
				if !ok {
					return nil, nil
				}

				if !fc.Flamed.IsClusterIDAvailable(clusterID.Value()) {
					return nil,
						gqlerrors.NewFormattedError(
							fmt.Sprintf("clusterID [%d] is not available", clusterID.Value()))
				}

				return fc.Flamed.NewAdmin(clusterID.Value(), fc.GlobalRequestTimeout), nil
			},
		},
	},
})

func FlamedMutator(flamedContext *flamedContext.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Name:        "FlamedMutator",
		Type:        FlamedMutatorType,
		Description: "Flamed mutator helps to modify cluster",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			//TODO: Request must be authenticated

			return flamedContext, nil
		},
	}
}
