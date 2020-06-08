package graphql

import (
	goGraphQL "github.com/graphql-go/graphql"
	flamedContext "github.com/mkawserm/flamed/pkg/context"
)

type GQLHandler func(flamedContext *flamedContext.FlamedContext) *goGraphQL.Field

type GraphQL struct {
	mQueryFields        goGraphQL.Fields
	mMutationFields     goGraphQL.Fields
	mSubscriptionFields goGraphQL.Fields

	mFlamedContext *flamedContext.FlamedContext

	mSchema          goGraphQL.Schema
	mSchemaBuildFlag bool
}

func (g *GraphQL) AddQueryField(name string, handler GQLHandler) {
	g.mQueryFields[name] = handler(g.mFlamedContext)
}

func (g *GraphQL) AddMutationField(name string, handler GQLHandler) {
	g.mMutationFields[name] = handler(g.mFlamedContext)
}

func (g *GraphQL) AddSubscriptionField(name string, handler GQLHandler) {
	g.mSubscriptionFields[name] = handler(g.mFlamedContext)
}

func (g *GraphQL) IsSchemaAvailable() bool {
	return g.mSchemaBuildFlag
}

func (g *GraphQL) GetSchema() goGraphQL.Schema {
	return g.mSchema
}

func (g *GraphQL) BuildSchema() (goGraphQL.Schema, error) {
	if g.mSchemaBuildFlag {
		return g.mSchema, nil
	}

	g.register()
	/* build schema */
	query := goGraphQL.ObjectConfig{
		Name:        "Query",
		Fields:      g.mQueryFields,
		Description: "All available GraphQL queries",
	}

	mutation := goGraphQL.ObjectConfig{
		Name:        "Mutation",
		Fields:      g.mMutationFields,
		Description: "All available GraphQL mutations",
	}

	subscription := goGraphQL.ObjectConfig{
		Name:        "Subscription",
		Description: "All available GraphQL subscriptions",
		Fields:      g.mSubscriptionFields,
	}

	schemaConfig := goGraphQL.SchemaConfig{
		Query: goGraphQL.NewObject(query),
	}

	if len(g.mMutationFields) != 0 {
		schemaConfig.Mutation = goGraphQL.NewObject(mutation)
	}

	if len(g.mSubscriptionFields) != 0 {
		schemaConfig.Subscription = goGraphQL.NewObject(subscription)
	}

	var err error
	g.mSchema, err = goGraphQL.NewSchema(schemaConfig)
	if err != nil {
		panic(err)
	}
	g.mSchemaBuildFlag = true
	return g.mSchema, err
}

func NewGraphQL(flamedContext *flamedContext.FlamedContext) *GraphQL {
	return &GraphQL{
		mFlamedContext:      flamedContext,
		mQueryFields:        map[string]*goGraphQL.Field{},
		mMutationFields:     map[string]*goGraphQL.Field{},
		mSubscriptionFields: map[string]*goGraphQL.Field{},
	}
}
