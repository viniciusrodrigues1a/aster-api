package projector

import (
	"inventory-service/cmd/inventory-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryCreationProjector struct{}

func (i *InventoryCreationProjector) Project(e *eventlib.BaseEvent) *InventoryState {
	payload := e.Payload.(event.InventoryWasCreatedEvent)

	return &InventoryState{
		AccountId:    payload.AccountId,
		Participants: []primitive.ObjectID{},
	}
}
