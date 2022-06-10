package command

import (
	"encoding/json"
	"expense-service/cmd/expense-service/domain/event"
	"expense-service/cmd/expense-service/domain/projector"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type UpdateExpenseCommand struct {
	ProductID               *string
	ID                      string
	Title                   string
	Description             string
	Value                   int64
	EventStoreWriter        eventstorelib.EventStoreWriter
	StateStoreReader        statestorelib.StateStoreReader
	ProductStateStoreReader statestorelib.StateStoreReader
}

// Handle stores an ExpenseWasUpdatedEvent to the event store and returns the resulting event.
// returns ErrExpenseDoesntExist if it can't read the expense state from the state store.
func (u *UpdateExpenseCommand) Handle() (*eventlib.BaseEvent, error) {
	stateString, err := u.StateStoreReader.ReadState(u.ID)
	if err != nil {
		return nil, ErrExpenseDoesntExist
	}

	expense := projector.ExpenseState{}
	json.Unmarshal([]byte(stateString), &expense)

	if expense.DeletedAt > 0 {
		return nil, ErrExpenseDoesntExist
	}

	title := u.Title
	if u.Title == "" {
		title = expense.Title
	}

	if u.ProductID != nil {
		_, productStateErr := u.ProductStateStoreReader.ReadState(*u.ProductID)
		if productStateErr != nil {
			return nil, ErrProductCouldntBeFound
		}
	}

	evt := event.NewExpenseWasUpdatedEvent(u.ProductID, title, u.Description, u.Value, u.ID)

	_, storeErr := u.EventStoreWriter.StoreEvent(evt)
	if storeErr != nil {
		return nil, storeErr
	}

	return evt, nil
}
