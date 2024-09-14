package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh-solanki21/golang-gin-crud-api/models"
	"github.com/harsh-solanki21/golang-gin-crud-api/services"
	"github.com/harsh-solanki21/golang-gin-crud-api/utils"
	"github.com/harsh-solanki21/golang-gin-crud-api/validations"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := validations.ValidateUserCreate(&user); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Validation error", err)
		return
	}

	if err := uc.userService.CreateUser(&user); err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, "User created successfully", user.ToJSON())
}

func (uc *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.userService.GetUser(id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "User retrieved successfully", user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := validations.ValidateUserUpdate(&user); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Validation error", err)
		return
	}

	updatedUser, err := uc.userService.UpdateUser(id, &user)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "User updated successfully", updatedUser)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := uc.userService.DeleteUser(id); err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "User deleted successfully", nil)
}

func (uc *UserController) ListUsers(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	paginatedData, err := uc.userService.ListUsers(pagination)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Users retrieved successfully", paginatedData)
}
