package messaging

import (
	"encoding/json"
	"log"
	"product-service/cmd/product-service/domain/projector"

	"github.com/streadway/amqp"
)

type ProductEventStateEmitter struct {
	messaging *Messaging
}

func NewProductEventStateEmitter(m *Messaging) *ProductEventStateEmitter {
	return &ProductEventStateEmitter{messaging: m}
}

type ProductEventState struct {
	ID            string `json:"id"`
	AccountID     string `json:"account_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Quantity      int32  `json:"quantity"`
	PurchasePrice int64  `json:"purchase_price"`
	SalePrice     int64  `json:"sale_price"`
	CreatedAt     int64  `json:"created_at"`
	DeletedAt     int64  `json:"deleted_at,omitempty"`
}

func (e *ProductEventStateEmitter) Emit(state projector.ProductState, id, accountID string) {
	ch, err := e.messaging.Connection.Channel()
	if err != nil {
		log.Fatalf("Couldn't open channel: %s", err)
	}

	eventState := ProductEventState{
		ID:            id,
		AccountID:     accountID,
		Title:         state.Title,
		Description:   state.Description,
		Quantity:      state.Quantity,
		PurchasePrice: state.PurchasePrice,
		SalePrice:     state.SalePrice,
		CreatedAt:     state.CreatedAt,
		DeletedAt:     state.DeletedAt,
	}

	bytes, err := json.Marshal(eventState)
	if err != nil {
		log.Fatalf("Couldn't marshal message: %s", err)
	}

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        bytes,
	}

	ch.Publish("event-state-transfer.direct", "product", false, false, message)
}
