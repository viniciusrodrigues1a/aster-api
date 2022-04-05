package consumer

import (
	"encoding/json"
	usecase "inventory-service/cmd/inventory-service/application/use-case"
	"inventory-service/cmd/inventory-service/external/messaging"
	"log"
)

type TransactionEventStateConsumer struct {
	messaging *messaging.Messaging
	useCase   *usecase.AddTransactionToInventoryUseCase
}

func NewTransactionEventStateConsumer(m *messaging.Messaging, useCase *usecase.AddTransactionToInventoryUseCase) *TransactionEventStateConsumer {
	return &TransactionEventStateConsumer{
		messaging: m,
		useCase:   useCase,
	}
}

func (t *TransactionEventStateConsumer) Consume() {
	ch, err := t.messaging.Connection.Channel()
	if err != nil {
		log.Fatalf("Couldn't open channel: %s", err)
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

			request := &usecase.AddTransactionToInventoryUseCaseRequest{}
			json.Unmarshal(m.Body, request)

			err := t.useCase.Execute(request)
			if err != nil {
				log.Printf("Error: `%s`", err.Error())
			}
		}
	}()

	log.Printf("Waiting for transaction event state transfers.")
	<-forever
}
