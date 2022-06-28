package event

import (
	"inventory-service/cmd/inventory-service/domain/dto"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductWasAddedToInventoryEvent struct {
	ProductID     string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Quantity      int64  `json:"quantity"`
	PurchasePrice int64  `json:"purchase_price"`
	SalePrice     int64  `json:"sale_price"`
	DeletedAt     int64  `json:"deleted_at"`
	Image         *dto.ProductImage
}

func NewProductWasAddedToInventoryEvent(productID, inventoryID, title, description string, quantity, purchasePrice, salePrice, deletedAt int64, image *dto.ProductImage) *eventlib.BaseEvent {
	payload := ProductWasAddedToInventoryEvent{
		ProductID:     productID,
		Title:         title,
		Description:   description,
		Quantity:      quantity,
		PurchasePrice: purchasePrice,
		SalePrice:     salePrice,
		DeletedAt:     deletedAt,
		Image:         image,
	}

	oid, _ := primitive.ObjectIDFromHex(inventoryID)
	return eventlib.NewBaseEvent("product-was-added-to-inventory", oid, payload)

}
