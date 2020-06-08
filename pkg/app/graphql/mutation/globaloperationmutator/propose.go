package globaloperationmutator

import (
	"encoding/base64"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/types"
	"github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
)

var Propose = &graphql.Field{
	Name:        "Propose",
	Description: "`Propose`",
	Type:        types.GQLProposalResponseType,
	Args: graphql.FieldConfigArgument{
		"proposal": &graphql.ArgumentConfig{
			Description: "Proposal",
			Type:        graphql.NewNonNull(types.GQLProposalInputType),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		ctx, ok := p.Source.(*context.GlobalOperationContext)
		if !ok {
			return nil, nil
		}

		namespace := []byte(ctx.GlobalOperation.Namespace())
		proposalMap := p.Args["proposal"].(map[string]interface{})
		if !utility.HasGlobalCRUDPermission(ctx.AccessControl) {
			return nil, gqlerrors.NewFormattedError("global CRUD permission required")
		}

		meta := []byte("")
		if v, found := proposalMap["meta"]; found {
			v2, err2 := base64.StdEncoding.DecodeString(v.(string))
			if err2 != nil {
				return nil, gqlerrors.NewFormattedError(err2.Error())
			}
			meta = v2
		}

		proposal := pb.NewProposal()
		proposal.Meta = meta
		for _, st := range proposalMap["transactions"].([]interface{}) {
			stm := st.(map[string]interface{})
			payload := stm["payload"].(string)
			familyName := stm["familyName"].(string)
			familyVersion := stm["familyVersion"].(string)

			if payload == "" {
				return nil, gqlerrors.NewFormattedError("Payload can not be empty")
			}
			if familyName == "" {
				return nil, gqlerrors.NewFormattedError("FamilyName can not be empty")
			}
			if familyVersion == "" {
				return nil, gqlerrors.NewFormattedError("FamilyVersion can not be empty")
			}

			payloadBytes, err := base64.StdEncoding.DecodeString(payload)
			if err != nil {
				return nil, gqlerrors.NewFormattedError(err.Error())
			}

			proposal.AddTransaction(namespace, familyName, familyVersion, payloadBytes)
		}

		o, err := ctx.GlobalOperation.Propose(proposal)
		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}
		return o, nil
	},
}
