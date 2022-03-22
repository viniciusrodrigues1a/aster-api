package event

import (
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpenseWasCreatedEvent struct {
	Title       string
	Description string
	Value       int64
}

func NewExpenseWasCreatedEvent(title, description string, value int64) *eventlib.BaseEvent {
	payload := &ExpenseWasCreatedEvent{
		Title:       title,
		Description: description,
		Value:       value,
	}

	return eventlib.NewBaseEvent("expense-was-created", primitive.NewObjectID(), payload)
}
