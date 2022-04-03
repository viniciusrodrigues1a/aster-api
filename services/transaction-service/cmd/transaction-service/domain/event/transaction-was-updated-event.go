package event

import (
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionWasUpdatedEvent struct {
	ValuePaid   int64
	Description string
}

func NewTransactionWasUpdatedEvent(valuePaid int64, description, id string) *eventlib.BaseEvent {
	payload := &TransactionWasUpdatedEvent{
		ValuePaid:   valuePaid,
		Description: description,
	}

	oid, _ := primitive.ObjectIDFromHex(id)
	return eventlib.NewBaseEvent("transaction-was-updated", oid, payload)
}
