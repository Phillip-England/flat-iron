package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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