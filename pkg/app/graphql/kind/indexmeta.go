package kind

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/pb"
)

var GQLIndexDocumentType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "IndexDocument",
	Description: "`IndexDocument`",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Name:        "Name",
			Description: "",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexDocument, ok := p.Source.(*pb.IndexDocument)
				if !ok {
					return nil, nil
				}

				return indexDocument.Name, nil
			},
		},

		"enabled": &graphql.Field{
			Name:        "Enabled",
			Description: "",
			Type:        graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexDocument, ok := p.Source.(*pb.IndexDocument)
				if !ok {
					return nil, nil
				}
				return indexDocument.Enabled, nil
			},
		},

		"defaultAnalyzer": &graphql.Field{
			Name:        "DefaultAnalyzer",
			Description: "",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexDocument, ok := p.Source.(*pb.IndexDocument)
				if !ok {
					return nil, nil
				}

				return indexDocument.DefaultAnalyzer, nil
			},
		},

		"dynamic": &graphql.Field{
			Name:        "Dynamic",
			Description: "",
			Type:        graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexDocument, ok := p.Source.(*pb.IndexDocument)
				if !ok {
					return nil, nil
				}

				return indexDocument.Dynamic, nil
			},
		},

		"indexFieldList": &graphql.Field{
			Name:        "IndexFieldList",
			Description: "",
			Type:        graphql.NewList(GQLIndexFieldType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexDocument, ok := p.Source.(*pb.IndexDocument)
				if !ok {
					return nil, nil
				}

				return indexDocument.IndexFieldList, nil
			},
		},
	},
})

var IndexMetaType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "IndexMeta",
	Description: "`IndexMeta` contains information related to indexing rules",
	Fields: graphql.Fields{
		"namespace": &graphql.Field{
			Name:        "Namespace",
			Description: "The namespace of which the indexing rule is for",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexMeta, ok := p.Source.(*pb.IndexMeta)
				if !ok {
					return nil, nil
				}

				return string(indexMeta.Namespace), nil
			},
		},

		"version": &graphql.Field{
			Name:        "Version",
			Description: "Index meta version",
			Type:        GQLUInt64Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexMeta, ok := p.Source.(*pb.IndexMeta)
				if !ok {
					return nil, nil
				}

				return NewUInt64FromUInt64(indexMeta.Version), nil
			},
		},

		"enabled": &graphql.Field{
			Name:        "Enabled",
			Description: "",
			Type:        graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexMeta, ok := p.Source.(*pb.IndexMeta)
				if !ok {
					return nil, nil
				}

				return indexMeta.Enabled, nil
			},
		},

		"default": &graphql.Field{
			Name:        "Default",
			Description: "",
			Type:        graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexMeta, ok := p.Source.(*pb.IndexMeta)
				if !ok {
					return nil, nil
				}

				return indexMeta.Default, nil
			},
		},

		"indexDynamic": &graphql.Field{
			Name:        "IndexDynamic",
			Description: "",
			Type:        graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexMeta, ok := p.Source.(*pb.IndexMeta)
				if !ok {
					return nil, nil
				}

				return indexMeta.IndexDynamic, nil
			},
		},

		"storeDynamic": &graphql.Field{
			Name:        "StoreDynamic",
			Description: "",
			Type:        graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexMeta, ok := p.Source.(*pb.IndexMeta)
				if !ok {
					return nil, nil
				}

				return indexMeta.StoreDynamic, nil
			},
		},

		"docValuesDynamic": &graphql.Field{
			Name:        "DocValuesDynamic",
			Description: "",
			Type:        graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexMeta, ok := p.Source.(*pb.IndexMeta)
				if !ok {
					return nil, nil
				}

				return indexMeta.DocValuesDynamic, nil
			},
		},

		"defaultType": &graphql.Field{
			Name:        "DefaultType",
			Description: "",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexMeta, ok := p.Source.(*pb.IndexMeta)
				if !ok {
					return nil, nil
				}

				return indexMeta.DefaultType, nil
			},
		},

		"defaultAnalyzer": &graphql.Field{
			Name:        "DefaultAnalyzer",
			Description: "",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexMeta, ok := p.Source.(*pb.IndexMeta)
				if !ok {
					return nil, nil
				}

				return indexMeta.DefaultAnalyzer, nil
			},
		},

		"defaultDateTimeParser": &graphql.Field{
			Name:        "DefaultDateTimeParser",
			Description: "",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexMeta, ok := p.Source.(*pb.IndexMeta)
				if !ok {
					return nil, nil
				}

				return indexMeta.DefaultDateTimeParser, nil
			},
		},

		"defaultField": &graphql.Field{
			Name:        "DefaultField",
			Description: "",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexMeta, ok := p.Source.(*pb.IndexMeta)
				if !ok {
					return nil, nil
				}

				return indexMeta.DefaultField, nil
			},
		},

		"typeField": &graphql.Field{
			Name:        "TypeField",
			Description: "",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexMeta, ok := p.Source.(*pb.IndexMeta)
				if !ok {
					return nil, nil
				}

				return indexMeta.TypeField, nil
			},
		},

		"customAnalysis": &graphql.Field{
			Name:        "CustomAnalysis",
			Description: "",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexMeta, ok := p.Source.(*pb.IndexMeta)
				if !ok {
					return nil, nil
				}

				return indexMeta.CustomAnalysis, nil
			},
		},

		"createdAt": &graphql.Field{
			Name:        "CreatedAt",
			Description: "",
			Type:        GQLUInt64Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexMeta, ok := p.Source.(*pb.IndexMeta)
				if !ok {
					return nil, nil
				}

				return NewUInt64FromUInt64(indexMeta.CreatedAt), nil
			},
		},

		"updatedAt": &graphql.Field{
			Name:        "UpdatedAt",
			Description: "",
			Type:        GQLUInt64Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexMeta, ok := p.Source.(*pb.IndexMeta)
				if !ok {
					return nil, nil
				}

				return NewUInt64FromUInt64(indexMeta.UpdatedAt), nil
			},
		},

		"indexDocumentList": &graphql.Field{
			Name:        "IndexDocumentList",
			Description: "",
			Type:        graphql.NewList(GQLIndexDocumentType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexMeta, ok := p.Source.(*pb.IndexMeta)
				if !ok {
					return nil, nil
				}

				return indexMeta.IndexDocumentList, nil
			},
		},
	},
})
