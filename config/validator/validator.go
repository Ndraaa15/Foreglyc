package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func New() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("password", PasswordValidation)
	return validate
}

func PasswordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Regex: At least one lowercase, one uppercase, and one digit
	var passwordRegex = regexp.MustCompile(`^[a-zA-Z\d]+$`)

	return passwordRegex.MatchString(password)
}
