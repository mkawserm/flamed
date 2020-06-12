package jsontp

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/tp/json"
	"github.com/mkawserm/flamed/pkg/utility"
)

var GQLGet = &graphql.Field{
	Name:        "Get",
	Type:        kind.GQLJSONType,
	Description: "",

	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Description: "ID",
			Type:        graphql.NewNonNull(graphql.ID),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := p.Args["id"].(string)

		ikc, ok := p.Source.(*json.Context)
		if !ok {
			return nil, nil
		}

		if !utility.HasReadPermission(ikc.AccessControl) {
			return nil, gqlerrors.NewFormattedError("read permission required")
		}

		obj := make(map[string]interface{})

		_, err := ikc.Client.Get(id, &obj)
		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}
		return obj, nil
	},
}
