package consumer

import (
	"encoding/json"
	"log"
	"transaction-service/cmd/transaction-service/external/messaging"

	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type ProductEventStateConsumer struct {
	messaging        *messaging.Messaging
	stateStoreWriter statestorelib.StateStoreWriter
}

func NewProductEventStateConsumer(m *messaging.Messaging, sttStoreW statestorelib.StateStoreWriter) *ProductEventStateConsumer {
	return &ProductEventStateConsumer{
		messaging:        m,
		stateStoreWriter: sttStoreW,
	}
}

type Product struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	SalePrice int64  `json:"sale_price"`
	Quantity  int64  `json:"quantity"`
	DeletedAt int64  `json:"deleted_at"`
}

func (p *ProductEventStateConsumer) Consume() {
	ch, err := p.messaging.Connection.Channel()
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
		log.Fatalf("Couldn't declare exchange: %s", err)
	}

	q, err := ch.QueueDeclare(
		"product-event-state-transaction-service",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Couldn't declare queue: %s", err)
	}

	bindKey := "product"
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

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for m := range msgs {
			log.Printf("Received message: `%s`\n", m.Body)

			product := &Product{}
			json.Unmarshal(m.Body, product)

			if product.DeletedAt > 0 {
				p.stateStoreWriter.StoreState(product.ID, nil)
			} else {
				p.stateStoreWriter.StoreState(product.ID, product)
			}
		}
	}()

	log.Printf("Waiting for product event state transfers.")
	<-forever
}
