package graphql

import (
	goContext "context"

	"github.com/mkawserm/flamed/pkg/logger"
	"go.uber.org/zap"
	"strings"

	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"

	"io/ioutil"
	"net/http"

	flamedContext "github.com/mkawserm/flamed/pkg/context"
)

type GQLHandler func(flamedContext *flamedContext.FlamedContext) *graphql.Field

type View struct {
	mQueryFields    graphql.Fields
	mMutationFields graphql.Fields

	mFlamedContext *flamedContext.FlamedContext
}

func (v *View) AddQueryField(name string, handler GQLHandler) {
	v.mQueryFields[name] = handler(v.mFlamedContext)
}

func (v *View) AddMutationField(name string, handler GQLHandler) {
	v.mMutationFields[name] = handler(v.mFlamedContext)
}

func (v *View) GetHTTPHandler() http.HandlerFunc {
	v.register()

	/* build schema */
	query := graphql.ObjectConfig{
		Name:        "Query",
		Fields:      v.mQueryFields,
		Description: "All available GraphQL queries",
	}

	mutation := graphql.ObjectConfig{
		Name:        "Mutation",
		Fields:      v.mMutationFields,
		Description: "All available GraphQL mutations",
	}

	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(query),
		Mutation: graphql.NewObject(mutation),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		panic(err)
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		logger.L("graphql").Debug("processing graphql request")
		writer.Header().Add("Content-Type", "application/json; charset=utf-8")

		var result *graphql.Result

		if request.Method != http.MethodPost {
			writer.WriteHeader(http.StatusBadRequest)
			result = &graphql.Result{
				Data: nil,
				Errors: []gqlerrors.FormattedError{
					gqlerrors.NewFormattedError("request method must be `POST`"),
				},
				Extensions: nil,
			}
			rJSON, _ := json.Marshal(result)
			_, _ = writer.Write(rJSON)
			return
		}

		bodyBytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			result = &graphql.Result{
				Data:       nil,
				Errors:     []gqlerrors.FormattedError{gqlerrors.NewFormattedError(err.Error())},
				Extensions: nil,
			}
			rJSON, _ := json.Marshal(result)
			_, _ = writer.Write(rJSON)
			return
		}

		logger.L("graphql").Debug("graphql request body",
			zap.ByteString("request", bodyBytes))

		var fields []zap.Field
		header := make(http.Header)
		for k, v := range request.Header {
			header[k] = v
			fields = append(fields, zap.String(k, strings.Join(v, ",")))
		}

		logger.L("graphql").Debug("graphql request header", fields...)

		ro := ParseGraphQLQuery(bodyBytes)

		var params graphql.Params

		var graphQLContext flamedContext.GraphQLContext
		graphQLContext.Header = header
		graphQLContext.Host = request.Host
		graphQLContext.URL = request.URL.String()
		graphQLContext.RemoteAddr = request.RemoteAddr
		graphQLContext.RequestURI = request.RequestURI

		params = graphql.Params{
			Schema:         schema,
			RequestString:  ro.Query,
			VariableValues: ro.Variables,
			OperationName:  ro.OperationName,
			Context: goContext.WithValue(goContext.Background(),
				"GraphQLContext",
				graphQLContext),
		}

		result = graphql.Do(params)
		if len(result.Errors) > 0 {
			writer.WriteHeader(http.StatusBadRequest)
		} else {
			writer.WriteHeader(http.StatusOK)
		}

		rJSON, _ := json.Marshal(result)
		_, _ = writer.Write(rJSON)

		logger.L("graphql").Debug("graphql request response",
			zap.ByteString("response", rJSON))
	}
}

func NewView(flamedContext *flamedContext.FlamedContext) *View {
	return &View{
		mFlamedContext:  flamedContext,
		mQueryFields:    map[string]*graphql.Field{},
		mMutationFields: map[string]*graphql.Field{},
	}
}
