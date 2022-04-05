package command

import (
	"inventory-service/cmd/inventory-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
)

type Expense struct {
	TransactionID string
	InventoryID   string
	Title         string
	Description   string
	Value         int64
}

type AddExpenseToInventoryCommand struct {
	Expense
	EventStoreWriter eventstorelib.EventStoreWriter
}

func NewAddExpenseToInventoryCommand(transactionID, inventoryID, title, description string, value int64, evtStore eventstorelib.EventStoreWriter) *AddExpenseToInventoryCommand {
	return &AddExpenseToInventoryCommand{
		Expense: Expense{
			TransactionID: transactionID,
			InventoryID:   inventoryID,
			Title:         title,
			Description:   description,
			Value:         value,
		},
		EventStoreWriter: evtStore,
	}
}

func (a *AddExpenseToInventoryCommand) Handle() (*eventlib.BaseEvent, error) {
	evt := event.NewExpenseWasAddedToInventoryEvent(a.TransactionID, a.InventoryID, a.Title, a.Description, a.Value)

	_, err := a.EventStoreWriter.StoreEvent(evt)
	if err != nil {
		return nil, err
	}

	return evt, nil
}
