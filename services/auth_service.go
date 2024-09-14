package services

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/harsh-solanki21/golang-gin-crud-api/models"
	"github.com/harsh-solanki21/golang-gin-crud-api/repositories"
	"github.com/harsh-solanki21/golang-gin-crud-api/utils"
	"github.com/harsh-solanki21/golang-gin-crud-api/validations"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService(userRepository *repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (as *AuthService) Register(user *models.User) error {
	if err := validations.ValidateUserCreate(user); err != nil {
		return utils.NewCustomError(http.StatusBadRequest, "Validation error", err)
	}

	existingUser, _ := as.userRepository.GetUserByEmail(user.Email)
	if existingUser != nil {
		return utils.NewCustomError(http.StatusConflict, "Email already in use", nil)
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Error hashing password", err)
	}
	user.Password = hashedPassword

	user.Role = "user" // Default role for new users
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = as.userRepository.CreateUser(context.Background(), user)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Error creating user", err)
	}

	return nil
}

func (as *AuthService) Login(c *gin.Context, email, password string) error {
	user, err := as.userRepository.GetUserByEmail(email)
	if err != nil {
		return utils.NewCustomError(http.StatusUnauthorized, "Invalid email or password", nil)
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return utils.NewCustomError(http.StatusUnauthorized, "Invalid email or password", nil)
	}

	accessToken, err := utils.GenerateAccessToken(user.ID.Hex(), user.Role)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Error generating access token", err)
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID.Hex())
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Error generating refresh token", err)
	}

	// Set the access token as an HTTP-only cookie
	c.SetCookie(
		"access_token",
		accessToken,
		int(15*time.Minute.Seconds()), // 15 Minutes
		"/",
		"",
		false,
		true,
	)

	// Set the refresh token as an HTTP-only cookie
	c.SetCookie(
		"refresh_token",
		refreshToken,
		int(7*24*time.Hour.Seconds()), // 7 Days
		"/",
		"",
		false,
		true,
	)

	return nil
}

func (as *AuthService) Logout(c *gin.Context) {
	// Clear the access token cookie
	c.SetCookie(
		"access_token",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)

	// Clear the refresh token cookie
	c.SetCookie(
		"refresh_token",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)
}

func (as *AuthService) RefreshToken(c *gin.Context) error {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		return utils.NewCustomError(http.StatusUnauthorized, "Refresh token not found", err)
	}

	claims, err := utils.ValidateRefreshToken(refreshToken)
	if err != nil {
		return utils.NewCustomError(http.StatusUnauthorized, "Invalid refresh token", err)
	}

	userId, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		return utils.NewCustomError(http.StatusUnauthorized, "Invalid user ID", err)
	}

	user, err := as.userRepository.GetUser(context.Background(), userId)
	if err != nil {
		return utils.NewCustomError(http.StatusUnauthorized, "User not found", err)
	}

	newAccessToken, err := utils.GenerateAccessToken(user.ID.Hex(), user.Role)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Error generating new access token", err)
	}

	// Set the new access token as an HTTP-only cookie
	c.SetCookie(
		"access_token",
		newAccessToken,
		int(15*time.Minute.Seconds()),
		"/",
		"",
		false,
		true,
	)

	return nil
}
