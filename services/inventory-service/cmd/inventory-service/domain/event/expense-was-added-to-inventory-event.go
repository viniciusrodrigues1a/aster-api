package event

import (
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpenseWasAddedToInventoryEvent struct {
	TransactionID string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Value         int64  `json:"value"`
}

func NewExpenseWasAddedToInventoryEvent(transactionID, inventoryID, title, description string, value int64) *eventlib.BaseEvent {
	payload := ExpenseWasAddedToInventoryEvent{
		TransactionID: transactionID,
		Title:         title,
		Description:   description,
		Value:         value,
	}

	oid, _ := primitive.ObjectIDFromHex(inventoryID)
	return eventlib.NewBaseEvent("expense-was-added-to-inventory", oid, payload)
}
