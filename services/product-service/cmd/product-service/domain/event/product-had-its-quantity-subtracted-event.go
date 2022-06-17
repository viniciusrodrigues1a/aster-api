package event

import (
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductHadItsQuantitySubtractedEvent struct {
	Reason     string `bson:"reason"`
	ByQuantity int32  `bson:"by_quantity"`
}

func NewProductHadItsQuantitySubtractedEvent(reason string, byQuantity int32, id string) *eventlib.BaseEvent {
	payload := &ProductHadItsQuantitySubtractedEvent{
		Reason:     reason,
		ByQuantity: byQuantity,
	}

	oid, _ := primitive.ObjectIDFromHex(id)
	return eventlib.NewBaseEvent("product-had-its-quantity-subtracted", oid, payload)
}
