package types

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/crypto"
	"github.com/mkawserm/flamed/pkg/iface"
)

var GQLTerm = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Term",
	Description: "Term",
	Fields: graphql.Fields{
		"term": &graphql.Field{
			Name:        "Term",
			Description: "Term",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if term, ok := p.Source.(iface.ITerm); ok {
					return term.Term(), nil
				}
				return nil, nil
			},
		},

		"count": &graphql.Field{
			Name:        "Count",
			Description: "Count",
			Type:        graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if term, ok := p.Source.(iface.ITerm); ok {
					return term.Count(), nil
				}
				return nil, nil
			},
		},
	},
})

var GQLDateRangeFacet = graphql.NewObject(graphql.ObjectConfig{
	Name:        "DateRangeFacet",
	Description: "Date range facet",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Name:        "Name",
			Description: "Name",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if dateRangeFacet, ok := p.Source.(iface.IDateRangeFacet); ok {
					return dateRangeFacet.Name(), nil
				}
				return nil, nil
			},
		},

		"start": &graphql.Field{
			Name:        "Start",
			Description: "Start",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if dateRangeFacet, ok := p.Source.(iface.IDateRangeFacet); ok {
					return dateRangeFacet.Start(), nil
				}
				return nil, nil
			},
		},
		"end": &graphql.Field{
			Name:        "End",
			Description: "End",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if dateRangeFacet, ok := p.Source.(iface.IDateRangeFacet); ok {
					return dateRangeFacet.End(), nil
				}
				return nil, nil
			},
		},
		"count": &graphql.Field{
			Name:        "Count",
			Description: "Count",
			Type:        graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if dateRangeFacet, ok := p.Source.(iface.IDateRangeFacet); ok {
					return dateRangeFacet.Count(), nil
				}
				return nil, nil
			},
		},
	},
})

var GQLNumericRangeFacet = graphql.NewObject(graphql.ObjectConfig{
	Name:        "NumericRangeFacet",
	Description: "Numeric range facet",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Name:        "Name",
			Description: "Name",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if numericRangeFacet, ok := p.Source.(iface.INumericRangeFacet); ok {
					return numericRangeFacet.Name(), nil
				}
				return nil, nil
			},
		},

		"min": &graphql.Field{
			Name:        "Min",
			Description: "Min",
			Type:        graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if numericRangeFacet, ok := p.Source.(iface.INumericRangeFacet); ok {
					return numericRangeFacet.Min(), nil
				}
				return nil, nil
			},
		},
		"max": &graphql.Field{
			Name:        "Max",
			Description: "Max",
			Type:        graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if numericRangeFacet, ok := p.Source.(iface.INumericRangeFacet); ok {
					return numericRangeFacet.Max(), nil
				}
				return nil, nil
			},
		},
		"count": &graphql.Field{
			Name:        "Count",
			Description: "Count",
			Type:        graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if numericRangeFacet, ok := p.Source.(iface.INumericRangeFacet); ok {
					return numericRangeFacet.Count(), nil
				}
				return nil, nil
			},
		},
	},
})

var GQLDocument = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Document",
	Description: "Show document",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Name:        "ID",
			Description: "ID",
			Type:        graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				dm, _ := p.Source.(iface.IDocument)
				return dm.ID(), nil
			},
		},

		"score": &graphql.Field{
			Name:        "Score",
			Description: "Score",
			Type:        graphql.NewNonNull(graphql.Float),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				dm, _ := p.Source.(iface.IDocument)
				return dm.Score(), nil
			},
		},

		"index": &graphql.Field{
			Name:        "Index",
			Description: "Index",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				dm, _ := p.Source.(iface.IDocument)
				return dm.Index(), nil
			},
		},

		"namespace": &graphql.Field{
			Name:        "Namespace",
			Description: "Namespace",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				dm, _ := p.Source.(iface.IDocument)
				return string(crypto.GetNamespaceFromStateAddressHexString(dm.ID())), nil
			},
		},
	},
})

var GQLFacet = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Facet",
	Description: "Facet",
	Fields: &graphql.Fields{
		"name": &graphql.Field{
			Name: "Name",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if facet, ok := p.Source.(iface.IFacet); ok {
					return facet.Name(), nil
				}
				return nil, nil
			},
		},

		"field": &graphql.Field{
			Name: "Field",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if facet, ok := p.Source.(iface.IFacet); ok {
					return facet.Field(), nil
				}
				return nil, nil
			},
		},

		"total": &graphql.Field{
			Name: "Total",
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if facet, ok := p.Source.(iface.IFacet); ok {
					return facet.Total(), nil
				}
				return nil, nil
			},
		},

		"missing": &graphql.Field{
			Name: "Missing",
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if facet, ok := p.Source.(iface.IFacet); ok {
					return facet.Missing(), nil
				}
				return nil, nil
			},
		},

		"other": &graphql.Field{
			Name: "Other",
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if facet, ok := p.Source.(iface.IFacet); ok {
					return facet.Other(), nil
				}
				return nil, nil
			},
		},

		"terms": &graphql.Field{
			Name: "Terms",
			Type: graphql.NewList(GQLTerm),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if facet, ok := p.Source.(iface.IFacet); ok {
					return facet.Terms(), nil
				}
				return nil, nil
			},
		},

		"numericRanges": &graphql.Field{
			Name: "NumericRanges",
			Type: graphql.NewList(GQLNumericRangeFacet),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if facet, ok := p.Source.(iface.IFacet); ok {
					return facet.NumericRanges(), nil
				}
				return nil, nil
			},
		},

		"dateRanges": &graphql.Field{
			Name: "DateRanges",
			Type: graphql.NewList(GQLDateRangeFacet),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if facet, ok := p.Source.(iface.IFacet); ok {
					return facet.DateRanges(), nil
				}
				return nil, nil
			},
		},
	},
})

var GQLSearchResponse = graphql.NewObject(graphql.ObjectConfig{
	Name:        "SearchResponse",
	Description: "Search response",
	Fields: graphql.Fields{
		"hits": &graphql.Field{
			Name:        "Hits",
			Description: "Hits",
			Type:        graphql.NewList(GQLDocument),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				sr, ok := p.Source.(iface.ISearchResponse)
				if !ok {
					return nil, nil
				}

				return sr.Hits(), nil
			},
		},

		"facets": &graphql.Field{
			Name:        "Facets",
			Description: "Facets",
			Type:        graphql.NewList(GQLFacet),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				sr, ok := p.Source.(iface.ISearchResponse)
				if !ok {
					return nil, nil
				}
				return sr.Facets(), nil
			},
		},

		"total": &graphql.Field{
			Name:        "Total",
			Description: "Total result",
			Type:        GQLUInt64Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				sr, ok := p.Source.(iface.ISearchResponse)
				if !ok {
					return nil, nil
				}

				return NewUInt64FromUInt64(sr.Total()), nil
			},
		},

		"maxScore": &graphql.Field{
			Name:        "MaxScore",
			Description: "Maximum score",
			Type:        graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				sr, ok := p.Source.(iface.ISearchResponse)
				if !ok {
					return nil, nil
				}

				return sr.MaxScore(), nil
			},
		},

		"took": &graphql.Field{
			Name:        "Took",
			Description: "Search time",
			Type:        GQLUInt64Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				sr, ok := p.Source.(iface.ISearchResponse)
				if !ok {
					return nil, nil
				}

				return NewUInt64FromUInt64(uint64(sr.Took())), nil
			},
		},

		"searchTime": &graphql.Field{
			Name:        "SearchTime",
			Description: "Search time in string",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				sr, ok := p.Source.(iface.ISearchResponse)
				if !ok {
					return nil, nil
				}

				return sr.Took().String(), nil
			},
		},
	},
})
