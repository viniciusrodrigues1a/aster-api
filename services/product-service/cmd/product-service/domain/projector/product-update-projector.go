package projector

import (
	"product-service/cmd/product-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type ProductUpdateProjector struct {
	CurrentState ProductState
}

func (p *ProductUpdateProjector) Project(e *eventlib.BaseEvent) *ProductState {
	payload := e.Payload.(*event.ProductWasUpdatedEvent)

	newState := &ProductState{
		Title:       payload.Title,
		Description: payload.Description,
		Quantity:    payload.Quantity,
		CreatedAt:   p.CurrentState.CreatedAt,
		DeletedAt:   p.CurrentState.DeletedAt,
	}

	return newState
}
