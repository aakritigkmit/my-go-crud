package repository

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aakritigkmit/my-go-crud/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	Client *mongo.Collection
}

func NewUserRepo(client *mongo.Collection) *UserRepo {
	return &UserRepo{
		Client: client,
	}
}

func (c *UserRepo) CreateUser(user model.User) (string, error) {
	result, err := c.Client.InsertOne(context.Background(), user)
	if err != nil {
		return "", nil
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
	//
	// Hex returns the hex encoding of the ObjectID as a string.
}

func (c *UserRepo) GetUserByID(id string) (model.User, error) {
	docId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return model.User{}, fmt.Errorf("invalid  ID")
	}
	var user model.User
	filter := bson.D{{Key: "_id", Value: docId}}
	// // D is an ordered representation of a BSON document. This type should be used when the order of the elements matters,
	// such as MongoDB command documents.
	err = c.Client.FindOne(context.Background(), filter).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return model.User{}, fmt.Errorf(" no document found")
	} else if err != nil {
		return model.User{}, err
	}

	return user, nil
}
func (c *UserRepo) GetAllUsers() ([]model.User, error) {
	filter := bson.D{}

	cur, err := c.Client.Find(context.Background(), filter)
	if err != nil {
		return []model.User{}, err
	}
	defer cur.Close(context.Background())

	var users []model.User

	for cur.Next(context.Background()) {
		var user model.User
		err := cur.Decode(&user)
		if err != nil {
			slog.Error("error while decoding get users", slog.String("error", err.Error()))
			continue
		}
		users = append(users, user)

	}

	return users, nil
}

func (c *UserRepo) UpdateUserAgeByID(id string, age int) (int, error) {
	docId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return 0, fmt.Errorf("invalid  ID")
	}
	filter := bson.D{{Key: "_id", Value: docId}}

	updateStmt := bson.D{{Key: "$set", Value: bson.D{{Key: "age", Value: age}}}}
	result, err := c.Client.UpdateOne(context.Background(), filter, updateStmt)
	if err != nil {
		return 0, err
	}
	return int(result.ModifiedCount), nil
}

func (c *UserRepo) DeleteUserByID(id string) (int, error) {
	docId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return 0, fmt.Errorf("invalid  ID")
	}
	filter := bson.D{{Key: "_id", Value: docId}}
	result, err := c.Client.DeleteOne(context.Background(), filter)

	if err != nil {
		return 0, err
	}
	return int(result.DeletedCount), nil
}

func (c *UserRepo) DeleteAllUsers() (int, error) {
	filter := bson.D{}

	result, err := c.Client.DeleteMany(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return int(result.DeletedCount), nil
}
