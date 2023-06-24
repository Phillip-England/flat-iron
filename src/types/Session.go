package types

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Session struct {
	Id primitive.ObjectID `bson:"_id" json:"_id"`
	User primitive.ObjectID `bson:"user" json:"user"`
	Expiration time.Time `bson:"expiration" json:"expiration"`
}

func NewSession(userId primitive.ObjectID) *Session {
	expiration := time.Now().Add(24 * time.Hour) // 24 hours
	return &Session{
		User:       userId,
		Expiration: expiration.UTC(),
	}
}

func (m *Session) ClearUserSessions(sessionCollection *mongo.Collection) (*HttpErr) {
	_, err := sessionCollection.DeleteMany(context.Background(), bson.D{
		{Key: "user", Value: m.User},
	})
	if err != nil {
		return NewHttpErr(0, 500, "internal server error")
	}
	return nil
}

func (m *Session) Insert(sessionCollection *mongo.Collection) (*HttpErr) {
	result, err := sessionCollection.InsertOne(context.Background(), bson.D{
		{Key: "user", Value: m.User},
		{Key: "expiration", Value: m.Expiration},
	})
	if err != nil {
		return NewHttpErr(0, 500, "internal server error")
	}
	stringId := result.InsertedID
	objectId, ok := stringId.(primitive.ObjectID)
	if !ok {
		return NewHttpErr(1, 500, "internal server error")
	}
	m.Id = objectId
	return nil
}