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

	return &Messaging{
		Connection: conn,
	}
}

func (m *Messaging) Disconnect() {
	m.Connection.Close()
}
