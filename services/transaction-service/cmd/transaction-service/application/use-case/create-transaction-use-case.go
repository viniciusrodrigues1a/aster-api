package usecase

import (
	"transaction-service/cmd/transaction-service/domain/command"
	"transaction-service/cmd/transaction-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type CreateTransactionUseCase struct {
	stateEmitter           StateEmitter
	eventStoreStreamWriter eventstorelib.EventStoreStreamWriter
	stateStoreWriter       statestorelib.StateStoreWriter
}

func NewCreateTransactionUseCase(sttEmitter StateEmitter, evtStore eventstorelib.EventStoreStreamWriter, sttStore statestorelib.StateStoreWriter) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		stateEmitter:           sttEmitter,
		eventStoreStreamWriter: evtStore,
		stateStoreWriter:       sttStore,
	}
}

type CreateTransactionUseCaseRequest struct {
	AccountID   string
	ValuePaid   int64  `json:"value_paid"`
	Description string `json:"description"`
}

// Issues the CreateTransactionCommand, projects the new state, stores it in the state store
// and emits a message with the new projected state
func (c *CreateTransactionUseCase) Execute(request *CreateTransactionUseCaseRequest) error {
	cmd := command.CreateTransactionCommand{
		ValuePaid:              request.ValuePaid,
		Description:            request.Description,
		EventStoreStreamWriter: c.eventStoreStreamWriter,
	}
	event, err := cmd.Handle()
	if err != nil {
		return err
	}

	projector := projector.TransactionCreationProjector{}
	state := projector.Project(event)

	stateErr := c.stateStoreWriter.StoreState(event.Data.StreamId.Hex(), state)
	if stateErr != nil {
		return stateErr
	}

	c.stateEmitter.Emit(*state, event.Data.StreamId.Hex(), request.AccountID)

	return nil
}
