package valid

import (
	"reflect"

	"gopkg.in/go-playground/validator.v8"
)

//OrderDirectValid doc
//Summary validator int is order direct is valid
//Method OrderDirectValid
//Return (bool)
func OrderDirectValid(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if s, ok := field.Interface().(int); ok {
		if s > 1 || s < 0 {
			return false
		}
		return true
	}

	return false
}
