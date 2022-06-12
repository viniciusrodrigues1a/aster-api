package event

import (
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MoneyWasDebitedToTransactionEvent struct {
	AmountDebited int64  `bson:"amount_debited"`
	Status        string `bson:"status"`
}

func NewMoneyWasDebitedToTransactionEvent(amountDebited int64, status, id string) *eventlib.BaseEvent {
	payload := &MoneyWasDebitedToTransactionEvent{
		AmountDebited: amountDebited,
		Status:        status,
	}

	oid, _ := primitive.ObjectIDFromHex(id)
	return eventlib.NewBaseEvent("money-was-debited-to-transaction", oid, payload)
}
