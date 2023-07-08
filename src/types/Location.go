package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	Id primitive.ObjectID `bson:"_id" json:"_id"`
	User primitive.ObjectID `bson:"user" json:"user"`
	Name string `bson:"name" json:"name"`
	Number string `bson:"number" json:"number"`
}

func NewLocation(userId primitive.ObjectID, name string, number string) (*Location) {
	return &Location{
		User: userId,
		Name: name,
		Number: number,
	}
}