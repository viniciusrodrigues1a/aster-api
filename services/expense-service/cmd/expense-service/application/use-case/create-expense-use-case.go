package usecase

import (
	"expense-service/cmd/expense-service/domain/command"
	"expense-service/cmd/expense-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type CreateExpenseUseCase struct {
	stateEmitter           StateEmitter
	eventStoreStreamWriter eventstorelib.EventStoreStreamWriter
	stateStoreWriter       statestorelib.StateStoreWriter
}

func NewCreateExpenseUseCase(sttEmitter StateEmitter, evtStore eventstorelib.EventStoreStreamWriter, sttStore statestorelib.StateStoreWriter) *CreateExpenseUseCase {
	return &CreateExpenseUseCase{
		stateEmitter:           sttEmitter,
		eventStoreStreamWriter: evtStore,
		stateStoreWriter:       sttStore,
	}
}

type CreateExpenseUseCaseRequest struct {
	Title       string
	Description string
	Value       int64
}

func (c *CreateExpenseUseCase) Execute(request *CreateExpenseUseCaseRequest) error {
	command := command.CreateExpenseCommand{
		Title:       request.Title,
		Description: request.Description,
		Value:       request.Value,
	}
	event := command.Handle()

	id, err := c.eventStoreStreamWriter.StoreEventStream(event)
	if err != nil {
		return err
	}

	projector := projector.ExpenseCreationProjector{}
	state := projector.Project(event)

	stateErr := c.stateStoreWriter.StoreState(id, state)
	if stateErr != nil {
		return stateErr
	}

	c.stateEmitter.Emit(*state, id)

	return nil
}
