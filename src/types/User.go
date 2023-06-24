package types

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id primitive.ObjectID `bson:"_id" json:"_id"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

func NewUser(email string, password string) (*User) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &User{
		Email: email,
		Password: string(hashedPassword),
	}
}

func (m *User) Insert(userCollection *mongo.Collection) (*HttpErr) {
	var userExists User
	err := userCollection.FindOne(context.Background(), bson.D{
		{Key: "email", Value:m.Email},
	}).Decode(&userExists)
	if err == nil && err != mongo.ErrNoDocuments {
		return NewHttpErr(0, 400, "user already exists")
	}
	result, err := userCollection.InsertOne(context.Background(), bson.D{
		{Key: "email", Value: m.Email},
		{Key: "password", Value: m.Password},
	})
	if err != nil {
		return NewHttpErr(1, 500, "internal server error")
	}
	stringId := result.InsertedID
	objectId, ok := stringId.(primitive.ObjectID)
	if !ok {
		return NewHttpErr(2, 500, "internal server error")
	}
	m.Id = objectId
	return nil
}

func (m *User) Find(userCollection *mongo.Collection) (*HttpErr) {
	err := userCollection.FindOne(context.Background(), bson.D{{
		Key: "email", Value: m.Email,
	}}).Decode(m)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return NewHttpErr(0, 400, "invalid credentials")
		}
		return NewHttpErr(1, 500, "internal server error")
	}
	return nil
}
