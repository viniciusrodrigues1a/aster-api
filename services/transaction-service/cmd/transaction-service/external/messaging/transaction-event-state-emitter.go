package messaging

import (
	"encoding/json"
	"log"
	"transaction-service/cmd/transaction-service/domain/projector"

	"github.com/streadway/amqp"
)

type ExpenseEventStateEmitter struct {
	messaging *Messaging
}

func NewExpenseEventStateEmitter(m *Messaging) *ExpenseEventStateEmitter {
	return &ExpenseEventStateEmitter{messaging: m}
}

type TransactionEventState struct {
	ID          string `json:"id"`
	AccountID   string `json:"account_id"`
	Description string `json:"description"`
	ValuePaid   int64  `json:"value_paid"`
	CreatedAt   int64  `json:"created_at"`
	DeletedAt   int64  `json:"deleted_at,omitempty"`
}

func (e *ExpenseEventStateEmitter) Emit(state projector.TransactionState, id string, accountID string) {
	ch, err := e.messaging.Connection.Channel()
	if err != nil {
		log.Fatalf("Couldn't open channel: %s", err)
	}

	eventState := TransactionEventState{
		ID:          id,
		AccountID:   accountID,
		Description: state.Description,
		ValuePaid:   state.ValuePaid,
		CreatedAt:   state.CreatedAt,
		DeletedAt:   state.DeletedAt,
	}

	bytes, err := json.Marshal(eventState)
	if err != nil {
		log.Fatalf("Couldn't marshal message: %s", err)
	}

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        bytes,
	}

	ch.Publish("event-state-transfer.direct", "transaction", false, false, message)
}
