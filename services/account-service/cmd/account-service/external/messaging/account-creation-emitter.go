package messaging

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type AccountCreationEmitter struct {
	messaging *Messaging
}

func NewAccountCreationEmitter(m *Messaging) *AccountCreationEmitter {
	return &AccountCreationEmitter{
		messaging: m,
	}
}

func (a *AccountCreationEmitter) Emit(body interface{}) {
	ch, err := a.messaging.Connection.Channel()
	if err != nil {
		log.Fatalf("Couldn't open channel: %s", err)
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("Couldn't marshal message: %s", err)
	}

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        bytes,
	}

	ch.Publish("command.direct", "create.inventory", false, false, message)
}
