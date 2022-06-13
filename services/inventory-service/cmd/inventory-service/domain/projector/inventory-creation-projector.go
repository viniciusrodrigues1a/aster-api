package projector

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryCreationProjector struct{}

func (i *InventoryCreationProjector) Project(id string) *InventoryState {

	return &InventoryState{
		ID:           id,
		Participants: []primitive.ObjectID{},
		Expenses:     []Expense{},
		Transactions: []Transaction{},
	}
}
