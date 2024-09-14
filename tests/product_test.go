package tests

import (
	"testing"
	"time"

	"github.com/harsh-solanki21/golang-gin-crud-api/models"
	"github.com/harsh-solanki21/golang-gin-crud-api/validations"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestProductValidation(t *testing.T) {
	tests := []struct {
		name    string
		product models.Product
		want    bool
	}{
		{
			name: "Valid Product",
			product: models.Product{
				Name:        "Test Product 1",
				Description: "This is a test product 1",
				Price:       9.99,
				Category:    "Test Category 1",
				InStock:     true,
			},
			want: true,
		},
		{
			name: "Short Name",
			product: models.Product{
				Name:        "T",
				Description: "This is a test product 2",
				Price:       9.99,
				Category:    "Test Category 2",
				InStock:     true,
			},
			want: false,
		},
		{
			name: "Long Name",
			product: models.Product{
				Name:        string(make([]rune, 101)),
				Description: "This is a test product 3",
				Price:       9.99,
				Category:    "Test Category 3",
				InStock:     true,
			},
			want: false,
		},
		{
			name: "Long Description",
			product: models.Product{
				Name:        "Test Product 4",
				Description: string(make([]rune, 501)),
				Price:       9.99,
				Category:    "Test Category 4",
				InStock:     true,
			},
			want: false,
		},
		{
			name: "Negative Price",
			product: models.Product{
				Name:        "Test Product 5",
				Description: "This is a test product 5",
				Price:       -1.00,
				Category:    "Test Category 5",
				InStock:     true,
			},
			want: false,
		},
		{
			name: "Empty Category",
			product: models.Product{
				Name:        "Test Product 6",
				Description: "This is a test product 6",
				Price:       9.99,
				Category:    "",
				InStock:     true,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := validations.ValidateProduct(&tt.product)
			assert.Equal(t, tt.want, len(errors) == 0)
		})
	}
}

func TestProductMarshalBSON(t *testing.T) {
	product := &models.Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Category:    "Test Category",
		InStock:     true,
	}

	_, err := product.MarshalBSON()
	assert.NoError(t, err)

	assert.False(t, product.CreatedAt.IsZero())
	assert.False(t, product.UpdatedAt.IsZero())
	assert.True(t, product.UpdatedAt.After(product.CreatedAt) || product.UpdatedAt.Equal(product.CreatedAt))
}

func TestProductToJSON(t *testing.T) {
	product := &models.Product{
		ID:          primitive.NewObjectID(),
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       9.99,
		Category:    "Test Category",
		InStock:     true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	json := product.ToJSON()

	assert.Equal(t, product.Name, json["name"])
	assert.Equal(t, product.Description, json["description"])
	assert.Equal(t, product.Price, json["price"])
	assert.Equal(t, product.Category, json["category"])
	assert.Equal(t, product.InStock, json["in_stock"])
	assert.NotContains(t, json, "created_at")
	assert.NotContains(t, json, "updated_at")
}
