package jsontpmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	fContext "github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/tp/json"
	"github.com/mkawserm/flamed/pkg/x"
	"strings"
)

var GQLJSONTPMutatorType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "JSONTPMutator",
	Description: "",
	Fields: graphql.Fields{
		"merge":        GQLMerge,
		"insert":       GQLInsert,
		"upsert":       GQLUpsert,
		"update":       GQLUpdate,
		"delete":       GQLDelete,
		"batchProcess": GQLBatchProcess,
	},
})

func JSONTPMutator(flamedContext *fContext.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Name:        "JSONTPMutator",
		Type:        GQLJSONTPMutatorType,
		Description: "",

		Args: graphql.FieldConfigArgument{
			"clusterID": &graphql.ArgumentConfig{
				Description: "Cluster ID",
				Type:        graphql.NewNonNull(kind.GQLUInt64Type),
			},
			"namespace": &graphql.ArgumentConfig{
				Description: "Namespace",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			clusterID := p.Args["clusterID"].(*kind.UInt64)
			namespace := p.Args["namespace"].(string)

			if !flamedContext.Flamed().IsClusterIDAvailable(clusterID.Value()) {
				return nil,
					gqlerrors.NewFormattedError(x.ErrClusterIsNotAvailable.Error())
			}

			if strings.EqualFold(namespace, "meta") {
				return nil, gqlerrors.NewFormattedError(x.ErrMetaNamespaceIsReserved.Error())
			}

			//if p.Context.Value("GraphQLContext") == nil {
			//	return nil, nil
			//}

			gqlContext := p.Context.Value("GraphQLContext").(*fContext.AuthContext)
			admin := flamedContext.Flamed().NewAdmin(clusterID.Value(), flamedContext.GlobalRequestTimeout())

			// Authenticate user
			username, password := gqlContext.GetUsernameAndPasswordFromAuth()
			if len(username) == 0 || len(password) == 0 {
				return nil, gqlerrors.NewFormattedError(x.ErrAccessDenied.Error())
			}
			if !gqlContext.IsUserPasswordValid(admin, username, password) {
				return nil, gqlerrors.NewFormattedError(x.ErrAccessDenied.Error())
			}

			accessControl, err := admin.GetAccessControl(username, []byte(namespace))
			if err != nil {
				accessControl, err = admin.GetAccessControl(username, []byte("*"))
				if err != nil {
					return nil, gqlerrors.NewFormattedError(x.ErrAccessControlNotFound.Error())
				}
			}

			jsonClient := &json.Client{}
			err = jsonClient.Setup(clusterID.Value(),
				namespace,
				flamedContext.Flamed(),
				flamedContext.GlobalRequestTimeout())

			if err != nil {
				return nil, gqlerrors.NewFormattedError(err.Error())
			}

			return &json.Context{
				AccessControl: accessControl,
				Client:        jsonClient,
			}, nil
		},
	}
}
