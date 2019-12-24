package valid

import (
	"reflect"
	"time"

	"gopkg.in/go-playground/validator.v8"
)

//DateTimeValid doc
//Summary validator string is datatime is valid
//Method DateTimeValid
//Return (bool)
func DateTimeValid(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if s, ok := field.Interface().(string); ok {
		if _, err := time.Parse("2006-01-02 15:04:05", s); err != nil {
			return false
		}
		return true
	}

	return false
}
