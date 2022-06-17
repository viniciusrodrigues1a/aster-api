package messaging

import (
	"log"

	"github.com/streadway/amqp"
)

type Messaging struct {
	Connection *amqp.Connection
}

func New() *Messaging {
	conn, err := amqp.Dial("amqp://aster:pa55@localhost:5672/")
	if err != nil {
		log.Fatalf("Couldn't connect to RabbitMQ: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Couldn't open channel: %s", err)
	}

	exchangeName := "event-state-transfer.direct"
	exchangeErr := ch.ExchangeDeclare(
		exchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if exchangeErr != nil {
		log.Fatalf("Couldn't declare exchange: %s", exchangeErr)
	}

	q, err := ch.QueueDeclare(
		"transaction-event-state",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Couldn't declare queue: %s", err)
	}

	bindKey := "transaction"
	bindErr := ch.QueueBind(
		q.Name,
		bindKey,
		exchangeName,
		false,
		nil,
	)
	if bindErr != nil {
		log.Fatalf("Couldn't declare bind: %s", err)
	}

	return &Messaging{
		Connection: conn,
	}
}

func (m *Messaging) Disconnect() {
	m.Connection.Close()
}
