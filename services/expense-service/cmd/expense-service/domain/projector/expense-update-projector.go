package projector

import (
	"expense-service/cmd/expense-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type ExpenseUpdateProjector struct {
	CurrentState *ExpenseState
}

func (p *ExpenseUpdateProjector) Project(e *eventlib.BaseEvent) *ExpenseState {
	payload := e.Payload.(*event.ExpenseWasUpdatedEvent)

	var productID *string = p.CurrentState.ProductID
	if productID != nil {
		productID = payload.ProductID
	}

	return &ExpenseState{
		ProductID:   productID,
		Title:       payload.Title,
		Description: payload.Description,
		Value:       payload.Value,
		CreatedAt:   p.CurrentState.CreatedAt,
	}
}
