package projector

import (
	"transaction-service/cmd/transaction-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type DebitMoneyToTransactionProjector struct {
	CurrentState *TransactionState
}

func (p *DebitMoneyToTransactionProjector) Project(e *eventlib.BaseEvent) *TransactionState {
	payload := e.Payload.(*event.MoneyWasDebitedToTransactionEvent)

	return &TransactionState{
		ProductID:   p.CurrentState.ProductID,
		Status:      payload.Status,
		Quantity:    p.CurrentState.Quantity,
		ValuePaid:   p.CurrentState.ValuePaid + payload.AmountDebited,
		Description: p.CurrentState.Description,
		CreatedAt:   p.CurrentState.CreatedAt,
	}
}
