package paramverify

import (
	"reflect"

	"gopkg.in/go-playground/validator.v8"
)

// NameValid 自定义验证器
func NameValid(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if s, ok := field.Interface().(string); ok {
		if s == "admin" {
			return true
		}
	}
	return false
}
