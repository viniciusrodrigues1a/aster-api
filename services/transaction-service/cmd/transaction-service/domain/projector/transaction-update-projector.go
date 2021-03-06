package projector

import (
	"transaction-service/cmd/transaction-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type TransactionUpdateProjector struct {
	CurrentState *TransactionState
}

func (p *TransactionUpdateProjector) Project(e *eventlib.BaseEvent) *TransactionState {
	payload := e.Payload.(*event.TransactionWasUpdatedEvent)

	productID := p.CurrentState.ProductID
	if payload.ProductID != nil {
		productID = payload.ProductID
	}

	status := p.CurrentState.Status
	if payload.Status != "" {
		status = payload.Status
	}

	return &TransactionState{
		ProductID:   productID,
		Status:      status,
		Quantity:    payload.Quantity,
		ValuePaid:   payload.ValuePaid,
		Description: payload.Description,
		CreatedAt:   p.CurrentState.CreatedAt,
	}
}
