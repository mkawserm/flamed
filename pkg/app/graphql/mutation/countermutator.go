package mutation

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	flamedContext "github.com/mkawserm/flamed/pkg/context"
	"sync/atomic"
)

var GQLCounterMutatorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "CounterMutator",
	Description: "Counter mutator provides an in-memory" +
		" unsigned 64 bit counter which can be incremented or decremented by one",
	Fields: graphql.Fields{
		"increment": &graphql.Field{
			Name:        "Increment",
			Type:        kind.GQLUInt64Type,
			Description: "Increment counter by one",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i, ok := p.Source.(*uint64)
				if !ok {
					return nil, nil
				}
				if *i < 18446744073709551615 {
					var m uint64 = 1
					atomic.AddUint64(i, m)
				}
				return kind.NewUInt64FromUInt64(*i), nil
			},
		},

		"decrement": &graphql.Field{
			Name:        "Decrement",
			Type:        kind.GQLUInt64Type,
			Description: "Decrement counter by one",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i, ok := p.Source.(*uint64)
				if !ok {
					return nil, nil
				}

				if *i > 0 {
					var m uint64 = 1
					atomic.AddUint64(i, -m)
				}

				return kind.NewUInt64FromUInt64(*i), nil
			},
		},
	},
})

func CounterMutator(_ *flamedContext.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Name:        "CounterMutator",
		Type:        GQLCounterMutatorType,
		Description: "Counter mutator provides an in-memory unsigned 64 bit counter",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return &i, nil
		},
	}
}
