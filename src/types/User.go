package types

import (
	"context"
	"time"

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

func FindUserBySession(token string, mongoStore *MongoStore) (*User, *HttpErr) {
	var session Session
	objectId, err := primitive.ObjectIDFromHex(token)
	if err != nil {
		return nil, NewHttpErr(0, 401, "unauthorized")
	}
	err = mongoStore.SessionCollection.FindOne(context.Background(), bson.D{{
		Key: "_id", Value: objectId,
	}}).Decode(&session)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, NewHttpErr(1, 401, "unauthorized")
		}
		return nil, NewHttpErr(2, 500, "internal server error")
	}
	expiration := session.Expiration
	currentTime := time.Now().UTC()
	if expiration.Before(currentTime) {
		return nil, NewHttpErr(3, 401, "unauthorized")
	}
	var user User
	err = mongoStore.UserCollection.FindOne(context.Background(), bson.D{
		{Key: "_id", Value: session.User},
	}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, NewHttpErr(4, 401, "unauthorized")
		}
		return nil, NewHttpErr(5, 500, "internal server error")
	}
	return &user, nil
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

func (m *User) Exists(userCollection *mongo.Collection) (*HttpErr) {
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


