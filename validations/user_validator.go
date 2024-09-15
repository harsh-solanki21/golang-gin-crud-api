package validations

import (
	"github.com/harsh-solanki21/golang-gin-crud-api/models"
)

func ValidateUser(user *models.User) []map[string]string {
	return extractValidationErrors(validate.Struct(user))
}

func ValidateUserCreate(user *models.User) []map[string]string {
	return extractValidationErrors(validate.Struct(user))
}

func ValidateUserUpdate(user *models.User) []map[string]string {
	return extractValidationErrors(validate.StructPartial(user))
}
