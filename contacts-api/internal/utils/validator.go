package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}) map[string]string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)

	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Field()] = getErrorMessage(err)
	}

	return errors
}

func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Required field"
	case "email":
		return "Invalid e-mail"
	case "min":
		return fmt.Sprintf("Min %s characters", fe.Param())
	case "max":
		return fmt.Sprintf("Max %s characters", fe.Param())
	default:
		return fmt.Sprintf("Fail at validation: %s", fe.Tag())
	}
}
