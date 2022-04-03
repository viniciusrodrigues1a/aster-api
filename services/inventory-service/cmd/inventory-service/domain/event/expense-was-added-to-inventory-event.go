package event

import (
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpenseWasAddedToInventoryEvent struct {
	InventoryID string `json:"inventory_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       int64  `json:"value"`
}

func NewExpenseWasAddedToInventoryEvent(inventoryID, title, description string, value int64) *eventlib.BaseEvent {
	payload := ExpenseWasAddedToInventoryEvent{
		InventoryID: inventoryID,
		Title:       title,
		Description: description,
		Value:       value,
	}

	oid, _ := primitive.ObjectIDFromHex(inventoryID)
	return eventlib.NewBaseEvent("expense-was-added-to-inventory", oid, payload)
}
