package userdb

import (
	"context"
	"htmx-scorecard/src/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Insert(user *types.User, userCollection *mongo.Collection) (*types.HttpErr) {
	var userExists types.User
	err := userCollection.FindOne(context.Background(), bson.D{
		{Key: "email", Value:user.Email},
	}).Decode(&userExists)
	if err == nil && err != mongo.ErrNoDocuments {
		return types.NewHttpErr(400, "user already exists")
	}
	result, err := userCollection.InsertOne(context.Background(), bson.D{
		{Key: "email", Value: user.Email},
		{Key: "password", Value: user.Password},
	})
	if err != nil {
		return types.NewHttpErr(500, "internal server error")
	}
	stringId := result.InsertedID
	objectId, ok := stringId.(primitive.ObjectID)
	if !ok {
		return types.NewHttpErr(500, "internal server error")
	}
	user.Id = objectId
	return nil
}