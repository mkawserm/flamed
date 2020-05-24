package types

import "github.com/graphql-go/graphql"

var StatusEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "Status",
	Description: "`Status` can be either REJECTED or ACCEPTED",
	Values: graphql.EnumValueConfigMap{
		"REJECTED": &graphql.EnumValueConfig{
			Value:       0,
			Description: "Rejected",
		},
		"ACCEPTED": &graphql.EnumValueConfig{
			Value:       1,
			Description: "Accepted",
		},
	},
})
