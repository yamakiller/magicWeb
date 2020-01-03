package valid

import (
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

var (
	defaultFilter = [...]string{"admin", "master", "super"}
)

//UserNameValid doc
//Summary validator username is valid
//Method UserNameValid
//Return (bool)
func UserNameValid(fl validator.FieldLevel) bool {
	s := strings.ToLower(fl.Field().String())
	strings.Index(strings.ToLower(s), "")
	for _, v := range defaultFilter {
		if strings.Index(s, v) >= 0 {
			return false
		}
	}

	return true
}

//UserStateValid doc
//Summary validator user state is valid
//Method UserStateValid
//Return (bool)
func UserStateValid(fl validator.FieldLevel) bool {
	s := strings.ToLower(fl.Field().String())
	if s != "lock" && s != "unlock" {
		return false
	}

	return true
}
