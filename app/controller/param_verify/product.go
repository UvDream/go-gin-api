package paramverify

import (
	"gopkg.in/go-playground/validator.v9"
)

// NameValid 校验
func NameValid(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	if val == "admin" {
		return true
	}
	return false
}
