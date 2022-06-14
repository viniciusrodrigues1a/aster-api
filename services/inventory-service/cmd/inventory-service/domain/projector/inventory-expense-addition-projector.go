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

	index := findIndexOfExpenseByID(i.CurrentState.Expenses, payload.ExpenseID)
	newExpenses := i.CurrentState.Expenses
	newExpense := Expense{
		ID:          payload.ExpenseID,
		Title:       payload.Title,
		Description: payload.Description,
		Value:       payload.Value,
	}

	if index == -1 { // add a new Expense to the slice
		newExpenses = append(newExpenses, newExpense)
	} else { // update Expense in the slice
		newExpenses[index] = newExpense
	}

	if payload.DeletedAt > 0 && index > -1 { // remove Expense from the slice
		newExpenses = append(newExpenses[:index], newExpenses[index+1:]...)
	}

	return &InventoryState{
		ID:           i.CurrentState.ID,
		Participants: i.CurrentState.Participants,
		Expenses:     newExpenses,
		Transactions: i.CurrentState.Transactions,
		Products:     i.CurrentState.Products,
	}
}

func findIndexOfExpenseByID(expenses []Expense, id string) int {
	for index, v := range expenses {
		if v.ID == id {
			return index
		}
	}

	return -1
}
