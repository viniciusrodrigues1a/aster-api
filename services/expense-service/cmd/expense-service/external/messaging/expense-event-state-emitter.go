package messaging

import (
	"encoding/json"
	"expense-service/cmd/expense-service/domain/projector"
	"log"

	"github.com/streadway/amqp"
)

type ExpenseEventStateEmitter struct {
	messaging *Messaging
}

func NewExpenseEventStateEmitter(m *Messaging) *ExpenseEventStateEmitter {
	return &ExpenseEventStateEmitter{messaging: m}
}

type ExpenseEventState struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       int64  `json:"value"`
	CreatedAt   int64  `json:"created_at"`
	DeletedAt   int64  `json:"deleted_at,omitempty"`
}

func (e *ExpenseEventStateEmitter) Emit(state projector.ExpenseState, id string) {
	ch, err := e.messaging.Connection.Channel()
	if err != nil {
		log.Fatalf("Couldn't open channel: %s", err)
	}

	eventState := ExpenseEventState{
		Id:          id,
		Title:       state.Title,
		Description: state.Description,
		Value:       state.Value,
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

	ch.Publish("event-state-transfer.direct", "expense", false, false, message)
}
