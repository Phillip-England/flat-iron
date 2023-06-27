package userdb

import (
	"context"
	"htmx-scorecard/src/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Exists(user *types.User, userCollection *mongo.Collection) (*types.HttpErr) {
	err := userCollection.FindOne(context.Background(), bson.D{{
		Key: "email", Value: user.Email,
	}}).Decode(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return types.NewHttpErr(400, "invalid credentials")
		}
		return types.NewHttpErr(500, "internal server error")
	}
	return nil
}