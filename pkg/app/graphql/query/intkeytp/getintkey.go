package intkeytp

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/tp/intkey"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"
)

var GQLGetIntKey = &graphql.Field{
	Name:        "GetIntKey",
	Type:        kind.GQLIntKeyStateType,
	Description: "",

	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Description: "Name",
			Type:        graphql.NewNonNull(graphql.String),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name := p.Args["name"].(string)

		ikc, ok := p.Source.(*intkey.Context)
		if !ok {
			return nil, nil
		}

		if !utility.HasReadPermission(ikc.AccessControl) {
			return nil, gqlerrors.NewFormattedError(x.ErrReadPermissionRequired.Error())
		}

		intKeyState, err := ikc.Client.GetIntKeyState(name)

		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}

		return intKeyState, nil
	},
}
