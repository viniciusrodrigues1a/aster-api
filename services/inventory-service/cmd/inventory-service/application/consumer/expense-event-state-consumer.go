package consumer

import (
	"encoding/json"
	usecase "inventory-service/cmd/inventory-service/application/use-case"
	"inventory-service/cmd/inventory-service/external/messaging"
	"log"
)

type ExpenseEventStateConsumer struct {
	messaging *messaging.Messaging
	useCase   *usecase.AddExpenseToInventoryUseCase
}

func NewExpenseEventStateConsumer(m *messaging.Messaging, useCase *usecase.AddExpenseToInventoryUseCase) *ExpenseEventStateConsumer {
	return &ExpenseEventStateConsumer{
		messaging: m,
		useCase:   useCase,
	}
}

func (e *ExpenseEventStateConsumer) Consume() {
	ch, err := e.messaging.Connection.Channel()
	if err != nil {
		log.Fatalf("Couldn't open channel: %s", err)
	}

	q, err := ch.QueueDeclare(
		"expense-event-state",
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

			request := &usecase.AddExpenseToInventoryUseCaseRequest{}
			json.Unmarshal(m.Body, request)

			err := e.useCase.Execute(request)
			if err != nil {
				log.Printf("Error: `%s`", err.Error())
			}
		}
	}()

	log.Printf("Waiting for expense event state transfers.")
	<-forever
}
