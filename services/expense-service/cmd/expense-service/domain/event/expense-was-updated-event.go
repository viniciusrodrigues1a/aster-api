package event

import (
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpenseWasUpdatedEvent struct {
	ProductID   *string
	Title       string
	Description string
	Value       int64
}

func NewExpenseWasUpdatedEvent(productID *string, title, description string, value int64, id string) *eventlib.BaseEvent {
	payload := &ExpenseWasUpdatedEvent{
		ProductID:   productID,
		Title:       title,
		Description: description,
		Value:       value,
	}

	oid, _ := primitive.ObjectIDFromHex(id)
	return eventlib.NewBaseEvent("expense-was-updated", oid, payload)
}
