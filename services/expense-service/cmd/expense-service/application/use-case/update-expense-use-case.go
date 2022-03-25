package usecase

import (
	"errors"
	"expense-service/cmd/expense-service/domain/command"
	"expense-service/cmd/expense-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

var ErrExpenseDoesntExist = errors.New("expense doesn't exist")

type UpdateExpenseUseCase struct {
	eventStoreWriter eventstorelib.EventStoreWriter
	stateStoreReader statestorelib.StateStoreReader
	stateStoreWriter statestorelib.StateStoreWriter
}

func NewUpdateExpenseUseCase(evtStore eventstorelib.EventStoreWriter, sttStoreR statestorelib.StateStoreReader, sttStoreW statestorelib.StateStoreWriter) *UpdateExpenseUseCase {
	return &UpdateExpenseUseCase{
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

	if _, err := u.stateStoreReader.ReadState(request.Id); err != nil {
		return ErrExpenseDoesntExist
	}

	id, err := u.eventStoreWriter.StoreEvent(event)
	if err != nil {
		return err
	}

	projector := projector.ExpenseUpdateProjector{}
	state := projector.Project(event)

	stateErr := u.stateStoreWriter.StoreState(id, state)
	if stateErr != nil {
		return stateErr
	}

	return nil
}
