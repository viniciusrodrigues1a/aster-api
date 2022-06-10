package projector

import (
	"time"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type TransactionDeletionProjector struct {
	CurrentState *TransactionState
}

func (p *TransactionDeletionProjector) Project(e *eventlib.BaseEvent) *TransactionState {
	return &TransactionState{
		ProductID:   p.CurrentState.ProductID,
		ValuePaid:   p.CurrentState.ValuePaid,
		Description: p.CurrentState.Description,
		CreatedAt:   p.CurrentState.CreatedAt,
		DeletedAt:   time.Now().Unix(),
	}
}
