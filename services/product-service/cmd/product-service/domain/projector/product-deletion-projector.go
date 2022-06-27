package projector

import (
	"time"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type ProductDeletionProjector struct {
	CurrentState *ProductState
}

func (p *ProductDeletionProjector) Project(e *eventlib.BaseEvent) *ProductState {
	return &ProductState{
		Title:         p.CurrentState.Title,
		Description:   p.CurrentState.Description,
		Quantity:      p.CurrentState.Quantity,
		PurchasePrice: p.CurrentState.PurchasePrice,
		SalePrice:     p.CurrentState.SalePrice,
		CreatedAt:     p.CurrentState.CreatedAt,
		DeletedAt:     time.Now().Unix(),
	}
}
