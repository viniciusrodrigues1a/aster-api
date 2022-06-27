package usecase

import (
	"encoding/json"
	"product-service/cmd/product-service/domain/command"
	"product-service/cmd/product-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type DeleteProductUseCase struct {
	stateEmitter         StateEmitter
	eventStoreRepository eventstorelib.EventStoreWriter
	stateStoreReader     statestorelib.StateStoreReader
	stateStoreWriter     statestorelib.StateStoreWriter
}

func NewDeleteProductUseCase(sttEmitter StateEmitter,
	evtStore eventstorelib.EventStoreWriter,
	sttStoreR statestorelib.StateStoreReader,
	sttStoreW statestorelib.StateStoreWriter,
) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		stateEmitter:         sttEmitter,
		eventStoreRepository: evtStore,
		stateStoreReader:     sttStoreR,
		stateStoreWriter:     sttStoreW,
	}
}

type DeleteProductUseCaseRequest struct {
	AccountID string `json:"account_id"`
	ID        string `json:"id"`
}

func (d *DeleteProductUseCase) Execute(request *DeleteProductUseCaseRequest) error {
	cmd := command.DeleteProductCommand{
		ID:               request.ID,
		EventStoreWriter: d.eventStoreRepository,
		StateStoreReader: d.stateStoreReader,
	}
	evt, err := cmd.Handle()
	if err != nil {
		return err
	}

	val, _ := d.stateStoreReader.ReadState(request.ID)
	currentState := projector.ProductState{}
	json.Unmarshal([]byte(val), &currentState)
	prj := projector.ProductDeletionProjector{CurrentState: &currentState}
	state := prj.Project(evt)

	stateErr := d.stateStoreWriter.StoreState(request.ID, state)
	if stateErr != nil {
		return stateErr
	}

	d.stateEmitter.Emit(*state, evt.Data.StreamId.Hex(), request.AccountID)

	return nil
}
