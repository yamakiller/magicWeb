package valid

import (
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v8"
)

var (
	defaultFilter = [...]string{"admin", "master", "super"}
)

//UserNameValid doc
//Summary validator username is valid
//Method UserNameValid
//Return (bool)
func UserNameValid(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if s, ok := field.Interface().(string); ok {
		s = strings.ToLower(s)
		strings.Index(strings.ToLower(s), "")
		for _, v := range defaultFilter {
			if strings.Index(s, v) >= 0 {
				return false
			}
		}

		return true
	}

	return false
}

//UserStateValid doc
//Summary validator user state is valid
//Method UserStateValid
//Return (bool)
func UserStateValid(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if s, ok := field.Interface().(string); ok {
		v := strings.ToLower(s)
		if v != "lock" && v != "unlock" {
			return false
		}

		return true
	}

	return false
}

//UserFeatureValid doc
//Summary validator user feature is valid
//Method UserFeatureValid
//Return (bool)
func UserFeatureValid(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if s, ok := field.Interface().(string); ok {
		v := strings.ToLower(s)
		if v != "admin" && v != "nomal" {
			return false
		}

		return true
	}

	return false
}
