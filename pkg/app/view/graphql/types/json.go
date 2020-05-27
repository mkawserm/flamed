package types

//
//import (
//	"encoding/json"
//	"github.com/graphql-go/graphql"
//	"github.com/graphql-go/graphql/language/ast"
//)
//
//type JSON struct {
//	value map[string]interface{}
//}
//
//func (j *JSON) Value() map[string]interface{} {
//	return j.value
//}
//
//func NewJSONFromString(v string) *JSON {
//	dataMap := make(map[string]interface{})
//	err := json.Unmarshal([]byte(v), &dataMap)
//	if err != nil {
//		return &JSON{value: make(map[string]interface{})}
//	}
//
//	return &JSON{value: dataMap}
//}
//
//func NewJSONFromMap(v map[string]interface{}) *JSON {
//	return &JSON{value: v}
//}
//
//var GQLJSONType = graphql.NewScalar(graphql.ScalarConfig{
//	Name:        "JSON",
//	Description: "The `JSON` scalar type represents arbitrary JSON value.",
//
//	// Serialize serializes `JSON` to JSON MAP.
//	Serialize: func(value interface{}) interface{} {
//		switch value := value.(type) {
//		case JSON:
//			return value.Value()
//		case *JSON:
//			v := *value
//			return v.Value()
//		default:
//			return nil
//		}
//	},
//
//	// ParseValue parses GraphQL variables from `string` to `JSON`.
//	ParseValue: func(value interface{}) interface{} {
//		switch value := value.(type) {
//		case string:
//			return NewJSONFromString(value)
//		case *string:
//			return NewJSONFromString(*value)
//		default:
//			return nil
//		}
//	},
//	// ParseLiteral parses GraphQL AST value to `JSON`.
//	ParseLiteral: func(valueAST ast.Value) interface{} {
//		switch valueAST := valueAST.(type) {
//		case *ast.StringValue:
//			return NewJSONFromString(valueAST.Value)
//		default:
//			return nil
//		}
//	},
//})

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
)

func parseLiteral(astValue ast.Value) interface{} {
	kind := astValue.GetKind()

	switch kind {
	case kinds.StringValue:
		return astValue.GetValue()
	case kinds.BooleanValue:
		return astValue.GetValue()
	case kinds.IntValue:
		return astValue.GetValue()
	case kinds.FloatValue:
		return astValue.GetValue()
	case kinds.ObjectValue:
		obj := make(map[string]interface{})
		for _, v := range astValue.GetValue().([]*ast.ObjectField) {
			obj[v.Name.Value] = parseLiteral(v.Value)
		}
		return obj
	case kinds.ListValue:
		list := make([]interface{}, 0)
		for _, v := range astValue.GetValue().([]ast.Value) {
			list = append(list, parseLiteral(v))
		}
		return list
	default:
		return nil
	}
}

// JSON json type
var GQLJSONType = graphql.NewScalar(
	graphql.ScalarConfig{
		Name:        "JSON",
		Description: "The `JSON` scalar type represents JSON values as specified by [ECMA-404](http://www.ecma-international.org/publications/files/ECMA-ST/ECMA-404.pdf)",
		Serialize: func(value interface{}) interface{} {
			return value
		},
		ParseValue: func(value interface{}) interface{} {
			return value
		},
		ParseLiteral: parseLiteral,
	},
)
