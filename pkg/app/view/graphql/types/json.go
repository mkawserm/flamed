package types

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

type JSON struct {
	value map[string]interface{}
}

func (j *JSON) Value() map[string]interface{} {
	return j.value
}

func NewJSONFromString(v string) *JSON {
	dataMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(v), &dataMap)
	if err != nil {
		return &JSON{value: make(map[string]interface{})}
	}

	return &JSON{value: dataMap}
}

func NewJSONFromMap(v map[string]interface{}) *JSON {
	return &JSON{value: v}
}

var GQLJSONType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "JSON",
	Description: "The `JSON` scalar type represents arbitrary JSON value.",

	// Serialize serializes `UInt64` to uint64.
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case JSON:
			return value.Value()
		case *JSON:
			v := *value
			return v.Value()
		default:
			return nil
		}
	},

	// ParseValue parses GraphQL variables from `string` to `JSON`.
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			return NewJSONFromString(value)
		case *string:
			return NewJSONFromString(*value)
		default:
			return nil
		}
	},
	// ParseLiteral parses GraphQL AST value to `JSON`.
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return NewJSONFromString(valueAST.Value)
		default:
			return nil
		}
	},
})
