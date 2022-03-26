package usecase

import (
	"encoding/json"
	"expense-service/cmd/expense-service/domain/command"
	"expense-service/cmd/expense-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type DeleteExpenseUseCase struct {
	stateEmitter     StateEmitter
	eventStoreWriter eventstorelib.EventStoreWriter
	stateStoreReader statestorelib.StateStoreReader
	stateStoreWriter statestorelib.StateStoreWriter
}

func NewDeleteExpenseUseCase(sttEmitter StateEmitter, evtStore eventstorelib.EventStoreWriter, sttStoreR statestorelib.StateStoreReader, sttStoreW statestorelib.StateStoreWriter) *DeleteExpenseUseCase {
	return &DeleteExpenseUseCase{
		stateEmitter:     sttEmitter,
		eventStoreWriter: evtStore,
		stateStoreReader: sttStoreR,
		stateStoreWriter: sttStoreW,
	}
}

type DeleteExpenseUseCaseRequest struct {
	Id string
}

func (d *DeleteExpenseUseCase) Execute(request *DeleteExpenseUseCaseRequest) error {
	command := command.DeleteExpenseCommand{
		Id:               request.Id,
		EventStoreWriter: d.eventStoreWriter,
		StateStoreReader: d.stateStoreReader,
	}
	event, err := command.Handle()
	if err != nil {
		return err
	}

	val, _ := d.stateStoreReader.ReadState(request.Id)
	currentState := projector.ExpenseState{}
	json.Unmarshal([]byte(val.(string)), &currentState)
	projector := projector.ExpenseDeletionProjector{CurrentState: &currentState}
	state := projector.Project(event)

	stateErr := d.stateStoreWriter.StoreState(event.Data.StreamId.Hex(), state)
	if stateErr != nil {
		return stateErr
	}

	d.stateEmitter.Emit(*state, event.Data.StreamId.Hex())

	return nil
}
