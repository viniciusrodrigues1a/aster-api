package eventstorelib

import (
	"context"
	"fmt"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoEventStoreRepository struct {
	Context    context.Context
	Client     *mongo.Client
	Collection string
}

func New(context context.Context, client *mongo.Client, collection string) *MongoEventStoreRepository {
	return &MongoEventStoreRepository{
		Context:    context,
		Client:     client,
		Collection: collection,
	}
}

type EventStoreStreamWriter interface {
	StoreEventStream(event *eventlib.BaseEvent) (string, error)
}

func (m *MongoEventStoreRepository) StoreEventStream(e *eventlib.BaseEvent) (string, error) {
	collection := m.Client.Database("aster").Collection(m.Collection)

	document := bson.D{{Key: "_id", Value: e.Data.StreamId}, {Key: "events", Value: []*eventlib.BaseEvent{e}}}
	result, err := collection.InsertOne(m.Context, document)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

type EventStoreWriter interface {
	StoreEvent(event *eventlib.BaseEvent) (string, error)
}

func (m *MongoEventStoreRepository) StoreEvent(e *eventlib.BaseEvent) (string, error) {
	collection := m.Client.Database("aster").Collection(m.Collection)

	filter := bson.D{{Key: "_id", Value: e.Data.StreamId}}
	query := bson.D{{Key: "$push", Value: bson.D{{Key: "events", Value: e}}}}

	result, err := collection.UpdateOne(m.Context, filter, query)
	if err != nil {
		return "", err
	}

	if result.ModifiedCount == 0 {
		return "", fmt.Errorf("Event stream doesn't exist")
	}

	return e.Data.StreamId.Hex(), nil
}
