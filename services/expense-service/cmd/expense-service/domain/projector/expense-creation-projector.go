package projector

import (
	"expense-service/cmd/expense-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type ExpenseCreationProjector struct{}

func (p *ExpenseCreationProjector) Project(e *eventlib.BaseEvent) *ExpenseState {
	payload := e.Payload.(*event.ExpenseWasCreatedEvent)

	return &ExpenseState{
		Title:       payload.Title,
		Description: payload.Description,
		Value:       payload.Value,
	}
}
