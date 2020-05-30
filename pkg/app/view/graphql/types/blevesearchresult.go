package types

import (
	bleveSearch "github.com/blevesearch/bleve"
	"github.com/graphql-go/graphql"
)

var GQLBleveSearchResult = graphql.NewObject(graphql.ObjectConfig{
	Name:        "BleveSearchResult",
	Description: "Bleve search result",
	Fields: graphql.Fields{
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
	},
})
