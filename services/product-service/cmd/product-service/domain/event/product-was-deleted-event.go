package event

import (
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewProductWasDeletedEvent(id string) *eventlib.BaseEvent {
	oid, _ := primitive.ObjectIDFromHex(id)
	return eventlib.NewBaseEvent("product-was-deleted", oid, struct{}{})
}
