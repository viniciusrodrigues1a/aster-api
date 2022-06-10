package event

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type ProductWasUpdatedEvent struct {
	Title         string `bson:"title"`
	Description   string `bson:"description"`
	Quantity      int32  `bson:"quantity"`
	PurchasePrice int64  `bson:"purchase_price"`
	SalePrice     int64  `bson:"sale_price"`
}

func NewProductWasUpdatedEvent(title, description string, quantity int32, purchase, sale int64, id string) *eventlib.BaseEvent {
	payload := &ProductWasUpdatedEvent{
		Title:         title,
		Description:   description,
		Quantity:      quantity,
		PurchasePrice: purchase,
		SalePrice:     sale,
	}

	oid, _ := primitive.ObjectIDFromHex(id)
	return eventlib.NewBaseEvent("product-was-updated", oid, payload)
}
