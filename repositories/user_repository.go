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

var ErrUserNotFound = errors.New("user not found")

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	dbName := configs.GetDatabaseName()
	collection := client.Database(dbName).Collection("users")
	return &UserRepository{
		collection: collection,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	_, err := ur.collection.InsertOne(ctx, user)
	return err
}

func (ur *UserRepository) GetUser(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := ur.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, err
}

func (ur *UserRepository) UpdateUser(ctx context.Context, id primitive.ObjectID, update bson.M) (*models.User, error) {
	filter := bson.M{"_id": id}
	updateDoc := bson.M{
		"$set": update,
	}

	result, err := ur.collection.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, ErrUserNotFound
	}

	return ur.GetUser(ctx, id)
}

func (ur *UserRepository) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	result, err := ur.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return ErrUserNotFound
	}
	return nil
}

func (ur *UserRepository) ListUsers(ctx context.Context, limit int, offset int, sort string) ([]*models.User, int64, error) {
	options := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(offset)).
		SetSort(bson.M{sort: 1})

	cursor, err := ur.collection.Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var users []*models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, 0, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, 0, err
	}

	// Get total count
	totalCount, err := ur.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := ur.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	return &user, err
}
