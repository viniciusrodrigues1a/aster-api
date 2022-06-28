package projector

import (
	"inventory-service/cmd/inventory-service/domain/dto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

type Product struct {
	ID            string            `json:"id"`
	Title         string            `json:"title"`
	Description   string            `json:"description"`
	Quantity      int64             `json:"quantity"`
	PurchasePrice int64             `json:"purchase_price"`
	SalePrice     int64             `json:"sale_price"`
	DeletedAt     int64             `json:"deleted_at"`
	Image         *dto.ProductImage `json:"image"`
}

type InventoryState struct {
	ID           string               `json:"id"`
	Participants []primitive.ObjectID `json:"participants"`
	Expenses     []Expense            `json:"expenses"`
	Transactions []Transaction        `json:"transactions"`
	Products     []Product            `json:"products"`
}
