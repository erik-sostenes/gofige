package repository

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	syncMongo     sync.Once
	mongoDatabase *mongo.Database
	mongoClient   *mongo.Client
	err           error
)

// NewClientMongo creates a new connection to mongoClient
// use syncOnce to create only one instance of mongoClient
func NewClientMongo(config MongoDB) (*mongo.Database, error) {
	syncMongo.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), config.ConnectTimeout*time.Second)
		defer cancel()

		mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(config.Uri))

		err = mongoClient.Ping(context.TODO(), readpref.Primary())

		mongoDatabase = mongoClient.Database(config.DatabaseName)
		err = mongoDatabase.Client().Ping(context.TODO(), readpref.Primary())
	})
	return mongoDatabase, err
}

func NewMDB(config MongoDB) (mongoDB *mongo.Database) {
	mongoDB, err := NewClientMongo(config)
	if err != nil {
		panic(err)
	}
	return
}
