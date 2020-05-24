package types

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/pb"
)

var TransactionResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "TransactionResponse",
	Description: "`TransactionResponse` gives detail information about a transaction",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Name:        "Status",
			Description: "Transaction status",
			Type:        StatusEnum,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				txr, ok := p.Source.(*pb.TransactionResponse)
				if !ok {
					return nil, nil
				}
				return txr.Status, nil
			},
		},

		"errorCode": &graphql.Field{
			Name:        "ErrorCode",
			Description: "Transaction error code if any",
			Type:        graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				txr, ok := p.Source.(*pb.TransactionResponse)
				if !ok {
					return nil, nil
				}
				return txr.ErrorCode, nil
			},
		},

		"errorText": &graphql.Field{
			Name:        "ErrorText",
			Description: "Transaction error text if any",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				txr, ok := p.Source.(*pb.TransactionResponse)
				if !ok {
					return nil, nil
				}
				return txr.ErrorText, nil
			},
		},

		"familyName": &graphql.Field{
			Name:        "FamilyName",
			Description: "Transaction family name",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				txr, ok := p.Source.(*pb.TransactionResponse)
				if !ok {
					return nil, nil
				}
				return txr.FamilyName, nil
			},
		},

		"familyVersion": &graphql.Field{
			Name:        "FamilyVersion",
			Description: "Transaction family version",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				txr, ok := p.Source.(*pb.TransactionResponse)
				if !ok {
					return nil, nil
				}
				return txr.FamilyVersion, nil
			},
		},
	},
})
