package kind

import (
	"encoding/base64"
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/pb"
)

var GQLStateEntry = graphql.NewObject(graphql.ObjectConfig{
	Name:        "StateEntry",
	Description: "`StateEntry`",
	Fields: graphql.Fields{
		"payload": &graphql.Field{
			Name:        "Payload",
			Type:        graphql.String,
			Description: "Payload in base64 format",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				stateEntry, ok := p.Source.(*pb.StateEntry)
				if !ok {
					return nil, nil
				}

				return base64.StdEncoding.EncodeToString(stateEntry.Payload), nil
			},
		},

		"namespace": &graphql.Field{
			Name:        "Namespace",
			Type:        graphql.String,
			Description: "Namespace",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				stateEntry, ok := p.Source.(*pb.StateEntry)
				if !ok {
					return nil, nil
				}

				return string(stateEntry.Namespace), nil
			},
		},

		"familyName": &graphql.Field{
			Name:        "FamilyName",
			Type:        graphql.String,
			Description: "FamilyName",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				stateEntry, ok := p.Source.(*pb.StateEntry)
				if !ok {
					return nil, nil
				}

				return stateEntry.FamilyName, nil
			},
		},

		"familyVersion": &graphql.Field{
			Name:        "FamilyVersion",
			Type:        graphql.String,
			Description: "FamilyVersion",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				stateEntry, ok := p.Source.(*pb.StateEntry)
				if !ok {
					return nil, nil
				}

				return stateEntry.FamilyVersion, nil
			},
		},
	},
})

var GQLStateEntryResponse = graphql.NewObject(graphql.ObjectConfig{
	Name:        "StateEntryResponse",
	Description: "`StateEntryResponse`",
	Fields: graphql.Fields{
		"stateAvailable": &graphql.Field{
			Name:        "StateAvailable",
			Description: "State availability flag",
			Type:        graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				ser, ok := p.Source.(*pb.StateEntryResponse)
				if !ok {
					return nil, nil
				}
				return ser.StateAvailable, nil
			},
		},
		"address": &graphql.Field{
			Name:        "Address",
			Description: "Address",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				ser, ok := p.Source.(*pb.StateEntryResponse)
				if !ok {
					return nil, nil
				}
				return ser.Address, nil
			},
		},

		"stateEntry": &graphql.Field{
			Name:        "StateEntry",
			Description: "StateEntry",
			Type:        GQLStateEntry,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				ser, ok := p.Source.(*pb.StateEntryResponse)
				if !ok {
					return nil, nil
				}
				return ser.StateEntry, nil
			},
		},
	},
})
