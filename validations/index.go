package validations

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// Helper function to format validation errors
func extractValidationErrors(err error) []map[string]string {
	if err == nil {
		return nil
	}

	var errors []map[string]string

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, validationErr := range validationErrs {
			field := validationErr.Field()
			tag := validationErr.Tag()
			value := validationErr.Value()

			errorMap := make(map[string]string)
			errorMap["field"] = field
			errorMap["message"] = getErrorMsg(field, tag, value)

			errors = append(errors, errorMap)
		}
	}

	return errors
}

func getErrorMsg(field, tag string, value interface{}) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "min":
		return fmt.Sprintf("%s must be at least %v characters long", field, value)
	case "max":
		return fmt.Sprintf("%s must not be longer than %v characters", field, value)
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}
