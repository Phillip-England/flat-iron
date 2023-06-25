package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	Id primitive.ObjectID `bson:"_id" json:"_id"`
	Name string `bson:"name" json:"name"`
	Number string `bson:"number" json:"number"`
}

func NewLocation(name string, number string) (*Location) {
	return &Location{
		Name: name,
		Number: number,
	}
}