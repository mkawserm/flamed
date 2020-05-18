package json

import (
	"reflect"
	"strings"
)

func GetId(document interface{}) string {
	return getId(reflect.TypeOf(document), reflect.ValueOf(document))
}

func getId(t reflect.Type, v reflect.Value) string {
	if t.Kind() == reflect.Ptr {
		return getId(t.Elem(), v.Elem())
	} else if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			tags := strings.Split(f.Tag.Get("json"), ",")
			if len(tags) > 0 {
				name := strings.TrimSpace(tags[0])
				if name == "id" {
					v2 := v.FieldByName(f.Name)
					if v2.Kind() == reflect.String {
						return v2.String()
					}
				}
			}
		}
	}

	return ""
}
