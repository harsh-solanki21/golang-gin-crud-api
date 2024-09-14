package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name" validate:"required,min=2,max=100"`
	Description string             `bson:"description" json:"description" validate:"required,max=500"`
	Price       float64            `bson:"price" json:"price" validate:"required,gte=0"`
	Category    string             `bson:"category" json:"category" validate:"required"`
	InStock     bool               `bson:"in_stock" json:"in_stock"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

func (p *Product) MarshalBSON() ([]byte, error) {
	if p.CreatedAt.IsZero() {
		p.CreatedAt = time.Now()
	}
	p.UpdatedAt = time.Now()

	type my Product
	return bson.Marshal((*my)(p))
}

// Response Data to send
func (p *Product) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
		"price":       p.Price,
		"category":    p.Category,
		"in_stock":    p.InStock,
	}
}
