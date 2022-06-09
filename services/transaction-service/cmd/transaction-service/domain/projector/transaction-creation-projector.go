package projector

import (
	"transaction-service/cmd/transaction-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type TransactionCreationProjector struct{}

func (p *TransactionCreationProjector) Project(e *eventlib.BaseEvent) *TransactionState {
	payload := e.Payload.(*event.TransactionWasCreatedEvent)

	return &TransactionState{
		Quantity:    payload.Quantity,
		ValuePaid:   payload.ValuePaid,
		Description: payload.Description,
		CreatedAt:   e.Data.CreatedAt,
	}
}
