package consumer

import (
	"encoding/json"
	usecase "inventory-service/cmd/inventory-service/application/use-case"
	"inventory-service/cmd/inventory-service/external/messaging"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountCreationConsumer struct {
	messaging *messaging.Messaging
	useCase   *usecase.CreateInventoryUseCase
}

func NewAccountCreationConsumer(m *messaging.Messaging, useCase *usecase.CreateInventoryUseCase) *AccountCreationConsumer {
	return &AccountCreationConsumer{
		messaging: m,
		useCase:   useCase,
	}
}

type accountCreationConsumerMessage struct {
	Id primitive.ObjectID
}

func (n *AccountCreationConsumer) Consume() {
	ch, err := n.messaging.Connection.Channel()
	if err != nil {
		log.Fatalf("Couldn't open channel: %s", err)
	}

	q, err := ch.QueueDeclare(
		"create-inventory-command",
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
			log.Printf("Received message: `%s`", m.Body)

			var message accountCreationConsumerMessage
			json.Unmarshal(m.Body, &message)

			request := &usecase.CreateInventoryUseCaseRequest{AccountId: message.Id}
			err := n.useCase.Execute(request)
			if err != nil {
				log.Printf("Error: `%s`", err.Error())
			}
		}
	}()

	log.Printf("Waiting for messages.")
	<-forever
}
