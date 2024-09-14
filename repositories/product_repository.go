package repositories

import (
	"context"
	"errors"

	"github.com/harsh-solanki21/golang-gin-crud-api/configs"
	"github.com/harsh-solanki21/golang-gin-crud-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrProductNotFound = errors.New("product not found")

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(client *mongo.Client) *ProductRepository {
	dbName := configs.GetDatabaseName()
	collection := client.Database(dbName).Collection("products")
	return &ProductRepository{
		collection: collection,
	}
}

func (pr *ProductRepository) CreateProduct(ctx context.Context, product *models.Product) error {
	_, err := pr.collection.InsertOne(ctx, product)
	return err
}

func (pr *ProductRepository) GetProduct(ctx context.Context, id primitive.ObjectID) (*models.Product, error) {
	var product models.Product
	err := pr.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrProductNotFound
		}
		return nil, err
	}
	return &product, nil
}

func (pr *ProductRepository) UpdateProduct(ctx context.Context, id primitive.ObjectID, update bson.M) (*models.Product, error) {
	filter := bson.M{"_id": id}
	updateDoc := bson.M{
		"$set": update,
	}

	result, err := pr.collection.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, ErrProductNotFound
	}

	return pr.GetProduct(ctx, id)
}

func (pr *ProductRepository) DeleteProduct(ctx context.Context, id primitive.ObjectID) error {
	result, err := pr.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return ErrProductNotFound
	}
	return nil
}

func (pr *ProductRepository) ListProducts(ctx context.Context, limit int, offset int, sort string) ([]*models.Product, int64, error) {
	options := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(offset)).
		SetSort(bson.M{sort: 1})

	cursor, err := pr.collection.Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var products []*models.Product
	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, 0, err
		}
		products = append(products, &product)
	}

	if err := cursor.Err(); err != nil {
		return nil, 0, err
	}

	// Get total count
	totalCount, err := pr.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	return products, totalCount, nil
}
