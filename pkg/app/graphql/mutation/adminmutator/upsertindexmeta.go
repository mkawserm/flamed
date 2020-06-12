package adminmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/app/utility"
	"github.com/mkawserm/flamed/pkg/flamed"
	"time"
)

var UpsertIndexMeta = &graphql.Field{
	Name:        "UpsertIndexMeta",
	Description: "Upsert index meta",
	Type:        kind.GQLProposalResponseType,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Description: "Index meta input",
			Type:        graphql.NewNonNull(kind.GQLIndexMetaInputType),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := p.Args["input"].(map[string]interface{})
		admin, ok := p.Source.(*flamed.Admin)
		if !ok {
			return nil, gqlerrors.NewFormattedError("Unknown source type." +
				" FlamedContext required")
		}

		indexMeta := utility.BuildIndexMetaFromMap(input)
		indexMeta.CreatedAt = uint64(time.Now().UnixNano())
		indexMeta.UpdatedAt = uint64(time.Now().UnixNano())

		pr, err := admin.UpsertIndexMeta(indexMeta)
		if err != nil {
			return nil, err
		}

		return pr, nil
	},
}
