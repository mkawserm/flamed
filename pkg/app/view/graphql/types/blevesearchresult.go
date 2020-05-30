package types

import (
	bleveSearch "github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search"
	"github.com/graphql-go/graphql"
)

var GQLDocumentMatch = graphql.NewObject(graphql.ObjectConfig{
	Name:        "DocumentMatch",
	Description: "Show matched document",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Name:        "id",
			Description: "ID",
			Type:        graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				dm, _ := p.Source.(*search.DocumentMatch)
				return dm.ID, nil
			},
		},

		"score": &graphql.Field{
			Name:        "score",
			Description: "ID",
			Type:        graphql.NewNonNull(graphql.Float),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				dm, _ := p.Source.(*search.DocumentMatch)
				return dm.Score, nil
			},
		},

		"index": &graphql.Field{
			Name:        "index",
			Description: "Index",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				dm, _ := p.Source.(*search.DocumentMatch)
				return dm.Index, nil
			},
		},
	},
})

var GQLBleveSearchResult = graphql.NewObject(graphql.ObjectConfig{
	Name:        "BleveSearchResult",
	Description: "Bleve search result",
	Fields: graphql.Fields{
		"hits": &graphql.Field{
			Name:        "hits",
			Description: "Hits",
			Type:        graphql.NewList(GQLDocumentMatch),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				sr, ok := p.Source.(*bleveSearch.SearchResult)
				if !ok {
					return nil, nil
				}

				return sr.Hits, nil
			},
		},

		"total": &graphql.Field{
			Name:        "total",
			Description: "Total result",
			Type:        GQLUInt64Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				sr, ok := p.Source.(*bleveSearch.SearchResult)
				if !ok {
					return nil, nil
				}

				return NewUInt64FromUInt64(sr.Total), nil
			},
		},

		"maxScore": &graphql.Field{
			Name:        "maxScore",
			Description: "Maximum score",
			Type:        graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				sr, ok := p.Source.(*bleveSearch.SearchResult)
				if !ok {
					return nil, nil
				}

				return sr.MaxScore, nil
			},
		},

		"took": &graphql.Field{
			Name:        "took",
			Description: "Search time",
			Type:        GQLUInt64Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				sr, ok := p.Source.(*bleveSearch.SearchResult)
				if !ok {
					return nil, nil
				}

				return NewUInt64FromUInt64(uint64(sr.Took)), nil
			},
		},

		"searchTime": &graphql.Field{
			Name:        "searchTime",
			Description: "Search time in string",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				sr, ok := p.Source.(*bleveSearch.SearchResult)
				if !ok {
					return nil, nil
				}

				return sr.Took.String(), nil
			},
		},
	},
})
