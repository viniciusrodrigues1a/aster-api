package usecase

import (
	"encoding/json"
	"inventory-service/cmd/inventory-service/domain/command"
	"inventory-service/cmd/inventory-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type AddTransactionToInventoryUseCase struct {
	eventStoreWriter eventstorelib.EventStoreWriter
	stateStoreReader statestorelib.StateStoreReader
	stateStoreWriter statestorelib.StateStoreWriter
}

func NewAddTransactionToInventoryUseCase(
	evtStore eventstorelib.EventStoreWriter,
	sttStoreR statestorelib.StateStoreReader,
	sttStoreW statestorelib.StateStoreWriter,
) *AddTransactionToInventoryUseCase {
	return &AddTransactionToInventoryUseCase{
		eventStoreWriter: evtStore,
		stateStoreReader: sttStoreR,
		stateStoreWriter: sttStoreW,
	}
}

type AddTransactionToInventoryUseCaseRequest struct {
	ID          string `json:"id"`
	Email       string `json:"account_id"`
	Description string `json:"description"`
	ValuePaid   int64  `json:"value_paid"`
	DeletedAt   int64  `json:"deleted_at"`
}

func (a *AddTransactionToInventoryUseCase) Execute(request *AddTransactionToInventoryUseCaseRequest) error {
	val, err := a.stateStoreReader.ReadState(request.Email)
	if err != nil {
		return ErrInventoryDoesntExist
	}
	currentState := projector.InventoryState{}
	json.Unmarshal([]byte(val), &currentState)

	cmd := command.NewAddTransactionToInventoryCommand(request.ID, currentState.ID, request.Description, request.ValuePaid, request.DeletedAt, a.eventStoreWriter)
	event, err := cmd.Handle()
	if err != nil {
		return err
	}

	prj := projector.InventoryTransactionAdditionProjector{CurrentState: &currentState}
	state := prj.Project(event)

	storeErr := a.stateStoreWriter.StoreState(request.Email, state)
	if storeErr != nil {
		return storeErr
	}

	return nil
}
