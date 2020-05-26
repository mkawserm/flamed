package intkeymutator

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	fContext "github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/tp/intkey"
	"strings"
)

type intKeyContext struct {
	Client        *intkey.Client
	AccessControl *pb.AccessControl
}

var GQLIntKeyMutatorType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "IntKeyMutator",
	Description: "`IntKeyMutator` provides mutation capability for `IntKey` transaction processor",
	Fields: graphql.Fields{
		"insert": Insert,
	},
})

func IntKeyMutator(flamedContext *fContext.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Name:        "IntKeyMutator",
		Type:        GQLIntKeyMutatorType,
		Description: "",

		Args: graphql.FieldConfigArgument{
			"clusterID": &graphql.ArgumentConfig{
				Description: "Cluster ID",
				Type:        graphql.NewNonNull(types.UInt64Type),
			},
			"namespace": &graphql.ArgumentConfig{
				Description: "Cluster ID",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			clusterID := p.Args["clusterID"].(*types.UInt64)
			namespace := p.Args["namespace"].(string)

			if !flamedContext.Flamed.IsClusterIDAvailable(clusterID.Value()) {
				return nil,
					gqlerrors.NewFormattedError(
						fmt.Sprintf("clusterID [%d] is not available", clusterID.Value()))
			}

			if strings.EqualFold(namespace, "meta") {
				return nil, gqlerrors.NewFormattedError("meta namespace is reserved")
			}

			//if p.Context.Value("GraphQLContext") == nil {
			//	return nil, nil
			//}

			gqlContext := p.Context.Value("GraphQLContext").(*fContext.GraphQLContext)
			admin := flamedContext.Flamed.NewAdmin(clusterID.Value(), flamedContext.GlobalRequestTimeout)

			// Authenticate user
			username, password := gqlContext.GetUsernameAndPasswordFromAuth()
			if len(username) == 0 || len(password) == 0 {
				return nil, gqlerrors.NewFormattedError("Access denied." +
					" Only authenticated user can access")
			}
			if !gqlContext.IsUserPasswordValid(admin, username, password) {
				return nil, gqlerrors.NewFormattedError("Access denied." +
					" Only authenticated user can access")
			}

			accessControl, err := admin.GetAccessControl(username, []byte(namespace))
			if err != nil {
				accessControl, err = admin.GetAccessControl(username, []byte("*"))
				if err != nil {
					return nil, gqlerrors.NewFormattedError("namespace access control not found")
				}
			}

			intKeyClient := &intkey.Client{}
			err = intKeyClient.Setup(clusterID.Value(),
				namespace,
				flamedContext.Flamed,
				flamedContext.GlobalRequestTimeout)

			if err != nil {
				return nil, gqlerrors.NewFormattedError(err.Error())
			}

			return &intKeyContext{
				AccessControl: accessControl,
				Client:        intKeyClient,
			}, nil
		},
	}
}
