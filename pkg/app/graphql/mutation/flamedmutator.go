package mutation

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/app/graphql/mutation/adminmutator"
	"github.com/mkawserm/flamed/pkg/app/graphql/mutation/nodeadminmutator"
	fContext "github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/x"
)

var GQLFlamedMutatorType = graphql.NewObject(graphql.ObjectConfig{
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
					Type:        graphql.NewNonNull(kind.GQLUInt64Type),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				clusterID := p.Args["clusterID"].(*kind.UInt64)
				fc, ok := p.Source.(*fContext.FlamedContext)
				if !ok {
					return nil, nil
				}

				if !fc.Flamed().IsClusterIDAvailable(clusterID.Value()) {
					return nil,
						gqlerrors.NewFormattedError(x.ErrClusterIsNotAvailable.Error())
				}
				return fc.Flamed().NewNodeAdmin(clusterID.Value(), fc.GlobalRequestTimeout()), nil
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
					Type:        graphql.NewNonNull(kind.GQLUInt64Type),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				clusterID := p.Args["clusterID"].(*kind.UInt64)
				fc, ok := p.Source.(*fContext.FlamedContext)
				if !ok {
					return nil, nil
				}

				if !fc.Flamed().IsClusterIDAvailable(clusterID.Value()) {
					return nil,
						gqlerrors.NewFormattedError(x.ErrClusterIsNotAvailable.Error())
				}

				return fc.Flamed().NewAdmin(clusterID.Value(), fc.GlobalRequestTimeout()), nil
			},
		},
	},
})

func FlamedMutator(flamedContext *fContext.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Name:        "FlamedMutator",
		Type:        GQLFlamedMutatorType,
		Description: "Flamed mutator helps to modify cluster",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if p.Context.Value("GraphQLContext") == nil {
				return nil, nil
			}

			gqlContext := p.Context.Value("GraphQLContext").(*fContext.AuthContext)
			if !gqlContext.AuthenticateSuperUser(flamedContext.Flamed().NewAdmin(
				1,
				flamedContext.GlobalRequestTimeout())) {
				return nil, gqlerrors.NewFormattedError(x.ErrAccessDenied.Error())
			}

			return flamedContext, nil
		},
	}
}
