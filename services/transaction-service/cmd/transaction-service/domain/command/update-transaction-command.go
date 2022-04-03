package command

import (
	"fmt"
	"transaction-service/cmd/transaction-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type UpdateTransactionCommand struct {
	ID               string
	ValuePaid        int64
	Description      string
	EventStoreWriter eventstorelib.EventStoreWriter
	StateStoreReader statestorelib.StateStoreReader
}

// Handle stores a TransactionWasUpdatedEvent to the event store and returns the resulting event.
// Returns ErrTransactionDoesntExist if it can't read the state from the state store
func (u *UpdateTransactionCommand) Handle() (*eventlib.BaseEvent, error) {
	evt := event.NewTransactionWasUpdatedEvent(u.ValuePaid, u.Description, u.ID)

	fmt.Println(u.ValuePaid)
	fmt.Println(u.Description)

	_, err := u.StateStoreReader.ReadState(u.ID)
	if err != nil {
		return nil, ErrTransactionDoesntExist
	}

	_, storeErr := u.EventStoreWriter.StoreEvent(evt)
	if storeErr != nil {
		return nil, storeErr
	}

	return evt, nil
}
