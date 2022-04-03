package usecase

import (
	"encoding/json"
	"transaction-service/cmd/transaction-service/domain/command"
	"transaction-service/cmd/transaction-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type UpdateTransactionUseCase struct {
	stateEmitter     StateEmitter
	eventStoreWriter eventstorelib.EventStoreWriter
	stateStoreReader statestorelib.StateStoreReader
	stateStoreWriter statestorelib.StateStoreWriter
}

func NewUpdateTransactionUseCase(
	sttEmitter StateEmitter,
	evtStore eventstorelib.EventStoreWriter,
	sttStoreR statestorelib.StateStoreReader,
	sttStoreW statestorelib.StateStoreWriter,
) *UpdateTransactionUseCase {
	return &UpdateTransactionUseCase{
		stateEmitter:     sttEmitter,
		eventStoreWriter: evtStore,
		stateStoreReader: sttStoreR,
		stateStoreWriter: sttStoreW,
	}
}

type UpdateTransactionUseCaseRequest struct {
	ID          string
	AccountID   string
	ValuePaid   int64  `json:"value_paid"`
	Description string `json:"description"`
}

// Issues the UpdateTransactionCommand, projects the new state, stores it in the state store
// and emits a message with the new projected state
func (u *UpdateTransactionUseCase) Execute(request *UpdateTransactionUseCaseRequest) error {
	cmd := command.UpdateTransactionCommand{
		ID:               request.ID,
		ValuePaid:        request.ValuePaid,
		Description:      request.Description,
		EventStoreWriter: u.eventStoreWriter,
		StateStoreReader: u.stateStoreReader,
	}
	event, err := cmd.Handle()
	if err != nil {
		return err
	}

	val, _ := u.stateStoreReader.ReadState(request.ID)
	currentState := projector.TransactionState{}
	json.Unmarshal([]byte(val.(string)), &currentState)
	prj := projector.TransactionUpdateProjector{CurrentState: &currentState}
	state := prj.Project(event)

	stateErr := u.stateStoreWriter.StoreState(event.Data.StreamId.Hex(), state)
	if stateErr != nil {
		return stateErr
	}

	u.stateEmitter.Emit(*state, event.Data.StreamId.Hex(), request.AccountID)

	return nil

}