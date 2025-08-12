package utils

import (
	"fmt"
	"strings"

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
		return "Este campo é obrigatório"
	case "email":
		return "Formato de e-mail inválido"
	case "min":
		return fmt.Sprintf("Deve ter no mínimo %s caracteres", fe.Param())
	case "max":
		return fmt.Sprintf("Deve ter no máximo %s caracteres", fe.Param())
	default:
		return fmt.Sprintf("Falha na validação: %s", fe.Tag())
	}
}

func toSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}
