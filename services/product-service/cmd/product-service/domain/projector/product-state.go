package projector

import "product-service/cmd/product-service/domain/dto"

type ProductState struct {
	Title         string
	Description   string
	Quantity      int32
	PurchasePrice int64
	SalePrice     int64
	Image         *dto.ProductImage
	CreatedAt     int64
	DeletedAt     int64
}
