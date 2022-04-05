package projector

import (
	"inventory-service/cmd/inventory-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type InventoryTransactionAdditionProjector struct {
	CurrentState *InventoryState
}

func (i *InventoryTransactionAdditionProjector) Project(e *eventlib.BaseEvent) *InventoryState {
	payload := e.Payload.(event.TransactionWasAddedToInventoryEvent)

	newTransactions := append(i.CurrentState.Transactions, payload)

	return &InventoryState{
		ID:           i.CurrentState.ID,
		Participants: i.CurrentState.Participants,
		Expenses:     i.CurrentState.Expenses,
		Transactions: newTransactions,
	}
}
