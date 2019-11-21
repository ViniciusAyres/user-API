package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User ...
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Birthdate string             `json:"birthdate" bson:"birthdate"`
	CreatedAt string             `json:"createdAt" bson:"createdAt"`
}
