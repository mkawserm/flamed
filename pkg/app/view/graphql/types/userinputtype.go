package types

import "github.com/graphql-go/graphql"

var UserInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"userType": &graphql.InputObjectFieldConfig{Type: UserTypeEnum},
		"roles":    &graphql.InputObjectFieldConfig{Type: graphql.String},
		"username": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"password": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"data": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: "Data in base64 encoded string",
		},
		"meta": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: "Meta in base64 encoded string",
		},
	},
})
