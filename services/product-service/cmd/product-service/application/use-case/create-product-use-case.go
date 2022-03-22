package usecase

import (
	"product-service/cmd/product-service/domain/command"
	"product-service/cmd/product-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type CreateProductUseCase struct {
	eventStoreRepository eventstorelib.EventStoreStreamWriter
	stateStoreRepository statestorelib.StateStoreWriter
}

func NewCreateProductUseCase(evtStore eventstorelib.EventStoreStreamWriter, sttStore statestorelib.StateStoreWriter) *CreateProductUseCase {
	return &CreateProductUseCase{
		eventStoreRepository: evtStore,
		stateStoreRepository: sttStore,
	}
}

type CreateProductUseCaseRequest struct {
	Title       string
	Description string
	Quantity    int32
}

func (c *CreateProductUseCase) Execute(request *CreateProductUseCaseRequest) error {
	command := command.CreateProductCommand{
		Title:       request.Title,
		Description: request.Description,
		Quantity:    request.Quantity,
	}
	event := command.Handle()

	id, err := c.eventStoreRepository.StoreEventStream(event)
	if err != nil {
		return err
	}

	projector := projector.ProductCreationProjector{}
	state := projector.Project(event)

	stateErr := c.stateStoreRepository.StoreState(id, state)
	if stateErr != nil {
		return stateErr
	}

	return nil
}
