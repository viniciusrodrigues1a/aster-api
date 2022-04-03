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
	AccountId   string
}

// Issues the CreateExpenseCommand, projects the new state, stores it in the state store
// and emits a message with the new projected state
func (c *CreateExpenseUseCase) Execute(request *CreateExpenseUseCaseRequest) error {
	command := command.CreateExpenseCommand{
		Title:                  request.Title,
		Description:            request.Description,
		Value:                  request.Value,
		EventStoreStreamWriter: c.eventStoreStreamWriter,
	}
	event, err := command.Handle()
	if err != nil {
		return err
	}

	projector := projector.ExpenseCreationProjector{}
	state := projector.Project(event)

	stateErr := c.stateStoreWriter.StoreState(event.Data.StreamId.Hex(), state)
	if stateErr != nil {
		return stateErr
	}

	c.stateEmitter.Emit(*state, event.Data.StreamId.Hex(), request.AccountId)

	return nil
}
