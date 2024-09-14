package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func SuccessResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string, err interface{}) APIResponse {
	return APIResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}

func RespondWithError(c *gin.Context, code int, message string, err interface{}) {
	c.JSON(code, ErrorResponse(message, err))
}

func RespondWithSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, SuccessResponse(message, data))
}

func HandleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case *CustomError:
		RespondWithError(c, e.StatusCode, e.Message, e.Details)
	default:
		RespondWithError(c, http.StatusInternalServerError, "Internal Server Error", nil)
	}
}

type CustomError struct {
	StatusCode int
	Message    string
	Details    interface{}
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewCustomError(statusCode int, message string, details interface{}) *CustomError {
	return &CustomError{
		StatusCode: statusCode,
		Message:    message,
		Details:    details,
	}
}
