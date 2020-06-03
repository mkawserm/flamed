package types

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/crypto"
	"github.com/mkawserm/flamed/pkg/iface"
)

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
