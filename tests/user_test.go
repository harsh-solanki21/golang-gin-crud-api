package tests

import (
	"testing"
	"time"

	"github.com/harsh-solanki21/golang-gin-crud-api/models"
	"github.com/harsh-solanki21/golang-gin-crud-api/validations"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserValidation(t *testing.T) {
	tests := []struct {
		name string
		user models.User
		want bool
	}{
		{
			name: "Valid User",
			user: models.User{
				Name:     "John Doe",
				Email:    "john@example.com",
				Password: "password123",
				Age:      30,
				Role:     "user",
			},
			want: true,
		},
		{
			name: "Invalid Email",
			user: models.User{
				Name:     "Jane Doe",
				Email:    "invalid-email",
				Password: "password123",
				Age:      25,
				Role:     "user",
			},
			want: false,
		},
		{
			name: "Short Password",
			user: models.User{
				Name:     "Bob Smith",
				Email:    "bob@example.com",
				Password: "short",
				Age:      40,
				Role:     "user",
			},
			want: false,
		},
		{
			name: "Invalid Age",
			user: models.User{
				Name:     "Alice Johnson",
				Email:    "alice@example.com",
				Password: "password123",
				Age:      150,
				Role:     "user",
			},
			want: false,
		},
		{
			name: "Invalid Role",
			user: models.User{
				Name:     "Charlie Brown",
				Email:    "charlie@example.com",
				Password: "password123",
				Age:      35,
				Role:     "superuser",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := validations.ValidateUser(&tt.user)
			assert.Equal(t, tt.want, len(errors) == 0)
		})
	}
}

func TestUserMarshalBSON(t *testing.T) {
	user := &models.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
		Age:      25,
		Role:     "user",
	}

	_, err := user.MarshalBSON()
	assert.NoError(t, err)

	assert.False(t, user.CreatedAt.IsZero())
	assert.False(t, user.UpdatedAt.IsZero())
	assert.True(t, user.UpdatedAt.After(user.CreatedAt) || user.UpdatedAt.Equal(user.CreatedAt))
}

func TestUserToJSON(t *testing.T) {
	user := &models.User{
		ID:        primitive.NewObjectID(),
		Name:      "Test User",
		Email:     "test@example.com",
		Password:  "password123",
		Age:       25,
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	json := user.ToJSON()

	assert.Equal(t, user.Name, json["name"])
	assert.Equal(t, user.Email, json["email"])
	assert.Equal(t, user.Role, json["role"])
	assert.NotContains(t, json, "password")
	assert.NotContains(t, json, "age")
	assert.NotContains(t, json, "created_at")
	assert.NotContains(t, json, "updated_at")
}
