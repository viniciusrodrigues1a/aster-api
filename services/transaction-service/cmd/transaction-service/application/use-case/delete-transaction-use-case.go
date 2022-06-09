package usecase

import (
	"encoding/json"
	"transaction-service/cmd/transaction-service/domain/command"
	"transaction-service/cmd/transaction-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type DeleteTransactionUseCase struct {
	stateEmitter     StateEmitter
	eventStoreWriter eventstorelib.EventStoreWriter
	stateStoreReader statestorelib.StateStoreReader
	stateStoreWriter statestorelib.StateStoreWriter
}

func NewDeleteTransactionUseCase(
	sttEmitter StateEmitter,
	evtStore eventstorelib.EventStoreWriter,
	sttStoreR statestorelib.StateStoreReader,
	sttStoreW statestorelib.StateStoreWriter,
) *DeleteTransactionUseCase {
	return &DeleteTransactionUseCase{
		stateEmitter:     sttEmitter,
		eventStoreWriter: evtStore,
		stateStoreReader: sttStoreR,
		stateStoreWriter: sttStoreW,
	}
}

type DeleteTransactionUseCaseRequest struct {
	ID        string
	AccountID string
}

// Issues the DeleteTransactionCommand, projects the new state, stores it in the state store
// and emits a message with the new projected state
func (d *DeleteTransactionUseCase) Execute(request *DeleteTransactionUseCaseRequest) error {
	cmd := command.DeleteTransactionCommand{
		ID:               request.ID,
		EventStoreWriter: d.eventStoreWriter,
		StateStoreReader: d.stateStoreReader,
	}
	event, err := cmd.Handle()
	if err != nil {
		return err
	}

	val, _ := d.stateStoreReader.ReadState(request.ID)
	currentState := projector.TransactionState{}
	json.Unmarshal([]byte(val), &currentState)
	prj := projector.TransactionDeletionProjector{CurrentState: &currentState}
	state := prj.Project(event)

	stateErr := d.stateStoreWriter.StoreState(event.Data.StreamId.Hex(), state)
	if stateErr != nil {
		return stateErr
	}

	d.stateEmitter.Emit(*state, event.Data.StreamId.Hex(), request.AccountID)

	return nil
}
