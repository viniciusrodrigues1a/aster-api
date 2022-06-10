package projector

import (
	"time"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type ExpenseDeletionProjector struct {
	CurrentState *ExpenseState
}

func (p *ExpenseDeletionProjector) Project(e *eventlib.BaseEvent) *ExpenseState {
	return &ExpenseState{
		ProductID:   p.CurrentState.ProductID,
		Title:       p.CurrentState.Title,
		Description: p.CurrentState.Description,
		Value:       p.CurrentState.Value,
		CreatedAt:   p.CurrentState.CreatedAt,
		DeletedAt:   time.Now().Unix(),
	}
}
