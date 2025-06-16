package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(regex, email)
	if !matched {
		return false
	}
	return true
}
