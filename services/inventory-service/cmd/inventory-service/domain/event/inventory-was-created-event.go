package event

import (
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryWasCreatedEvent struct {
	Email string
}

func NewInventoryWasCreatedEvent(email string) *eventlib.BaseEvent {
	payload := InventoryWasCreatedEvent{
		Email: email,
	}

	return eventlib.NewBaseEvent("inventory-was-created", primitive.NewObjectID(), payload)
}
