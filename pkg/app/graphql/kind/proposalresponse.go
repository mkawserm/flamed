package kind

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
)

var GQLProposalResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ProposalResponse",
	Description: "`ProposalResponse` gives detail information about a proposal",
	Fields: graphql.Fields{
		"uuid": &graphql.Field{
			Name:        "Uuid",
			Description: "Proposal uuid",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				pr, ok := p.Source.(*pb.ProposalResponse)
				if !ok {
					return nil, nil
				}
				return utility.UUIDToString(pr.Uuid), nil
			},
		},

		"status": &graphql.Field{
			Name:        "Status",
			Description: "Proposal status",
			Type:        GQLStatusEnum,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				pr, ok := p.Source.(*pb.ProposalResponse)
				if !ok {
					return nil, nil
				}
				return int(pr.Status), nil
			},
		},

		"errorCode": &graphql.Field{
			Name:        "ErrorCode",
			Description: "Proposal error code if any",
			Type:        graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				pr, ok := p.Source.(*pb.ProposalResponse)
				if !ok {
					return nil, nil
				}
				return pr.ErrorCode, nil
			},
		},

		"errorText": &graphql.Field{
			Name:        "ErrorText",
			Description: "Proposal error text if any",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				pr, ok := p.Source.(*pb.ProposalResponse)
				if !ok {
					return nil, nil
				}
				return pr.ErrorText, nil
			},
		},

		"transactionResponses": &graphql.Field{
			Name:        "TransactionResponses",
			Description: "Transaction responses",
			Type:        graphql.NewList(GQLTransactionResponseType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				pr, ok := p.Source.(*pb.ProposalResponse)
				if !ok {
					return nil, nil
				}
				return pr.TransactionResponses, nil
			},
		},
	},
})
