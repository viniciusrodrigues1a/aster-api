package command

import (
	"expense-service/cmd/expense-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type DeleteExpenseCommand struct {
	Id string
}

func (d *DeleteExpenseCommand) Handle() *eventlib.BaseEvent {
	return event.NewExpenseWasDeletedEvent(d.Id)
}
