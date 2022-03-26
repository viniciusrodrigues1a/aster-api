package command

import (
	"expense-service/cmd/expense-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type UpdateExpenseCommand struct {
	Id               string
	Title            string
	Description      string
	Value            int64
	EventStoreWriter eventstorelib.EventStoreWriter
	StateStoreReader statestorelib.StateStoreReader
}

func (u *UpdateExpenseCommand) Handle() (*eventlib.BaseEvent, error) {
	evt := event.NewExpenseWasUpdatedEvent(u.Title, u.Description, u.Value, u.Id)

	_, err := u.StateStoreReader.ReadState(u.Id)
	if err != nil {
		return nil, ErrExpenseDoesntExist
	}

	_, storeErr := u.EventStoreWriter.StoreEvent(evt)
	if storeErr != nil {
		return nil, storeErr
	}

	return evt, nil
}
