package sessiondb

import (
	"context"
	"htmx-scorecard/src/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ClearUserSessions(session *types.Session, sessionCollection *mongo.Collection) (*types.HttpErr) {
	_, err := sessionCollection.DeleteMany(context.Background(), bson.D{
		{Key: "user", Value: session.User},
	})
	if err != nil {
		return types.NewHttpErr(500, "internal server error")
	}
	return nil
}

