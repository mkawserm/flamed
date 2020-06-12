package kind

import "github.com/graphql-go/graphql"

var GQLTransactionInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "TransactionInput",
	Description: "Transaction input",
	Fields: graphql.InputObjectConfigFieldMap{
		"payload": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Payload in base64 string",
		},
		"familyName": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Family name",
		},
		"familyVersion": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Family version",
		},
	},
})

var GQLProposalInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name:        "ProposalInput",
		Description: "Proposal input",
		Fields: graphql.InputObjectConfigFieldMap{
			"meta": &graphql.InputObjectFieldConfig{
				Type:        graphql.String,
				Description: "Meta in base64 string",
			},
			"transactions": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(GQLTransactionInputType))),
				Description: "Transactions",
			},
		},
	},
)
