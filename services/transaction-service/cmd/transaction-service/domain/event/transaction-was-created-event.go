package event

import (
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionWasCreatedEvent struct {
	ValuePaid   int64
	Description string
}

func NewTransactionWasCreatedEvent(valuePaid int64, description string) *eventlib.BaseEvent {
	payload := &TransactionWasCreatedEvent{
		ValuePaid:   valuePaid,
		Description: description,
	}

	return eventlib.NewBaseEvent("transaction-was-created", primitive.NewObjectID(), payload)
}
