package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang-rest-api-kata/config"
	"time"
)

func NewMongoDb(c *config.Config) (*mongo.Database, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(c.MongoDB.MongoURI).SetConnectTimeout(time.Second * 5)
	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	defer cancel()

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		return nil, err
	}

	return client.Database(c.MongoDB.MongoDatabase), nil
}

func Disconnect(client *mongo.Client) {
	_ = client.Disconnect(context.TODO())
}
