package kind

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/pb"
)

var GQLIndexFieldType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "IndexField",
	Description: "`IndexField`",
	Fields: graphql.Fields{
		"indexFieldType": &graphql.Field{
			Name:        "IndexFieldType",
			Description: "",
			Type:        GQLIndexFieldTypeEnum,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexField, ok := p.Source.(*pb.IndexField)
				if !ok {
					return nil, nil
				}
				return int(indexField.IndexFieldType), nil
			},
		},

		"name": &graphql.Field{
			Name:        "Name",
			Description: "",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexField, ok := p.Source.(*pb.IndexField)
				if !ok {
					return nil, nil
				}
				return indexField.Name, nil
			},
		},

		"analyzer": &graphql.Field{
			Name:        "Analyzer",
			Description: "",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexField, ok := p.Source.(*pb.IndexField)
				if !ok {
					return nil, nil
				}
				return indexField.Analyzer, nil
			},
		},

		"enabled": &graphql.Field{
			Name:        "Enabled",
			Description: "",
			Type:        graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexField, ok := p.Source.(*pb.IndexField)
				if !ok {
					return nil, nil
				}
				return indexField.Enabled, nil
			},
		},

		"index": &graphql.Field{
			Name:        "Index",
			Description: "",
			Type:        graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexField, ok := p.Source.(*pb.IndexField)
				if !ok {
					return nil, nil
				}
				return indexField.Index, nil
			},
		},

		"store": &graphql.Field{
			Name:        "Store",
			Description: "",
			Type:        graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexField, ok := p.Source.(*pb.IndexField)
				if !ok {
					return nil, nil
				}
				return indexField.Store, nil
			},
		},

		"includeTermVectors": &graphql.Field{
			Name:        "IncludeTermVectors",
			Description: "",
			Type:        graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexField, ok := p.Source.(*pb.IndexField)
				if !ok {
					return nil, nil
				}
				return indexField.IncludeTermVectors, nil
			},
		},

		"includeInAll": &graphql.Field{
			Name:        "IncludeInAll",
			Description: "",
			Type:        graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexField, ok := p.Source.(*pb.IndexField)
				if !ok {
					return nil, nil
				}
				return indexField.IncludeInAll, nil
			},
		},

		"docValues": &graphql.Field{
			Name:        "DocValues",
			Description: "",
			Type:        graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexField, ok := p.Source.(*pb.IndexField)
				if !ok {
					return nil, nil
				}
				return indexField.DocValues, nil
			},
		},

		"dateFormat": &graphql.Field{
			Name:        "DateFormat",
			Description: "",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexField, ok := p.Source.(*pb.IndexField)
				if !ok {
					return nil, nil
				}
				return indexField.DateFormat, nil
			},
		},
	},
})
