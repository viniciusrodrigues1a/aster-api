package event

import (
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionWasAddedToInventoryEvent struct {
	TransactionID string `json:"id"`
	Description   string `json:"description"`
	ValuePaid     int64  `json:"value_paid"`
}

func NewTransactionWasAddedToInventoryEvent(transactionID, inventoryID, description string, valuePaid int64) *eventlib.BaseEvent {
	payload := TransactionWasAddedToInventoryEvent{
		TransactionID: transactionID,
		Description:   description,
		ValuePaid:     valuePaid,
	}

	oid, _ := primitive.ObjectIDFromHex(inventoryID)
	return eventlib.NewBaseEvent("transaction-was-added-to-inventory", oid, payload)
}
