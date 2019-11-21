package repository

import "go.mongodb.org/mongo-driver/mongo"

// MongoDB ...
type MongoDB struct {
	db *mongo.Database
}

// New ...
func New(db *mongo.Database) *MongoDB {
	return &MongoDB{
		db: db,
	}
}
