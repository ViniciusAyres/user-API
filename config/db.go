package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConfigDB Start db conection
func ConfigDB(ctx context.Context) (*mongo.Database, error) {
	uri := fmt.Sprintf(`mongodb://%s:%s@%s/%s`,
		"userAPI",
		"senha123",
		"127.0.0.1:27017",
		"test",
	// ctx.Value(usernameKey).(string),
	// ctx.Value(passwordKey).(string),
	// ctx.Value(hostKey).(string),
	// ctx.Value(databaseKey).(string),
	)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("todo: couldn't connect to mongo: %v", err)
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("todo: mongo client couldn't connect with background context: %v", err)
	}
	DB := client.Database("test")
	return DB, nil
}
