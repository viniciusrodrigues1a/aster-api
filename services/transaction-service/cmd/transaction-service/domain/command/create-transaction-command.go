package command

import (
	"transaction-service/cmd/transaction-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type CreateTransactionCommand struct {
	ProductID               *string
	Quantity                int64
	ValuePaid               int64
	Description             string
	EventStoreStreamWriter  eventstorelib.EventStoreStreamWriter
	ProductStateStoreReader statestorelib.StateStoreReader
}

// Handle stores an TransactionWasCreatedEvent to the event store and returns the resulting event
func (c *CreateTransactionCommand) Handle() (*eventlib.BaseEvent, error) {
	if c.Quantity <= 0 {
		return nil, ErrQuantityMustBeGreaterThanZero
	}

	if c.ProductID == nil {
		return nil, ErrProductIDIsRequired
	}

	_, stateErr := c.ProductStateStoreReader.ReadState(*c.ProductID)
	if stateErr != nil {
		return nil, ErrProductCouldntBeFound
	}

	evt := event.NewTransactionWasCreatedEvent(c.ProductID, c.Quantity, c.ValuePaid, c.Description)

	_, err := c.EventStoreStreamWriter.StoreEventStream(evt)
	if err != nil {
		return nil, err
	}

	return evt, nil
}
