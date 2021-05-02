package server

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database holds the MongoDB dependencies
type Database struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// Setup MongoDB database
func SetupDatabase(
	ctx context.Context,
	databaseURI string,
	databaseName string,
	databaseUser string,
	databasePassword string,
) (*Database, error) {
	var err error
	db := &Database{}
	db.Client, err = mongo.NewClient(
		options.Client().ApplyURI(databaseURI),
		options.Client().SetAuth(options.Credential{Username: databaseUser, Password: databasePassword}))
	if err != nil {
		return nil, err
	}

	ctxTO, cancel := context.WithTimeout(ctx, 15*time.Second)

	defer cancel()

	if err = db.Client.Connect(ctxTO); err != nil {
		return nil, err
	}

	db.Database = db.Client.Database(databaseName)

	return db, nil
}
