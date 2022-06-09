package command

import (
	"transaction-service/cmd/transaction-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
)

type CreateTransactionCommand struct {
	Quantity               int64
	ValuePaid              int64
	Description            string
	EventStoreStreamWriter eventstorelib.EventStoreStreamWriter
}

// Handle stores an TransactionWasCreatedEvent to the event store and returns the resulting event
func (c *CreateTransactionCommand) Handle() (*eventlib.BaseEvent, error) {
	if c.Quantity <= 0 {
		return nil, ErrQuantityMustBeGreaterThanZero
	}

	evt := event.NewTransactionWasCreatedEvent(c.Quantity, c.ValuePaid, c.Description)

	_, err := c.EventStoreStreamWriter.StoreEventStream(evt)
	if err != nil {
		return nil, err
	}

	return evt, nil
}
