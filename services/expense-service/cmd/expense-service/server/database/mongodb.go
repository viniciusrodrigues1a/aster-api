package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoConnection struct {
	Context context.Context
	Client  *mongo.Client
}

var MongoConn *mongoConnection

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	uri := "mongodb://localhost:27014"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	MongoConn = &mongoConnection{
		Context: ctx,
		Client:  client,
	}
}

func StopMongo() {
	MongoConn.Client.Disconnect(MongoConn.Context)
}
