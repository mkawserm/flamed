package jsontpmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/tp/json"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"
)

var GQLUpdate = &graphql.Field{
	Name:        "Update",
	Type:        kind.GQLProposalResponseType,
	Description: "",

	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Description: "Input object",
			Type:        graphql.NewNonNull(kind.GQLJSONType),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := p.Args["input"].(map[string]interface{})
		jsonContext, ok := p.Source.(*json.Context)
		if !ok {
			return nil, nil
		}

		if !utility.HasUpdatePermission(jsonContext.AccessControl) {
			return nil, gqlerrors.NewFormattedError(x.ErrUpdatePermissionRequired.Error())
		}

		if !utility.HasReadPermission(jsonContext.AccessControl) {
			return nil, gqlerrors.NewFormattedError(x.ErrReadPermissionRequired.Error())
		}

		if !utility.HasWritePermission(jsonContext.AccessControl) {
			return nil, gqlerrors.NewFormattedError(x.ErrWritePermissionRequired.Error())
		}

		pr, err := jsonContext.Client.UpdateJSONMap(input)

		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}

		return pr, nil
	},
}
