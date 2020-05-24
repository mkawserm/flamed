package types

import (
	"encoding/base64"
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
)

var AccessControlType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "AccessControl",
	Description: "`AccessControl` contains all access control related information of a user",
	Fields: graphql.Fields{
		"namespace": &graphql.Field{
			Name:        "Namespace",
			Type:        graphql.String,
			Description: "Namespace name which user has access to",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u, ok := p.Source.(*pb.AccessControl)
				if !ok {
					return nil, nil
				}

				return string(u.Namespace), nil
			},
		},

		"username": &graphql.Field{
			Name:        "Username",
			Type:        graphql.String,
			Description: "Username of which the access control belongs to",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u, ok := p.Source.(*pb.AccessControl)
				if !ok {
					return nil, nil
				}

				return u.Username, nil
			},
		},

		"permission": &graphql.Field{
			Name:        "Permission",
			Type:        graphql.Int,
			Description: "User permission",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u, ok := p.Source.(*pb.AccessControl)
				if !ok {
					return nil, nil
				}

				if len(u.Permission) == 0 {
					return 0, nil
				}

				return int(u.Permission[0]), nil
			},
		},

		"hasReadPermission": &graphql.Field{
			Name:        "HasReadPermission",
			Type:        graphql.Boolean,
			Description: "Read permission flag",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u, ok := p.Source.(*pb.AccessControl)
				if !ok {
					return nil, nil
				}

				return utility.HasReadPermission(u), nil
			},
		},

		"hasWritePermission": &graphql.Field{
			Name:        "HasWritePermission",
			Type:        graphql.Boolean,
			Description: "Write permission flag",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u, ok := p.Source.(*pb.AccessControl)
				if !ok {
					return nil, nil
				}

				return utility.HasWritePermission(u), nil
			},
		},

		"hasUpdatePermission": &graphql.Field{
			Name:        "HasUpdatePermission",
			Type:        graphql.Boolean,
			Description: "Update permission flag",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u, ok := p.Source.(*pb.AccessControl)
				if !ok {
					return nil, nil
				}

				return utility.HasUpdatePermission(u), nil
			},
		},

		"hasDeletePermission": &graphql.Field{
			Name:        "HasDeletePermission",
			Type:        graphql.Boolean,
			Description: "Delete permission flag",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u, ok := p.Source.(*pb.AccessControl)
				if !ok {
					return nil, nil
				}

				return utility.HasDeletePermission(u), nil
			},
		},

		"createdAt": &graphql.Field{
			Name:        "CreatedAt",
			Type:        UInt64Type,
			Description: "User creation time in unix nano timestamp",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u, ok := p.Source.(*pb.AccessControl)
				if !ok {
					return nil, nil
				}

				return NewUInt64FromUInt64(u.CreatedAt), nil
			},
		},

		"updatedAt": &graphql.Field{
			Name:        "UpdatedAt",
			Type:        UInt64Type,
			Description: "User updated time in unix nano timestamp",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u, ok := p.Source.(*pb.AccessControl)
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
				u, ok := p.Source.(*pb.AccessControl)
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
				u, ok := p.Source.(*pb.AccessControl)
				if !ok {
					return nil, nil
				}

				return base64.StdEncoding.EncodeToString(u.Meta), nil
			},
		},
	},
})
