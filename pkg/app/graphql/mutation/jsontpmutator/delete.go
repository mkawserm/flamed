package jsontpmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/tp/json"
	"github.com/mkawserm/flamed/pkg/utility"
)

var GQLDelete = &graphql.Field{
	Name:        "Delete",
	Type:        kind.GQLProposalResponseType,
	Description: "",

	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Description: "ID",
			Type:        graphql.NewNonNull(graphql.ID),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := p.Args["id"].(string)
		jsonContext, ok := p.Source.(*json.Context)
		if !ok {
			return nil, nil
		}

		if !utility.HasReadPermission(jsonContext.AccessControl) {
			return nil, gqlerrors.NewFormattedError("read permission required")
		}

		if !utility.HasWritePermission(jsonContext.AccessControl) {
			return nil, gqlerrors.NewFormattedError("write permission required")
		}

		pr, err := jsonContext.Client.DeleteJSONMap(map[string]interface{}{"id": id})

		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}

		return pr, nil
	},
}
