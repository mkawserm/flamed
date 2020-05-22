package graphql

import (
	goContext "context"

	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/context"
	"io/ioutil"
	"net/http"
)

type GQLHandler func(flamedContext *context.FlamedContext) *graphql.Field

type View struct {
	mQueryFields    graphql.Fields
	mMutationFields graphql.Fields

	mFlamedContext *context.FlamedContext
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
		writer.Header().Add("Content-Type", "application/json; charset=utf-8")

		var result *graphql.Result

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

		header := make(http.Header)
		for k, v := range request.Header {
			header[k] = v
		}
		ro := ParseGraphQLQuery(bodyBytes)

		var params graphql.Params
		params = graphql.Params{
			Schema:         schema,
			RequestString:  ro.Query,
			VariableValues: ro.Variables,
			OperationName:  ro.OperationName,
			Context: goContext.
				WithValue(goContext.Background(),
					"HTTPHeader",
					header),
		}

		result = graphql.Do(params)
		if len(result.Errors) > 0 {
			writer.WriteHeader(http.StatusBadRequest)
		} else {
			writer.WriteHeader(http.StatusOK)
		}

		rJSON, _ := json.Marshal(result)
		_, _ = writer.Write(rJSON)
	}
}

func NewView(flamedContext *context.FlamedContext) *View {
	return &View{
		mFlamedContext:  flamedContext,
		mQueryFields:    map[string]*graphql.Field{},
		mMutationFields: map[string]*graphql.Field{},
	}
}
