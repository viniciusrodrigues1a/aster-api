package projector

import (
	"product-service/cmd/product-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type ProductQuantityAdditionProjector struct {
	CurrentState ProductState
}

func (p *ProductQuantityAdditionProjector) Project(e *eventlib.BaseEvent) *ProductState {
	payload := e.Payload.(*event.ProductHadItsQuantityAddedEvent)

	newQuantity := p.CurrentState.Quantity + payload.ByQuantity

	return &ProductState{
		Title:         p.CurrentState.Title,
		Description:   p.CurrentState.Description,
		Quantity:      newQuantity,
		PurchasePrice: p.CurrentState.PurchasePrice,
		SalePrice:     p.CurrentState.SalePrice,
		CreatedAt:     p.CurrentState.CreatedAt,
		DeletedAt:     p.CurrentState.DeletedAt,
	}
}
