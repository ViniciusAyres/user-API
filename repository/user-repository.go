package repository

import (
	"context"
	"fmt"

	"github.com/ViniciusAyres/user-API/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -destination=../mocks/user-repository-mock.go -package=mocks github.com/ViniciusAyres/user-API/repository UserRepository

// UserRepository ...
type UserRepository interface {
	Save(u model.User) model.User
}

// Save ...
func (m *MongoDB) Save(u model.User) model.User {
	collection := m.db.Collection("User")
	u.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(context.Background(), u)
	if err != nil {
		fmt.Println(err)
	}

	u.ID = res.InsertedID.(primitive.ObjectID)

	return u
}
