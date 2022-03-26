package command

import (
	"expense-service/cmd/expense-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
)

type CreateExpenseCommand struct {
	Title                  string
	Description            string
	Value                  int64
	EventStoreStreamWriter eventstorelib.EventStoreStreamWriter
}

// Handle stores an ExpenseWasCreatedEvent to the event store and returns the resulting event
func (c *CreateExpenseCommand) Handle() (*eventlib.BaseEvent, error) {
	evt := event.NewExpenseWasCreatedEvent(c.Title, c.Description, c.Value)

	_, err := c.EventStoreStreamWriter.StoreEventStream(evt)
	if err != nil {
		return nil, err
	}

	return evt, nil
}
