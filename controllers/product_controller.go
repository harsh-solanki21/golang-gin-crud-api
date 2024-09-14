package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh-solanki21/golang-gin-crud-api/models"
	"github.com/harsh-solanki21/golang-gin-crud-api/services"
	"github.com/harsh-solanki21/golang-gin-crud-api/utils"
	"github.com/harsh-solanki21/golang-gin-crud-api/validations"
)

type ProductController struct {
	productService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := validations.ValidateProductCreate(&product); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Validation error", err)
		return
	}

	if err := pc.productService.CreateProduct(&product); err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, "Product created successfully", product.ToJSON())
}

func (pc *ProductController) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := pc.productService.GetProduct(id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Product retrieved successfully", product)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := validations.ValidateProductUpdate(&product); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Validation error", err)
		return
	}

	updatedProduct, err := pc.productService.UpdateProduct(id, &product)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Product updated successfully", updatedProduct)
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := pc.productService.DeleteProduct(id); err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Product deleted successfully", nil)
}

func (pc *ProductController) ListProducts(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	paginatedData, err := pc.productService.ListProducts(pagination)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Products retrieved successfully", paginatedData)
}
