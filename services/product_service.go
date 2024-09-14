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
	ErrProductNotFoundMessage = "Product not found"
	ErrInvalidIdMessage       = "Invalid product ID"
)

type ProductService struct {
	productRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (ps *ProductService) CreateProduct(product *models.Product) error {
	if validationErrors := validations.ValidateProduct(product); validationErrors != nil {
		return fmt.Errorf("validation error: %v", validationErrors)
	}

	return ps.productRepository.CreateProduct(context.Background(), product)
}

func (ps *ProductService) GetProduct(id string) (*models.Product, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, utils.NewCustomError(400, ErrInvalidIdMessage, err)
	}

	product, err := ps.productRepository.GetProduct(context.Background(), objectID)
	if err != nil {
		if errors.Is(err, repositories.ErrProductNotFound) {
			return nil, utils.NewCustomError(404, ErrProductNotFoundMessage, err)
		}
		return nil, utils.NewCustomError(500, "Error retrieving product", err)
	}

	return product, nil
}

func (ps *ProductService) UpdateProduct(id string, product *models.Product) (*models.Product, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, utils.NewCustomError(400, ErrInvalidIdMessage, err)
	}

	validationErrors := validations.ValidateProductUpdate(product)
	if validationErrors != nil {
		return nil, fmt.Errorf("validation error: %v", validationErrors)
	}

	update := bson.M{}
	if product.Name != "" {
		update["name"] = product.Name
	}
	if product.Description != "" {
		update["description"] = product.Description
	}
	if product.Price != 0 {
		update["price"] = product.Price
	}

	updatedProduct, err := ps.productRepository.UpdateProduct(context.Background(), objectID, update)
	if err != nil {
		if errors.Is(err, repositories.ErrProductNotFound) {
			return nil, utils.NewCustomError(404, ErrProductNotFoundMessage, err)
		}
		return nil, utils.NewCustomError(500, "Error updating product", err)
	}

	return updatedProduct, nil
}

func (ps *ProductService) DeleteProduct(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return utils.NewCustomError(400, ErrInvalidIdMessage, err)
	}

	err = ps.productRepository.DeleteProduct(context.Background(), objectID)
	if err != nil {
		if errors.Is(err, repositories.ErrProductNotFound) {
			return utils.NewCustomError(404, ErrProductNotFoundMessage, err)
		}
		return utils.NewCustomError(500, "Error deleting product", err)
	}

	return nil
}

func (ps *ProductService) ListProducts(pagination utils.Pagination) (utils.PaginatedResponse, error) {
	products, totalRows, err := ps.productRepository.ListProducts(
		context.Background(),
		pagination.GetLimit(),
		pagination.GetOffset(),
		pagination.GetSort(),
	)
	if err != nil {
		return utils.PaginatedResponse{}, utils.NewCustomError(500, "Error listing products", err)
	}

	return pagination.GenerateResponse(products, totalRows), nil
}
