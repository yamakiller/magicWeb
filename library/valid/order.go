package valid

import (
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

//OrderDirectValid doc
//Summary validator int is order direct is valid
//Method OrderDirectValid
//Return (bool)
func OrderDirectValid(fl validator.FieldLevel) bool {
	value := strings.ToLower(fl.Field().String())
	if value == "asc" || value == "desc" {
		return true
	}
	return false
}
