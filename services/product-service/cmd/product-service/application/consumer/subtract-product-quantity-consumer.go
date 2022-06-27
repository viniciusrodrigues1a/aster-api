package consumer

import (
	"encoding/json"
	"log"
	usecase "product-service/cmd/product-service/application/use-case"
	"product-service/cmd/product-service/external/messaging"
)

type SubtractProductQuantityCommandConsumer struct {
	messaging *messaging.Messaging
	useCase   *usecase.SubtractProductQuantityUseCase
}

func NewSubtractProductQuantityCommandConsumer(m *messaging.Messaging, u *usecase.SubtractProductQuantityUseCase) *SubtractProductQuantityCommandConsumer {
	return &SubtractProductQuantityCommandConsumer{
		messaging: m,
		useCase:   u,
	}
}

type Product struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	SalePrice int64  `json:"sale_price"`
	Quantity  int64  `json:"quantity"`
}

type consumerReceivingMessage struct {
	ProductID   string `json:"product_id"`
	AccountID   string `json:"account_id"`
	Description string `json:"description"`
	Quantity    int32  `json:"quantity"`
	ValuePaid   int64  `json:"value_paid"`
}

func (s *SubtractProductQuantityCommandConsumer) Consume() {
	ch, err := s.messaging.Connection.Channel()
	if err != nil {
		log.Fatalf("Couldn't open channel: %s", err)
	}

	exchangeName := "command.direct"
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
		"subtract-product-quantity-command",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Couldn't declare queue: %s", err)
	}

	bindKey := "subtract.product"
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

			body := &consumerReceivingMessage{}
			json.Unmarshal(m.Body, body)

			request := &usecase.SubtractProductQuantityUseCaseRequest{
				AccountID:  body.AccountID,
				ID:         body.ProductID,
				Reason:     "Automatically subtracting quantity because a Transaction was created",
				ByQuantity: body.Quantity,
			}

			err := s.useCase.Execute(request)
			if err != nil {
				log.Printf("Error: `%s`", err.Error())
			}
		}
	}()

	log.Printf("Waiting for product event state transfers.")
	<-forever
}
