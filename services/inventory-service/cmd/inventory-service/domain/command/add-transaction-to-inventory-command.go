package command

import (
	"inventory-service/cmd/inventory-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
)

type Transaction struct {
	InventoryID   string
	TransactionID string
	Description   string
	ValuePaid     int64
	DeletedAt     int64
}

type AddTransactionToInventoryCommand struct {
	Transaction
	EventStoreWriter eventstorelib.EventStoreWriter
}

func NewAddTransactionToInventoryCommand(transactionID, inventoryID, description string, valuePaid, deletedAt int64, evtStore eventstorelib.EventStoreWriter) *AddTransactionToInventoryCommand {
	return &AddTransactionToInventoryCommand{
		Transaction: Transaction{
			TransactionID: transactionID,
			InventoryID:   inventoryID,
			Description:   description,
			ValuePaid:     valuePaid,
			DeletedAt:     deletedAt,
		},
		EventStoreWriter: evtStore,
	}
}

func (a *AddTransactionToInventoryCommand) Handle() (*eventlib.BaseEvent, error) {
	evt := event.NewTransactionWasAddedToInventoryEvent(a.TransactionID, a.InventoryID, a.Description, a.ValuePaid, a.DeletedAt)

	_, err := a.EventStoreWriter.StoreEvent(evt)
	if err != nil {
		return nil, err
	}

	return evt, nil
}
