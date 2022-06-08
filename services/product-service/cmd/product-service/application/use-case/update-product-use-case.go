package usecase

import (
	"encoding/json"
	"product-service/cmd/product-service/domain/command"
	"product-service/cmd/product-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type StateStoreRepository interface {
	statestorelib.StateStoreWriter
	statestorelib.StateStoreReader
}

type UpdateProductUseCase struct {
	eventStoreRepository eventstorelib.EventStoreWriter
	stateStoreRepository StateStoreRepository
}

func NewUpdateProductUseCase(evtStore eventstorelib.EventStoreWriter, sttStore StateStoreRepository) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		eventStoreRepository: evtStore,
		stateStoreRepository: sttStore,
	}
}

type UpdateProductUseCaseRequest struct {
	Id          string
	Title       string
	Description string
	Quantity    int32
}

func (u *UpdateProductUseCase) Execute(request *UpdateProductUseCaseRequest) error {
	command := command.UpdateProductCommand{
		Title:       request.Title,
		Description: request.Description,
		Quantity:    request.Quantity,
		Id:          request.Id,
	}
	event := command.Handle()

	val, err := u.stateStoreRepository.ReadState(request.Id)
	if err != nil {
		return err
	}

	id, err := u.eventStoreRepository.StoreEvent(event)
	if err != nil {
		return err
	}

	currentState := projector.ProductState{}
	json.Unmarshal([]byte(val.(string)), &currentState)
	projector := projector.ProductUpdateProjector{CurrentState: currentState}
	state := projector.Project(event)

	stateErr := u.stateStoreRepository.StoreState(id, state)
	if stateErr != nil {
		return stateErr
	}

	return nil
}
