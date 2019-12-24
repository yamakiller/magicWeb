package valid

import (
	"reflect"

	"gopkg.in/go-playground/validator.v8"
)

//CaptchaImageWidthValid doc
//Summary validator captcha image width is valid
//Method CaptchaImageWidthValid
//Return (bool)
func CaptchaImageWidthValid(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if s, ok := field.Interface().(int); ok {
		if s > 540 {
			return false
		}

		return true
	}

	return false
}

//CaptchaImageHeightValid doc
//Summary validator captcha image height is valid
//Method CaptchaImageHeightValid
//Return (bool)
func CaptchaImageHeightValid(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if s, ok := field.Interface().(int); ok {
		if s > 240 {
			return false
		}

		return true
	}

	return false
}

//CaptchaImageModeValid doc
//Summary validator captcha image mode is valid
//Method CaptchaImageHeightValid
//Return (bool)
func CaptchaImageModeValid(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if s, ok := field.Interface().(int); ok {
		if s < 0 || s > 4 {
			return false
		}

		return true
	}

	return false
}
