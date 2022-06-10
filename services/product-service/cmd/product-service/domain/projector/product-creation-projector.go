package projector

import (
	"product-service/cmd/product-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type ProductCreationProjector struct{}

func (p *ProductCreationProjector) Project(e *eventlib.BaseEvent) *ProductState {
	payload := e.Payload.(*event.ProductWasCreatedEvent)

	return &ProductState{
		Title:         payload.Title,
		Description:   payload.Description,
		Quantity:      payload.Quantity,
		PurchasePrice: payload.PurchasePrice,
		SalePrice:     payload.SalePrice,
		CreatedAt:     e.Data.CreatedAt,
		DeletedAt:     e.Data.DeletedAt,
	}
}
