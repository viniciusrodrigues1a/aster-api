package usecase

import (
	"inventory-service/cmd/inventory-service/domain/command"
	"inventory-service/cmd/inventory-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateInventoryUseCase struct {
	eventStoreStreamWriter eventstorelib.EventStoreStreamWriter
	stateStoreWriter       statestorelib.StateStoreWriter
}

func NewCreateInventoryUseCase(evtStore eventstorelib.EventStoreStreamWriter, sttStore statestorelib.StateStoreWriter) *CreateInventoryUseCase {
	return &CreateInventoryUseCase{
		eventStoreStreamWriter: evtStore,
		stateStoreWriter:       sttStore,
	}
}

type CreateInventoryUseCaseRequest struct {
	AccountId primitive.ObjectID
}

func (c *CreateInventoryUseCase) Execute(request *CreateInventoryUseCaseRequest) error {
	command := command.NewCreateInventoryCommand(request.AccountId, c.eventStoreStreamWriter)
	event, err := command.Handle()
	if err != nil {
		return err
	}

	projector := projector.InventoryCreationProjector{}
	state := projector.Project(event)

	stateErr := c.stateStoreWriter.StoreState(event.Data.StreamId.Hex(), state)
	if stateErr != nil {
		return stateErr
	}

	return nil
}
