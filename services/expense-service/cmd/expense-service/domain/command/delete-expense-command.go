package command

import (
	"errors"
	"expense-service/cmd/expense-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

var ErrExpenseDoesntExist = errors.New("expense doesn't exist")

type DeleteExpenseCommand struct {
	Id               string
	EventStoreWriter eventstorelib.EventStoreWriter
	StateStoreReader statestorelib.StateStoreReader
}

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
