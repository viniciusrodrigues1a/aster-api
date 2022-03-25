package usecase

import (
	"encoding/json"
	"expense-service/cmd/expense-service/domain/command"
	"expense-service/cmd/expense-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type DeleteExpenseUseCase struct {
	eventStoreWriter eventstorelib.EventStoreWriter
	stateStoreReader statestorelib.StateStoreReader
	stateStoreWriter statestorelib.StateStoreWriter
}

func NewDeleteExpenseUseCase(evtStore eventstorelib.EventStoreWriter, sttStoreR statestorelib.StateStoreReader, sttStoreW statestorelib.StateStoreWriter) *DeleteExpenseUseCase {
	return &DeleteExpenseUseCase{
		eventStoreWriter: evtStore,
		stateStoreReader: sttStoreR,
		stateStoreWriter: sttStoreW,
	}
}

type DeleteExpenseUseCaseRequest struct {
	Id string
}

func (d *DeleteExpenseUseCase) Execute(request *DeleteExpenseUseCaseRequest) error {
	command := command.DeleteExpenseCommand{Id: request.Id}
	event := command.Handle()

	val, err := d.stateStoreReader.ReadState(request.Id)
	if err != nil {
		return ErrExpenseDoesntExist
	}

	id, err := d.eventStoreWriter.StoreEvent(event)
	if err != nil {
		return err
	}

	currentState := projector.ExpenseState{}
	json.Unmarshal([]byte(val.(string)), &currentState)
	projector := projector.ExpenseDeletionProjector{CurrentState: &currentState}
	state := projector.Project(event)

	stateErr := d.stateStoreWriter.StoreState(id, state)
	if stateErr != nil {
		return stateErr
	}

	return nil
}
