package command

import (
	"inventory-service/cmd/inventory-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateInventoryCommand struct {
	AccountId primitive.ObjectID
}

func NewCreateInventoryCommand(accountId primitive.ObjectID) *CreateInventoryCommand {
	return &CreateInventoryCommand{
		AccountId: accountId,
	}
}

func (c *CreateInventoryCommand) Handle() *eventlib.BaseEvent {
	return event.NewInventoryWasCreatedEvent(c.AccountId)
}
