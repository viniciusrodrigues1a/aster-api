package command

import (
	"product-service/cmd/product-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type CreateProductCommand struct {
	Title         string
	Description   string
	Quantity      int32
	PurchasePrice int64
	SalePrice     int64
}

func (c *CreateProductCommand) Handle() *eventlib.BaseEvent {
	return event.NewProductWasCreatedEvent(c.Title, c.Description, c.Quantity, c.PurchasePrice, c.SalePrice)
}
