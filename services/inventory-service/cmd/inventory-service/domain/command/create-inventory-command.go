package command

import (
	"inventory-service/cmd/inventory-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
)

type CreateInventoryCommand struct {
	Email                  string
	EventStoreStreamWriter eventstorelib.EventStoreStreamWriter
}

func NewCreateInventoryCommand(email string, evtStore eventstorelib.EventStoreStreamWriter) *CreateInventoryCommand {
	return &CreateInventoryCommand{
		Email:                  email,
		EventStoreStreamWriter: evtStore,
	}
}

func (c *CreateInventoryCommand) Handle() (*eventlib.BaseEvent, error) {
	evt := event.NewInventoryWasCreatedEvent(c.Email)

	_, err := c.EventStoreStreamWriter.StoreEventStream(evt)
	if err != nil {
		return nil, err
	}

	return evt, nil
}
