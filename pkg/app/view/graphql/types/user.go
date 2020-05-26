package types

import (
	"encoding/base64"
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/pb"
)

var GQLUserType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "`User` contains all user related information",
	Fields: graphql.Fields{
		"userType": &graphql.Field{
			Name:        "GQLUserType",
			Type:        GQLUserTypeEnum,
			Description: "User type",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u, ok := p.Source.(*pb.User)
				if !ok {
					return nil, nil
				}

				return int(u.UserType), nil
			},
		},

		"roles": &graphql.Field{
			Name:        "Roles",
			Type:        graphql.String,
			Description: "User roles",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u, ok := p.Source.(*pb.User)
				if !ok {
					return nil, nil
				}

				return u.Roles, nil
			},
		},

		"username": &graphql.Field{
			Name:        "Username",
			Type:        graphql.String,
			Description: "Username",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u, ok := p.Source.(*pb.User)
				if !ok {
					return nil, nil
				}

				return u.Username, nil
			},
		},

		"password": &graphql.Field{
			Name:        "Password",
			Type:        graphql.String,
			Description: "User password",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u, ok := p.Source.(*pb.User)
				if !ok {
					return nil, nil
				}

				return u.Password, nil
			},
		},

		"createdAt": &graphql.Field{
			Name:        "CreatedAt",
			Type:        GQLUInt64Type,
			Description: "User creation time in unix nano timestamp",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u, ok := p.Source.(*pb.User)
				if !ok {
					return nil, nil
				}

				return NewUInt64FromUInt64(u.CreatedAt), nil
			},
		},

		"updatedAt": &graphql.Field{
			Name:        "UpdatedAt",
			Type:        GQLUInt64Type,
			Description: "User updated time in unix nano timestamp",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u, ok := p.Source.(*pb.User)
				if !ok {
					return nil, nil
				}

				return NewUInt64FromUInt64(u.UpdatedAt), nil
			},
		},

		"data": &graphql.Field{
			Name:        "Data",
			Type:        graphql.String,
			Description: "Data in base64 format",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u, ok := p.Source.(*pb.User)
				if !ok {
					return nil, nil
				}

				return base64.StdEncoding.EncodeToString(u.Data), nil
			},
		},

		"meta": &graphql.Field{
			Name:        "Meta",
			Type:        graphql.String,
			Description: "Meta data in base64 format",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u, ok := p.Source.(*pb.User)
				if !ok {
					return nil, nil
				}

				return base64.StdEncoding.EncodeToString(u.Meta), nil
			},
		},
	},
})
