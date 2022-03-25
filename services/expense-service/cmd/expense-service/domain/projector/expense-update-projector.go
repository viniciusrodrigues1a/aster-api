package projector

import (
	"expense-service/cmd/expense-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type ExpenseUpdateProjector struct{}

func (p *ExpenseUpdateProjector) Project(e *eventlib.BaseEvent) *ExpenseState {
	payload := e.Payload.(*event.ExpenseWasUpdatedEvent)

	return &ExpenseState{
		Title:       payload.Title,
		Description: payload.Description,
		Value:       payload.Value,
	}
}
