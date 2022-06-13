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

	index := findIndexOfTransactionByID(i.CurrentState.Transactions, payload.TransactionID)
	newTransactions := i.CurrentState.Transactions
	newTransaction := Transaction{
		ID:          payload.TransactionID,
		Description: payload.Description,
		ValuePaid:   payload.ValuePaid,
	}

	if index == -1 { // add a new Transaction to the slice
		newTransactions = append(newTransactions, newTransaction)
	} else { // update Transaction in the slice
		newTransactions[index] = newTransaction
	}

	if payload.DeletedAt > 0 && index > -1 { // remove Transaction from the slice
		newTransactions = append(newTransactions[:index], newTransactions[index+1:]...)
	}

	return &InventoryState{
		ID:           i.CurrentState.ID,
		Participants: i.CurrentState.Participants,
		Expenses:     i.CurrentState.Expenses,
		Transactions: newTransactions,
	}
}

func findIndexOfTransactionByID(transactions []Transaction, id string) int {
	for index, v := range transactions {
		if v.ID == id {
			return index
		}
	}

	return -1
}
