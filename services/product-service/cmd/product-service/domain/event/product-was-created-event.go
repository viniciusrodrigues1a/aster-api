package event

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type ProductWasCreatedEvent struct {
	Title         string `bson:"title"`
	Description   string `bson:"description"`
	Quantity      int32  `bson:"quantity"`
	PurchasePrice int64  `bson:"purchase_price"`
	SalePrice     int64  `bson:"sale_price"`
}

func NewProductWasCreatedEvent(title, description string, quantity int32, purchase, sale int64) *eventlib.BaseEvent {
	payload := &ProductWasCreatedEvent{
		Title:         title,
		Description:   description,
		Quantity:      quantity,
		PurchasePrice: purchase,
		SalePrice:     sale,
	}

	return eventlib.NewBaseEvent("product-was-created", primitive.NewObjectID(), payload)
}
