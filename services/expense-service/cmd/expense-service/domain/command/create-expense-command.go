package command

import (
	"expense-service/cmd/expense-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type CreateExpenseCommand struct {
	Title       string
	Description string
	Value       int64
}

func (c *CreateExpenseCommand) Handle() *eventlib.BaseEvent {
	return event.NewExpenseWasCreatedEvent(c.Title, c.Description, c.Value)
}
