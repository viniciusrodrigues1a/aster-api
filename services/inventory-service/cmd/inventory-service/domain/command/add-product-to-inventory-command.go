package command

import (
	"inventory-service/cmd/inventory-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
)

type Product struct {
	InventoryID   string
	ProductID     string
	Title         string
	Description   string
	Quantity      int64
	PurchasePrice int64
	SalePrice     int64
	DeletedAt     int64
}

type AddProductToInventoryCommand struct {
	Product
	EventStoreWriter eventstorelib.EventStoreWriter
}

func NewAddProductToInventoryCommand(productID, inventoryID, title, description string, quantity, purchasePrice, salePrice, deletedAt int64, evtStoreW eventstorelib.EventStoreWriter) *AddProductToInventoryCommand {
	return &AddProductToInventoryCommand{
		Product: Product{
			ProductID:     productID,
			InventoryID:   inventoryID,
			Title:         title,
			Description:   description,
			Quantity:      quantity,
			PurchasePrice: purchasePrice,
			SalePrice:     salePrice,
			DeletedAt:     deletedAt,
		},
		EventStoreWriter: evtStoreW,
	}
}

func (a *AddProductToInventoryCommand) Handle() (*eventlib.BaseEvent, error) {
	evt := event.NewProductWasAddedToInventoryEvent(a.ProductID, a.InventoryID, a.Title, a.Description, a.Quantity, a.PurchasePrice, a.SalePrice, a.DeletedAt)

	_, err := a.EventStoreWriter.StoreEvent(evt)
	if err != nil {
		return nil, err
	}

	return evt, nil
}
