package validations

import (
	"github.com/harsh-solanki21/golang-gin-crud-api/models"
)

func ValidateProduct(product *models.Product) []map[string]string {
	return extractValidationErrors(validate.Struct(product))
}

func ValidateProductCreate(product *models.Product) []map[string]string {
	return extractValidationErrors(validate.Struct(product))
}

func ValidateProductUpdate(product *models.Product) []map[string]string {
	return extractValidationErrors(validate.StructPartial(product))
}
