package graphql

import (
	"context"
	"encoding/json"
	goGraphQL "github.com/graphql-go/graphql"
	flamedContext "github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/logger"
	"github.com/mkawserm/flamed/pkg/utility"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

type Service struct {
	mFlamedContext *flamedContext.FlamedContext
	mSchema        goGraphQL.Schema

	UnimplementedGraphQLRPCServer
}

func (s *Service) RegisterGRPCService(server *grpc.Server) {
	RegisterGraphQLRPCServer(server, s)
}

func (s *Service) GetGraphQLResponse(ctx context.Context, request *GraphQLRequest) (*GraphQLResponse, error) {
	logger.L("app::grpc::service::graphql").Debug("graphql request processing started")

	var fields []zap.Field
	header := make(map[string][]string)
	if headers, ok := metadata.FromIncomingContext(ctx); ok {
		for k, v := range headers {
			header[k] = v[:]

			// skip authorization header from debug log
			if strings.EqualFold(k, "Authorization") {
				fields = append(fields, zap.String(k, "****"))
				continue
			}
			fields = append(fields, zap.String(k, strings.Join(v, ",")))
		}
	}

	logger.L("app::grpc::service::graphql").Debug("graphql request header", fields...)

	ro := utility.ParseGraphQLQuery(request.Payload)

	var params goGraphQL.Params

	var graphQLContext = &flamedContext.GraphQLContext{}
	graphQLContext.Protocol = "GRPC"
	graphQLContext.Header = header

	params = goGraphQL.Params{
		Schema:         s.mSchema,
		RequestString:  ro.Query,
		VariableValues: ro.Variables,
		OperationName:  ro.OperationName,
		Context: context.WithValue(ctx,
			"GraphQLContext",
			graphQLContext),
	}

	result := goGraphQL.Do(params)
	rJSON, _ := json.Marshal(result)
	response := &GraphQLResponse{Payload: rJSON}

	logger.L("app::grpc::service::graphql").Debug("graphql request response served")

	return response, nil
}

func NewGraphQLService(flamedContext *flamedContext.FlamedContext,
	schema goGraphQL.Schema) *Service {
	return &Service{
		mFlamedContext: flamedContext,
		mSchema:        schema,
	}
}
