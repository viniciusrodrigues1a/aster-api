package event

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type ProductWasUpdatedEvent struct {
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Quantity    int32  `bson:"quantity"`
}

func NewProductWasUpdatedEvent(title, description string, quantity int32, id string) *eventlib.BaseEvent {
	payload := &ProductWasUpdatedEvent{
		Title:       title,
		Description: description,
		Quantity:    quantity,
	}

	oid, _ := primitive.ObjectIDFromHex(id)
	return eventlib.NewBaseEvent("product-was-updated", oid, payload)
}
