package query

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/context"
)

type ServiceStatus struct {
	HTTPServer bool
	GRPCServer bool
	RAFTServer bool
}

var ServiceStatusType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ServiceStatus",
	Description: "`ServiceStatus` provides information about service availability",
	Fields: graphql.Fields{
		"httpServer": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Is HTTP server available?",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(ServiceStatus); ok {
					return val.HTTPServer, nil
				}
				return nil, nil
			},
		},

		"grpcServer": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Is GRPC server available?",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(ServiceStatus); ok {
					return val.GRPCServer, nil
				}
				return nil, nil
			},
		},

		"raftServer": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Is RAFT server available?",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if val, ok := p.Source.(ServiceStatus); ok {
					return val.RAFTServer, nil
				}
				return nil, nil
			},
		},
	},
})

func IsLive(_ *context.FlamedContext) *graphql.Field {
	return &graphql.Field{
		Name: "IsLive",
		Type: ServiceStatusType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return ServiceStatus{
				HTTPServer: true,
				GRPCServer: false,
				RAFTServer: true,
			}, nil
		},
		Description: "Query about service availability information",
	}
}
