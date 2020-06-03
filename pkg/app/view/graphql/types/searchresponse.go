package types

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/iface"
)

var GQLDocument = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Document",
	Description: "Show document",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Name:        "id",
			Description: "ID",
			Type:        graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				dm, _ := p.Source.(iface.IDocument)
				return dm.ID(), nil
			},
		},

		"score": &graphql.Field{
			Name:        "score",
			Description: "ID",
			Type:        graphql.NewNonNull(graphql.Float),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				dm, _ := p.Source.(iface.IDocument)
				return dm.Score(), nil
			},
		},

		"index": &graphql.Field{
			Name:        "index",
			Description: "Index",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				dm, _ := p.Source.(iface.IDocument)
				return dm.Index(), nil
			},
		},
	},
})

var GQLSearchResponse = graphql.NewObject(graphql.ObjectConfig{
	Name:        "SearchResponse",
	Description: "Search response",
	Fields: graphql.Fields{
		"hits": &graphql.Field{
			Name:        "hits",
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

		"total": &graphql.Field{
			Name:        "total",
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
			Name:        "maxScore",
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
			Name:        "took",
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
			Name:        "searchTime",
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
