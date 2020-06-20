package query

import (
	"github.com/graphql-go/graphql"
	utility2 "github.com/mkawserm/flamed/pkg/app/utility"
	"github.com/mkawserm/flamed/pkg/context"
)

var ServerStatusType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ServerStatus",
	Description: "`ServerStatus` provides information about service availability",
	Fields: graphql.Fields{
		"httpServer": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Is HTTP server available?",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*utility2.ServerStatus); ok {
					return val.HTTPServer(), nil
				}
				return nil, nil
			},
		},

		"grpcServer": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Is GRPC server available?",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*utility2.ServerStatus); ok {
					return val.GRPCServer(), nil
				}
				return nil, nil
			},
		},

		"raftServer": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Is RAFT server available?",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(*utility2.ServerStatus); ok {
					return val.RAFTServer(), nil
				}
				return nil, nil
			},
		},
	},
})

func ServerStatus(_ *context.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Name: "ServerStatus",
		Type: ServerStatusType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return utility2.GetServerStatus(), nil
		},
		Description: "GlobalOperation server availability information",
	}
}
