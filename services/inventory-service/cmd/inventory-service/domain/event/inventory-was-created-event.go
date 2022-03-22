package event

import (
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryWasCreatedEvent struct {
	AccountId primitive.ObjectID
}

func NewInventoryWasCreatedEvent(accountId primitive.ObjectID) *eventlib.BaseEvent {
	payload := InventoryWasCreatedEvent{
		AccountId: accountId,
	}

	return eventlib.NewBaseEvent("inventory-was-created", primitive.NewObjectID(), payload)
}
