package command

import (
	"transaction-service/cmd/transaction-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type DeleteTransactionCommand struct {
	ID               string
	EventStoreWriter eventstorelib.EventStoreWriter
	StateStoreReader statestorelib.StateStoreReader
}

// Handle stores a TransactionWasDeletedEvent to the event store and returns the resulting event.
// Returns ErrTransactionDoesntExist if it can't read the state from the state store
func (d *DeleteTransactionCommand) Handle() (*eventlib.BaseEvent, error) {
	evt := event.NewTransactionWasDeletedEvent(d.ID)

	_, err := d.StateStoreReader.ReadState(d.ID)
	if err != nil {
		return nil, ErrTransactionDoesntExist
	}

	_, storeErr := d.EventStoreWriter.StoreEvent(evt)
	if storeErr != nil {
		return nil, storeErr
	}

	return evt, nil
}
