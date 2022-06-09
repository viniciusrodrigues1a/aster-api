package usecase

import (
	"encoding/json"
	"expense-service/cmd/expense-service/domain/command"
	"expense-service/cmd/expense-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type UpdateExpenseUseCase struct {
	stateEmitter     StateEmitter
	eventStoreWriter eventstorelib.EventStoreWriter
	stateStoreReader statestorelib.StateStoreReader
	stateStoreWriter statestorelib.StateStoreWriter
}

func NewUpdateExpenseUseCase(sttEmitter StateEmitter, evtStore eventstorelib.EventStoreWriter, sttStoreR statestorelib.StateStoreReader, sttStoreW statestorelib.StateStoreWriter) *UpdateExpenseUseCase {
	return &UpdateExpenseUseCase{
		stateEmitter:     sttEmitter,
		eventStoreWriter: evtStore,
		stateStoreReader: sttStoreR,
		stateStoreWriter: sttStoreW,
	}
}

type UpdateExpenseUseCaseRequest struct {
	ID          string
	AccountID   string
	Title       string
	Description string
	Value       int64
}

// Issues the UpdateExpenseCommand, projects the new state, stores it in the state store
// and emits a message with the new projected state
func (u *UpdateExpenseUseCase) Execute(request *UpdateExpenseUseCaseRequest) error {
	command := command.UpdateExpenseCommand{
		Id:               request.ID,
		Title:            request.Title,
		Description:      request.Description,
		Value:            request.Value,
		EventStoreWriter: u.eventStoreWriter,
		StateStoreReader: u.stateStoreReader,
	}
	event, err := command.Handle()
	if err != nil {
		return err
	}

	val, _ := u.stateStoreReader.ReadState(request.ID)
	currentState := projector.ExpenseState{}
	json.Unmarshal([]byte(val), &currentState)
	projector := projector.ExpenseUpdateProjector{CurrentState: &currentState}
	state := projector.Project(event)

	stateErr := u.stateStoreWriter.StoreState(event.Data.StreamId.Hex(), state)
	if stateErr != nil {
		return stateErr
	}

	u.stateEmitter.Emit(*state, event.Data.StreamId.Hex(), request.AccountID)

	return nil
}
