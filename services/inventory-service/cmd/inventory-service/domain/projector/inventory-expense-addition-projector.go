package projector

import (
	"inventory-service/cmd/inventory-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type InventoryExpenseAdditionProjector struct {
	CurrentState *InventoryState
}

func (i *InventoryExpenseAdditionProjector) Project(e *eventlib.BaseEvent) *InventoryState {
	payload := e.Payload.(event.ExpenseWasAddedToInventoryEvent)

	newExpenses := append(i.CurrentState.Expenses, payload)

	return &InventoryState{
		ID:           i.CurrentState.ID,
		Participants: i.CurrentState.Participants,
		Expenses:     newExpenses,
		Transactions: i.CurrentState.Transactions,
	}
}
