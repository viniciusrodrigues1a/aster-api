package event

import (
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpenseWasAddedToInventoryEvent struct {
	ExpenseID   string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       int64  `json:"value"`
	DeletedAt   int64  `json:"deleted_at"`
}

func NewExpenseWasAddedToInventoryEvent(expenseID, inventoryID, title, description string, value, deletedAt int64) *eventlib.BaseEvent {
	payload := ExpenseWasAddedToInventoryEvent{
		ExpenseID:   expenseID,
		Title:       title,
		Description: description,
		Value:       value,
		DeletedAt:   deletedAt,
	}

	oid, _ := primitive.ObjectIDFromHex(inventoryID)
	return eventlib.NewBaseEvent("expense-was-added-to-inventory", oid, payload)
}
