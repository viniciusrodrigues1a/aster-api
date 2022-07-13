package command

import (
	"expense-service/cmd/expense-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type CreateExpenseCommand struct {
	ProductID               *string
	Title                   string
	Description             string
	Value                   int64
	EventStoreStreamWriter  eventstorelib.EventStoreStreamWriter
	ProductStateStoreReader statestorelib.StateStoreReader
}

// Handle stores an ExpenseWasCreatedEvent to the event store and returns the resulting event
func (c *CreateExpenseCommand) Handle() (*eventlib.BaseEvent, error) {
	if c.Title == "" {
		return nil, ErrTitleIsRequired
	}

	if c.ProductID != nil {
		_, err := c.ProductStateStoreReader.ReadState(*c.ProductID)
		if err != nil {
			return nil, ErrProductCouldntBeFound
		}
	}

	if c.Value == 0 {
		return nil, ErrValueCantBeZero
	}

	evt := event.NewExpenseWasCreatedEvent(c.ProductID, c.Title, c.Description, c.Value)

	_, storeErr := c.EventStoreStreamWriter.StoreEventStream(evt)
	if storeErr != nil {
		return nil, storeErr
	}

	return evt, nil
}
