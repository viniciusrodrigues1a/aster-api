package projector

import (
	"product-service/cmd/product-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type ProductQuantitySubtractionProjector struct {
	CurrentState ProductState
}

func (p *ProductQuantitySubtractionProjector) Project(e *eventlib.BaseEvent) *ProductState {
	payload := e.Payload.(*event.ProductHadItsQuantitySubtractedEvent)

	newQuantity := p.CurrentState.Quantity - payload.ByQuantity

	return &ProductState{
		Title:         p.CurrentState.Title,
		Description:   p.CurrentState.Description,
		Quantity:      newQuantity,
		Image:         p.CurrentState.Image,
		PurchasePrice: p.CurrentState.PurchasePrice,
		SalePrice:     p.CurrentState.SalePrice,
		CreatedAt:     p.CurrentState.CreatedAt,
		DeletedAt:     p.CurrentState.DeletedAt,
	}
}
