package validations

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/harsh-solanki21/golang-gin-crud-api/models"
)

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

func ValidateProduct(product *models.Product) []map[string]string {
	return extractValidationErrors(validate.Struct(product))
}

func ValidateProductCreate(product *models.Product) []map[string]string {
	return extractValidationErrors(validate.Struct(product))
}

func ValidateProductUpdate(product *models.Product) []map[string]string {
	return extractValidationErrors(validate.StructPartial(product))
}
