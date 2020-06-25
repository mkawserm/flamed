package intkeytp

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/tp/intkey"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"
)

var GQLGetIntKeyList = &graphql.Field{
	Name:        "GetIntKeyList",
	Type:        graphql.NewList(kind.GQLIntKeyStateType),
	Description: "",

	Args: graphql.FieldConfigArgument{
		"nameList": &graphql.ArgumentConfig{
			Description: "Name list",
			Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.String))),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		nameList := p.Args["nameList"].([]interface{})
		ikc, ok := p.Source.(*intkey.Context)
		if !ok {
			return nil, nil
		}

		if !utility.HasReadPermission(ikc.AccessControl) {
			return nil, gqlerrors.NewFormattedError(x.ErrReadPermissionRequired.Error())
		}

		intKeyStateList, err := ikc.Client.GetIntKeyStateList(nameList)

		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}

		return intKeyStateList, nil
	},
}
