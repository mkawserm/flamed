package kind

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/tp/intkey"
)

var GQLIntKeyStateType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "IntKeyState",
	Description: "`IntKeyState`",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Name:        "Name",
			Description: "",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				intKeyState, ok := p.Source.(*intkey.IntKeyState)
				if !ok {
					return nil, nil
				}
				return intKeyState.Name, nil
			},
		},
		"value": &graphql.Field{
			Name:        "Value",
			Description: "",
			Type:        GQLUInt64Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				intKeyState, ok := p.Source.(*intkey.IntKeyState)
				if !ok {
					return nil, nil
				}
				return NewUInt64FromUInt64(intKeyState.Value), nil
			},
		},
	},
})
