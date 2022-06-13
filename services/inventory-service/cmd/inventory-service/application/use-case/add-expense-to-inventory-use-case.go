package usecase

import (
	"encoding/json"
	"errors"
	"inventory-service/cmd/inventory-service/domain/command"
	"inventory-service/cmd/inventory-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

var ErrInventoryDoesntExist = errors.New("Inventory doesn't exist")

type AddExpenseToInventoryUseCase struct {
	eventStoreWriter eventstorelib.EventStoreWriter
	stateStoreReader statestorelib.StateStoreReader
	stateStoreWriter statestorelib.StateStoreWriter
}

func NewAddExpenseToInventoryUseCase(
	evtStore eventstorelib.EventStoreWriter,
	sttStoreR statestorelib.StateStoreReader,
	sttStoreW statestorelib.StateStoreWriter,
) *AddExpenseToInventoryUseCase {
	return &AddExpenseToInventoryUseCase{
		eventStoreWriter: evtStore,
		stateStoreReader: sttStoreR,
		stateStoreWriter: sttStoreW,
	}
}

type AddExpenseToInventoryUseCaseRequest struct {
	ID          string `json:"id"`
	Email       string `json:"account_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       int64  `json:"value"`
	DeletedAt   int64  `json:"deleted_at"`
}

func (a *AddExpenseToInventoryUseCase) Execute(request *AddExpenseToInventoryUseCaseRequest) error {
	val, err := a.stateStoreReader.ReadState(request.Email)
	if err != nil {
		return ErrInventoryDoesntExist
	}
	currentState := &projector.InventoryState{}
	json.Unmarshal([]byte(val), &currentState)

	command := command.NewAddExpenseToInventoryCommand(request.ID, currentState.ID, request.Title, request.Description, request.Value, request.DeletedAt, a.eventStoreWriter)
	event, err := command.Handle()
	if err != nil {
		return err
	}

	projector := projector.InventoryExpenseAdditionProjector{CurrentState: currentState}
	state := projector.Project(event)

	storeErr := a.stateStoreWriter.StoreState(request.Email, state)
	if storeErr != nil {
		return storeErr
	}

	return nil
}
