package userdb

import (
	"context"
	"htmx-scorecard/src/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUserBySession(token string, mongoStore *types.MongoStore) (*types.User, *types.HttpErr) {
	var session types.Session
	objectId, err := primitive.ObjectIDFromHex(token)
	if err != nil {
		return nil, types.NewHttpErr(0, 401, "unauthorized")
	}
	err = mongoStore.SessionCollection.FindOne(context.Background(), bson.D{{
		Key: "_id", Value: objectId,
	}}).Decode(&session)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, types.NewHttpErr(1, 401, "unauthorized")
		}
		return nil, types.NewHttpErr(2, 500, "internal server error")
	}
	expiration := session.Expiration
	currentTime := time.Now().UTC()
	if expiration.Before(currentTime) {
		return nil, types.NewHttpErr(3, 401, "unauthorized")
	}
	var user types.User
	err = mongoStore.UserCollection.FindOne(context.Background(), bson.D{
		{Key: "_id", Value: session.User},
	}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, types.NewHttpErr(4, 401, "unauthorized")
		}
		return nil, types.NewHttpErr(5, 500, "internal server error")
	}
	return &user, nil
} 