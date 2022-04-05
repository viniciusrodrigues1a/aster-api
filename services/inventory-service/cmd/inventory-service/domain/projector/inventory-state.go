package projector

import "go.mongodb.org/mongo-driver/bson/primitive"

type InventoryState struct {
	ID           string               `json:"id"`
	Participants []primitive.ObjectID `json:"participants"`
	Expenses     []interface{}        `json:"expenses"`
	Transactions []interface{}        `json:"transactions"`
}
