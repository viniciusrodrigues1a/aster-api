package event

import (
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpenseWasCreatedEvent struct {
	ProductID   *string
	Title       string
	Description string
	Value       int64
}

func NewExpenseWasCreatedEvent(productID *string, title, description string, value int64) *eventlib.BaseEvent {
	payload := &ExpenseWasCreatedEvent{
		ProductID:   productID,
		Title:       title,
		Description: description,
		Value:       value,
	}

	return eventlib.NewBaseEvent("expense-was-created", primitive.NewObjectID(), payload)
}
