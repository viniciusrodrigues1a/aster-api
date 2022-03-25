package command

import (
	"expense-service/cmd/expense-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type UpdateExpenseCommand struct {
	Id          string
	Title       string
	Description string
	Value       int64
}

func (u *UpdateExpenseCommand) Handle() *eventlib.BaseEvent {
	return event.NewExpenseWasUpdatedEvent(u.Title, u.Description, u.Value, u.Id)
}
