package utils

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateStruct(data interface{}) []string {
	validate = validator.New()
	err := validate.Struct(data)

	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			var message string
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is a required field", err.Field())
			case "min":
				message = fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
			case "max":
				message = fmt.Sprintf("%s cannot be longer than %s characters", err.Field(), err.Param())
			case "gte":
				message = fmt.Sprintf("%s must be greater than or equal to %s", err.Field(), err.Param())
			case "gt":
				message = fmt.Sprintf("%s must be greater than %s", err.Field(), err.Param())
			case "lte":
				message = fmt.Sprintf("%s must be less than or equal to %s", err.Field(), err.Param())
			case "lt":
				message = fmt.Sprintf("%s must be less than %s", err.Field(), err.Param())
			default:
				message = fmt.Sprintf("%s is invalid", err.Field())
			}
			errors = append(errors, message)
		}
		return errors
	}

	return nil
}

func HandleUnmarshalTypeError(err error) []string {
	if unmarshalErr, ok := err.(*json.UnmarshalTypeError); ok {
		return []string{
			fmt.Sprintf("Field '%s' expects a value of type '%s', but got '%s'", unmarshalErr.Field, unmarshalErr.Type, unmarshalErr.Value),
		}
	}
	if syntaxErr, ok := err.(*json.SyntaxError); ok {
		return []string{
			fmt.Sprintf("There is a syntax error at offset %v", syntaxErr.Offset),
		}
	}
	return []string{"Invalid JSON payload"}
}
