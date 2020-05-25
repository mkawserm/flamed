package types

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/pb"
)

var IndexFieldTypeEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "IndexFieldType",
	Description: "`IndexFieldType` defines enum for `IndexField`",
	Values: graphql.EnumValueConfigMap{
		"TEXT": &graphql.EnumValueConfig{
			Value:       0,
			Description: "TEXT",
		},

		"NUMERIC": &graphql.EnumValueConfig{
			Value:       1,
			Description: "NUMERIC",
		},

		"BOOLEAN": &graphql.EnumValueConfig{
			Value:       2,
			Description: "BOOLEAN",
		},

		"GEO_POINT": &graphql.EnumValueConfig{
			Value:       3,
			Description: "GEO_POINT",
		},

		"DATE_TIME": &graphql.EnumValueConfig{
			Value:       4,
			Description: "DATE_TIME",
		},
	},
})

var IndexFieldType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "IndexField",
	Description: "`IndexField`",
	Fields: graphql.Fields{
		"indexFieldType": &graphql.Field{
			Name:        "IndexFieldType",
			Description: "",
			Type:        IndexFieldTypeEnum,
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
	},
})

var IndexDocumentType = graphql.NewObject(graphql.ObjectConfig{
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

		"default": &graphql.Field{
			Name:        "Default",
			Description: "",
			Type:        graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				indexDocument, ok := p.Source.(*pb.IndexDocument)
				if !ok {
					return nil, nil
				}

				return indexDocument.Default, nil
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
			Type:        graphql.NewList(IndexFieldType),
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
			Type:        UInt64Type,
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
			Type:        UInt64Type,
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
			Type:        UInt64Type,
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
			Type:        graphql.NewList(IndexDocumentType),
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
