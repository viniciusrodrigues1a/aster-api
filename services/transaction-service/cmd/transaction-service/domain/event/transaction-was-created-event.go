package event

import (
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionWasCreatedEvent struct {
	ProductID   *string
	Status      string
	Quantity    int64
	ValuePaid   int64
	TotalValue  int64
	Description string
}

func NewTransactionWasCreatedEvent(productID *string, status string, quantity, valuePaid int64, description string) *eventlib.BaseEvent {
	payload := &TransactionWasCreatedEvent{
		ProductID:   productID,
		Status:      status,
		Quantity:    quantity,
		ValuePaid:   valuePaid,
		Description: description,
	}

	return eventlib.NewBaseEvent("transaction-was-created", primitive.NewObjectID(), payload)
}
