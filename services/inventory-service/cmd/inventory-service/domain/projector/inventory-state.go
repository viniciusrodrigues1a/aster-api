package projector

import "go.mongodb.org/mongo-driver/bson/primitive"

type Expense struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       int64  `json:"value"`
}

type Transaction struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	ValuePaid   int64  `json:"value_paid"`
}

type InventoryState struct {
	ID           string               `json:"id"`
	Participants []primitive.ObjectID `json:"participants"`
	Expenses     []Expense            `json:"expenses"`
	Transactions []Transaction        `json:"transactions"`
}
