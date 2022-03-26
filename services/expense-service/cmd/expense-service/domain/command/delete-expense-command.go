package command

import (
	"expense-service/cmd/expense-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type DeleteExpenseCommand struct {
	Id               string
	EventStoreWriter eventstorelib.EventStoreWriter
	StateStoreReader statestorelib.StateStoreReader
}

// Handle stores an ExpenseWasDeletedEvent to the event store and returns the resulting event.
// returns ErrExpenseDoesntExist if it can't read the expense state from the state store.
func (d *DeleteExpenseCommand) Handle() (*eventlib.BaseEvent, error) {
	evt := event.NewExpenseWasDeletedEvent(d.Id)

	_, err := d.StateStoreReader.ReadState(d.Id)
	if err != nil {
		return nil, ErrExpenseDoesntExist
	}

	_, storeErr := d.EventStoreWriter.StoreEvent(evt)
	if storeErr != nil {
		return nil, storeErr
	}

	return evt, nil
}
