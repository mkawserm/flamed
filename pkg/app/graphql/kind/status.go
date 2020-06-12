package kind

import "github.com/graphql-go/graphql"

var GQLStatusEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "Status",
	Description: "`Status` can be either REJECTED or ACCEPTED",
	Values: graphql.EnumValueConfigMap{
		"REJECTED": &graphql.EnumValueConfig{
			Value:       0,
			Description: "REJECTED",
		},
		"ACCEPTED": &graphql.EnumValueConfig{
			Value:       1,
			Description: "ACCEPTED",
		},
	},
})
