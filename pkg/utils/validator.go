package utils

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

// GetValidator Initiatilize validator in singleton way
func GetValidator() *validator.Validate {

	if validate == nil {
		validate = validator.New()
	}

	return validate
}
