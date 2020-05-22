package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/mkawserm/flamed/pkg/context"
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

	return func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello world"))
	}
}

func NewView(flamedContext *context.FlamedContext) *View {
	return &View{mFlamedContext: flamedContext}
}
