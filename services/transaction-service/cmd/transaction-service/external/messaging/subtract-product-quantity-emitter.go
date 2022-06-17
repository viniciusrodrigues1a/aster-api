package messaging

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type SubtractProductQuantityEmitter struct {
	messaging *Messaging
}

func NewSubtractProductQuantityEmitter(m *Messaging) *SubtractProductQuantityEmitter {
	return &SubtractProductQuantityEmitter{
		messaging: m,
	}
}

func (a *SubtractProductQuantityEmitter) Emit(body interface{}) {
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

	ch.Publish("command.direct", "subtract.product", false, false, message)
}
