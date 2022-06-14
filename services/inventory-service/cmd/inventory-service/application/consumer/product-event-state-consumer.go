package consumer

import (
	"encoding/json"
	usecase "inventory-service/cmd/inventory-service/application/use-case"
	"inventory-service/cmd/inventory-service/external/messaging"
	"log"
)

type ProductEventStateConsumer struct {
	messaging *messaging.Messaging
	useCase   *usecase.AddProductToInventoryUseCase
}

func NewProductEventStateConsumer(m *messaging.Messaging, u *usecase.AddProductToInventoryUseCase) *ProductEventStateConsumer {
	return &ProductEventStateConsumer{
		messaging: m,
		useCase:   u,
	}
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
		"product-event-state-inventory-service",
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

			request := &usecase.AddProductToInventoryUseCaseRequest{}
			json.Unmarshal(m.Body, request)

			err := p.useCase.Execute(request)
			if err != nil {
				log.Printf("Error: `%s`", err.Error())
			}
		}
	}()

	log.Printf("Waiting for product event state transfers.")
	<-forever
}
