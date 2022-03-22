package projector

import "go.mongodb.org/mongo-driver/bson/primitive"

type InventoryState struct {
	AccountId    primitive.ObjectID
	Participants []primitive.ObjectID
}
