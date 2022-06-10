package command

import (
	"encoding/json"
	"transaction-service/cmd/transaction-service/domain/event"
	"transaction-service/cmd/transaction-service/domain/projector"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type UpdateTransactionCommand struct {
	ProductID               *string
	ID                      string
	Quantity                int64
	ValuePaid               int64
	Description             string
	EventStoreWriter        eventstorelib.EventStoreWriter
	StateStoreReader        statestorelib.StateStoreReader
	ProductStateStoreReader statestorelib.StateStoreReader
}

// Handle stores a TransactionWasUpdatedEvent to the event store and returns the resulting event.
// Returns ErrTransactionDoesntExist if it can't read the state from the state store
func (u *UpdateTransactionCommand) Handle() (*eventlib.BaseEvent, error) {
	stateString, err := u.StateStoreReader.ReadState(u.ID)
	if err != nil {
		return nil, ErrTransactionDoesntExist
	}

	transaction := projector.TransactionState{}
	json.Unmarshal([]byte(stateString), &transaction)

	if transaction.DeletedAt > 0 {
		return nil, ErrTransactionDoesntExist
	}

	quantity := u.Quantity
	if u.Quantity <= 0 {
		quantity = transaction.Quantity
	}

	if err != nil {
		_, err := u.ProductStateStoreReader.ReadState(*u.ProductID)
		if err != nil {
			return nil, ErrProductCouldntBeFound
		}
	}

	evt := event.NewTransactionWasUpdatedEvent(u.ProductID, quantity, u.ValuePaid, u.Description, u.ID)

	_, storeErr := u.EventStoreWriter.StoreEvent(evt)
	if storeErr != nil {
		return nil, storeErr
	}

	return evt, nil
}
