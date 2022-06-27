package event

import (
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductHadItsQuantityAddedEvent struct {
	ByQuantity int32 `bson:"by_quantity"`
}

func NewProductHadItsQuantityAddedEvent(byQuantity int32, id string) *eventlib.BaseEvent {
	payload := &ProductHadItsQuantityAddedEvent{
		ByQuantity: byQuantity,
	}

	oid, _ := primitive.ObjectIDFromHex(id)
	return eventlib.NewBaseEvent("product-had-its-quantity-added", oid, payload)
}
