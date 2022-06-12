package event

import (
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionWasUpdatedEvent struct {
	ProductID   *string
	Status      string
	Quantity    int64
	ValuePaid   int64
	Description string
}

func NewTransactionWasUpdatedEvent(productID *string, status string, quantity, valuePaid int64, description, id string) *eventlib.BaseEvent {
	payload := &TransactionWasUpdatedEvent{
		ProductID:   productID,
		Status:      status,
		Quantity:    quantity,
		ValuePaid:   valuePaid,
		Description: description,
	}

	oid, _ := primitive.ObjectIDFromHex(id)
	return eventlib.NewBaseEvent("transaction-was-updated", oid, payload)
}
