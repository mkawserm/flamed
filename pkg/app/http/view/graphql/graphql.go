package graphql

import (
	goContext "context"
	"encoding/json"
	goGraphQL "github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	utility2 "github.com/mkawserm/flamed/pkg/app/utility"
	flamedContext "github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/logger"
	"github.com/mkawserm/flamed/pkg/utility"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strings"
)

type GraphQL struct {
	mFlamedContext *flamedContext.FlamedContext
	mSchema        goGraphQL.Schema
}

func (v *GraphQL) GetHTTPHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger.L("app::http::view::graphql").Debug("processing graphql request")

		if !utility2.GetCORSOptions().CROSCheckAllowNext(writer, request) {
			return
		}

		writer.Header().Add("Content-Type", "application/json; charset=utf-8")

		var result *goGraphQL.Result

		if request.Method != http.MethodPost {
			writer.WriteHeader(http.StatusBadRequest)
			result = &goGraphQL.Result{
				Data: nil,
				Errors: []gqlerrors.FormattedError{
					gqlerrors.NewFormattedError("request method must be `POST`"),
				},
				Extensions: nil,
			}
			_ = json.NewEncoder(writer).Encode(result)
			//rJSON, _ := json.Marshal(result)
			//_, _ = writer.Write(rJSON)
			return
		}

		bodyBytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			result = &goGraphQL.Result{
				Data:       nil,
				Errors:     []gqlerrors.FormattedError{gqlerrors.NewFormattedError(err.Error())},
				Extensions: nil,
			}
			_ = json.NewEncoder(writer).Encode(result)
			//rJSON, _ := json.Marshal(result)
			//_, _ = writer.Write(rJSON)
			return
		}

		logger.L("app::http::view::graphql").Debug("graphql request body received")

		var fields []zap.Field
		header := make(map[string][]string)
		for k, v := range request.Header {
			header[k] = v[:]

			// skip authorization header from debug log
			if strings.EqualFold(k, "Authorization") {
				fields = append(fields, zap.String(k, "****"))
				continue
			}
			fields = append(fields, zap.String(k, strings.Join(v, ",")))
		}

		logger.L("app::http::view::graphql").Debug("graphql request header", fields...)

		ro := utility.ParseGraphQLQuery(bodyBytes)

		var params goGraphQL.Params

		var graphQLContext = &flamedContext.AuthContext{}
		graphQLContext.Protocol = "HTTP"
		graphQLContext.KVPair = header
		graphQLContext.Host = request.Host
		graphQLContext.URL = request.URL.String()
		graphQLContext.RemoteAddr = request.RemoteAddr
		graphQLContext.RequestURI = request.RequestURI

		params = goGraphQL.Params{
			Schema:         v.mSchema,
			RequestString:  ro.Query,
			VariableValues: ro.Variables,
			OperationName:  ro.OperationName,
			Context: goContext.WithValue(goContext.Background(),
				"GraphQLContext",
				graphQLContext),
		}

		result = goGraphQL.Do(params)
		if len(result.Errors) > 0 {
			writer.WriteHeader(http.StatusBadRequest)
		} else {
			writer.WriteHeader(http.StatusOK)
		}
		_ = json.NewEncoder(writer).Encode(result)
		//rJSON, _ := json.Marshal(result)
		//_, _ = writer.Write(rJSON)

		logger.L("app::http::view::graphql").Debug("graphql request response served")
	}
}

func NewGraphQLView(flamedContext *flamedContext.FlamedContext, schema goGraphQL.Schema) *GraphQL {
	return &GraphQL{
		mSchema:        schema,
		mFlamedContext: flamedContext,
	}
}
