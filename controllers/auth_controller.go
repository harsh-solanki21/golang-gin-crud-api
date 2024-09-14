package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh-solanki21/golang-gin-crud-api/models"
	"github.com/harsh-solanki21/golang-gin-crud-api/services"
	"github.com/harsh-solanki21/golang-gin-crud-api/utils"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginRequest models.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	err := ac.authService.Login(c, loginRequest.Email, loginRequest.Password)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Logged in successfully", nil)
}

func (ac *AuthController) Logout(c *gin.Context) {
	ac.authService.Logout(c)
	utils.RespondWithSuccess(c, http.StatusOK, "Logged out successfully", nil)
}

func (ac *AuthController) RefreshToken(c *gin.Context) {
	err := ac.authService.RefreshToken(c)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Access token refreshed successfully", nil)
}
