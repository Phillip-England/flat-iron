package sessiondb

import (
	"context"
	"htmx-scorecard/src/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Insert(session *types.Session, sessionCollection *mongo.Collection) (*types.HttpErr) {
	result, err := sessionCollection.InsertOne(context.Background(), bson.D{
		{Key: "user", Value: session.User},
		{Key: "expiration", Value: session.Expiration},
	})
	if err != nil {
		return types.NewHttpErr(500, "internal server error")
	}
	stringId := result.InsertedID
	objectId, ok := stringId.(primitive.ObjectID)
	if !ok {
		return types.NewHttpErr(500, "internal server error")
	}
	session.Id = objectId
	return nil
}