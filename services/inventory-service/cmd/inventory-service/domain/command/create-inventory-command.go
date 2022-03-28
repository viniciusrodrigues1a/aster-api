package command

import (
	"inventory-service/cmd/inventory-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateInventoryCommand struct {
	AccountId              primitive.ObjectID
	EventStoreStreamWriter eventstorelib.EventStoreStreamWriter
}

func NewCreateInventoryCommand(accountId primitive.ObjectID, evtStore eventstorelib.EventStoreStreamWriter) *CreateInventoryCommand {
	return &CreateInventoryCommand{
		AccountId:              accountId,
		EventStoreStreamWriter: evtStore,
	}
}

func (c *CreateInventoryCommand) Handle() (*eventlib.BaseEvent, error) {
	evt := event.NewInventoryWasCreatedEvent(c.AccountId)

	_, err := c.EventStoreStreamWriter.StoreEventStream(evt)
	if err != nil {
		return nil, err
	}

	return evt, nil
}
