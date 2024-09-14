package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/harsh-solanki21/golang-gin-crud-api/models"
	"github.com/harsh-solanki21/golang-gin-crud-api/repositories"
	"github.com/harsh-solanki21/golang-gin-crud-api/utils"
	"github.com/harsh-solanki21/golang-gin-crud-api/validations"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ErrUserNotFoundMessage = "User not found"
	ErrInvalidUserId       = "Invalid user ID"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (us *UserService) CreateUser(user *models.User) error {
	if validationErrors := validations.ValidateUser(user); validationErrors != nil {
		return fmt.Errorf("validation error: %v", validationErrors)
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return us.userRepository.CreateUser(context.Background(), user)
}

func (us *UserService) GetUser(id string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, utils.NewCustomError(400, ErrInvalidUserId, err)
	}

	user, err := us.userRepository.GetUser(context.Background(), objectID)
	if err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) {
			return nil, utils.NewCustomError(404, ErrUserNotFoundMessage, err)
		}
		return nil, utils.NewCustomError(500, "Error retrieving user", err)
	}

	return user, nil
}

func (us *UserService) UpdateUser(id string, user *models.User) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, utils.NewCustomError(400, ErrInvalidIdMessage, err)
	}

	validationErrors := validations.ValidateUserUpdate(user)
	if validationErrors != nil {
		return nil, fmt.Errorf("validation error: %v", validationErrors)
	}

	update := bson.M{}
	if user.Name != "" {
		update["name"] = user.Name
	}
	if user.Email != "" {
		update["email"] = user.Email
	}

	updatedUser, err := us.userRepository.UpdateUser(context.Background(), objectID, update)
	if err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) {
			return nil, utils.NewCustomError(404, ErrUserNotFoundMessage, err)
		}
		return nil, utils.NewCustomError(500, "Error updating user", err)
	}

	return updatedUser, nil
}

func (us *UserService) DeleteUser(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return utils.NewCustomError(400, ErrInvalidIdMessage, err)
	}

	err = us.userRepository.DeleteUser(context.Background(), objectID)
	if err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) {
			return utils.NewCustomError(404, ErrUserNotFoundMessage, err)
		}
		return utils.NewCustomError(500, "Error deleting user", err)
	}

	return nil
}

func (us *UserService) ListUsers(pagination utils.Pagination) (utils.PaginatedResponse, error) {
	users, totalRows, err := us.userRepository.ListUsers(
		context.Background(),
		pagination.GetLimit(),
		pagination.GetOffset(),
		pagination.GetSort(),
	)
	if err != nil {
		return utils.PaginatedResponse{}, utils.NewCustomError(500, "Error listing users", err)
	}

	return pagination.GenerateResponse(users, totalRows), nil
}
