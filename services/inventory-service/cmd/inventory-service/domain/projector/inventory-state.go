package projector

import "go.mongodb.org/mongo-driver/bson/primitive"

type InventoryState struct {
	ID           string
	Participants []primitive.ObjectID
	Expenses     []interface{}
	Transactions []interface{}
}
