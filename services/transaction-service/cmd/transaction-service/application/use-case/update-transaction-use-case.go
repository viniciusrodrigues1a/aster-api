package usecase

import (
	"encoding/json"
	"transaction-service/cmd/transaction-service/domain/command"
	"transaction-service/cmd/transaction-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type CommandEmitter interface {
	Emit(message interface{})
}

type UpdateTransactionUseCase struct {
	stateEmitter            StateEmitter
	eventStoreWriter        eventstorelib.EventStoreWriter
	stateStoreReader        statestorelib.StateStoreReader
	productStateStoreReader statestorelib.StateStoreReader
	stateStoreWriter        statestorelib.StateStoreWriter
	commandEmitter          CommandEmitter
}

func NewUpdateTransactionUseCase(
	sttEmitter StateEmitter,
	evtStore eventstorelib.EventStoreWriter,
	sttStoreR statestorelib.StateStoreReader,
	sttStoreW statestorelib.StateStoreWriter,
	productSttStoreR statestorelib.StateStoreReader,
	cmdEmitter CommandEmitter,
) *UpdateTransactionUseCase {
	return &UpdateTransactionUseCase{
		stateEmitter:            sttEmitter,
		eventStoreWriter:        evtStore,
		stateStoreReader:        sttStoreR,
		stateStoreWriter:        sttStoreW,
		productStateStoreReader: productSttStoreR,
		commandEmitter:          cmdEmitter,
	}
}

type UpdateTransactionUseCaseRequest struct {
	ProductID   *string `json:"product_id"`
	ID          string
	AccountID   string
	Quantity    int64  `json:"quantity"`
	ValuePaid   int64  `json:"value_paid"`
	Description string `json:"description"`
}

type emittedCommandMessage struct {
	ProductID string `json:"product_id"`
	Quantity  int64  `json:quantity`
}

// Issues the UpdateTransactionCommand, projects the new state, stores it in the state store
// and emits a message with the new projected state
func (u *UpdateTransactionUseCase) Execute(request *UpdateTransactionUseCaseRequest) error {
	cmd := command.UpdateTransactionCommand{
		ProductID:               request.ProductID,
		ID:                      request.ID,
		Quantity:                request.Quantity,
		ValuePaid:               request.ValuePaid,
		Description:             request.Description,
		EventStoreWriter:        u.eventStoreWriter,
		StateStoreReader:        u.stateStoreReader,
		ProductStateStoreReader: u.productStateStoreReader,
	}
	event, err := cmd.Handle()
	if err != nil {
		return err
	}

	val, _ := u.stateStoreReader.ReadState(request.ID)
	currentState := projector.TransactionState{}
	json.Unmarshal([]byte(val), &currentState)
	prj := projector.TransactionUpdateProjector{CurrentState: &currentState}
	state := prj.Project(event)

	stateErr := u.stateStoreWriter.StoreState(event.Data.StreamId.Hex(), state)
	if stateErr != nil {
		return stateErr
	}

	if request.Quantity != 0 && request.Quantity > currentState.Quantity {
		msg := emittedCommandMessage{
			ProductID: *request.ProductID,
			Quantity:  currentState.Quantity + request.Quantity,
		}
		u.commandEmitter.Emit(msg)
	}

	if request.Quantity != 0 && request.Quantity < currentState.Quantity {
		msg := emittedCommandMessage{
			ProductID: *request.ProductID,
			Quantity:  currentState.Quantity - request.Quantity,
		}
		u.commandEmitter.Emit(msg)
	}
	u.stateEmitter.Emit(*state, event.Data.StreamId.Hex(), request.AccountID)

	return nil

}
