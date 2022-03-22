package command

import (
	"product-service/cmd/product-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type UpdateProductCommand struct {
	Id          string
	Title       string
	Description string
	Quantity    int32
}

func (u *UpdateProductCommand) Handle() *eventlib.BaseEvent {
	return event.NewProductWasUpdatedEvent(u.Title, u.Description, u.Quantity, u.Id)
}
