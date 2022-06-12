package dto

type ProductState struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	PurchasePrice int64  `json:"purchase_price"`
}
