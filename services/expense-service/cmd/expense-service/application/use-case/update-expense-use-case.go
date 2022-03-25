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
	Id          string
	Title       string
	Description string
	Value       int64
}

func (u *UpdateExpenseUseCase) Execute(request *UpdateExpenseUseCaseRequest) error {
	command := command.UpdateExpenseCommand{
		Id:          request.Id,
		Title:       request.Title,
		Description: request.Description,
		Value:       request.Value,
	}
	event := command.Handle()

	val, err := u.stateStoreReader.ReadState(request.Id)
	if err != nil {
		return ErrExpenseDoesntExist
	}

	id, err := u.eventStoreWriter.StoreEvent(event)
	if err != nil {
		return err
	}

	currentState := projector.ExpenseState{}
	json.Unmarshal([]byte(val.(string)), &currentState)
	projector := projector.ExpenseUpdateProjector{CurrentState: &currentState}
	state := projector.Project(event)

	stateErr := u.stateStoreWriter.StoreState(id, state)
	if stateErr != nil {
		return stateErr
	}

	u.stateEmitter.Emit(*state, id)

	return nil
}
