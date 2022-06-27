package usecase

import (
	"encoding/json"
	"product-service/cmd/product-service/domain/command"
	"product-service/cmd/product-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type AddProductQuantityUseCase struct {
	stateEmitter         StateEmitter
	eventStoreRepository eventstorelib.EventStoreWriter
	stateStoreReader     statestorelib.StateStoreReader
	stateStoreWriter     statestorelib.StateStoreWriter
}

func NewAddProductQuantityUseCase(
	sttEmitter StateEmitter,
	evtStore eventstorelib.EventStoreWriter,
	sttStoreR statestorelib.StateStoreReader,
	sttStoreW statestorelib.StateStoreWriter,
) *AddProductQuantityUseCase {
	return &AddProductQuantityUseCase{
		stateEmitter:         sttEmitter,
		eventStoreRepository: evtStore,
		stateStoreReader:     sttStoreR,
		stateStoreWriter:     sttStoreW,
	}
}

type AddProductQuantityUseCaseRequest struct {
	AccountID  string `json:"account_id"`
	ID         string `json:"id"`
	ByQuantity int32  `json:"by_quantity"`
}

func (a *AddProductQuantityUseCase) Execute(request *AddProductQuantityUseCaseRequest) error {
	cmd := command.AddProductQuantityCommand{
		ID:               request.ID,
		ByQuantity:       request.ByQuantity,
		EventStoreWriter: a.eventStoreRepository,
		StateStoreReader: a.stateStoreReader,
	}
	evt, err := cmd.Handle()
	if err != nil {
		return err
	}

	val, _ := a.stateStoreReader.ReadState(request.ID)
	currentState := projector.ProductState{}
	json.Unmarshal([]byte(val), &currentState)
	prj := projector.ProductQuantityAdditionProjector{CurrentState: currentState}
	state := prj.Project(evt)

	stateErr := a.stateStoreWriter.StoreState(request.ID, state)
	if stateErr != nil {
		return stateErr
	}

	a.stateEmitter.Emit(*state, evt.Data.StreamId.Hex(), request.AccountID)

	return nil
}
