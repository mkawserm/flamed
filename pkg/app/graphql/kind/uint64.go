package kind

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"strconv"
)

type UInt64 struct {
	value uint64
}

func (u *UInt64) Value() uint64 {
	return u.value
}

func NewUInt64FromString(v string) *UInt64 {
	u, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return &UInt64{0}
	}
	return &UInt64{value: u}
}

func NewUInt64FromInt(v int) *UInt64 {
	return &UInt64{value: uint64(v)}
}

func NewUInt64FromUInt64(v uint64) *UInt64 {
	return &UInt64{value: v}
}

var GQLUInt64Type = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "UInt64",
	Description: "The `UInt64` scalar type represents an unsigned 64 bit integer.",
	// Serialize serializes `UInt64` to uint64.
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case UInt64:
			return value.Value()
		case *UInt64:
			v := *value
			return v.Value()
		default:
			return nil
		}
	},
	// ParseValue parses GraphQL variables from `string` to `UInt64`.
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			return NewUInt64FromString(value)
		case *string:
			return NewUInt64FromString(*value)
		case int:
			return NewUInt64FromInt(value)
		case *int:
			return NewUInt64FromInt(*value)

		default:
			return nil
		}
	},
	// ParseLiteral parses GraphQL AST value to `UInt64`.
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return NewUInt64FromString(valueAST.Value)
		case *ast.IntValue:
			return NewUInt64FromString(valueAST.Value)
		default:
			return nil
		}
	},
})
