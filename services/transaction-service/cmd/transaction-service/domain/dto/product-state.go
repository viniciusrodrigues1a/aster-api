package dto

type ProductState struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	SalePrice int64  `json:"sale_price"`
}
