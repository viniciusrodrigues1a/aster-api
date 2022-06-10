package command

import (
	"product-service/cmd/product-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type UpdateProductCommand struct {
	Id            string
	Title         string
	Description   string
	Quantity      int32
	PurchasePrice int64
	SalePrice     int64
}

func (u *UpdateProductCommand) Handle() *eventlib.BaseEvent {
	return event.NewProductWasUpdatedEvent(u.Title, u.Description, u.Quantity, u.PurchasePrice, u.SalePrice, u.Id)
}
