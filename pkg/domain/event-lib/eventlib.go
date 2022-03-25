package eventlib

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseEventData struct {
	StreamId  primitive.ObjectID `bson:"stream_id"`
	Id        primitive.ObjectID `bson:"id"`
	Name      string             `bson:"name"`
	CreatedAt int64              `bson:"created_at"`
}

type BaseEvent struct {
	Data    BaseEventData
	Payload interface{}
}

func NewBaseEvent(name string, streamId primitive.ObjectID, payload interface{}) *BaseEvent {
	return &BaseEvent{
		Data: BaseEventData{
			StreamId:  streamId,
			Id:        primitive.NewObjectID(),
			Name:      name,
			CreatedAt: time.Now().Unix(),
		},
		Payload: payload,
	}
}
