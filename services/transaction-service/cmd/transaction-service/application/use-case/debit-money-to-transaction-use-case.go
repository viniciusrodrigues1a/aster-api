package usecase

import (
	"encoding/json"
	"transaction-service/cmd/transaction-service/domain/command"
	"transaction-service/cmd/transaction-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type DebitMoneyToTransactionUseCase struct {
	stateEmitter            StateEmitter
	eventStoreWriter        eventstorelib.EventStoreWriter
	stateStoreReader        statestorelib.StateStoreReader
	productStateStoreReader statestorelib.StateStoreReader
	stateStoreWriter        statestorelib.StateStoreWriter
}

func NewDebitMoneyToTransactionUseCase(
	sttEmitter StateEmitter,
	evtStore eventstorelib.EventStoreWriter,
	sttStoreR statestorelib.StateStoreReader,
	sttStoreW statestorelib.StateStoreWriter,
	productSttStoreR statestorelib.StateStoreReader,
) *DebitMoneyToTransactionUseCase {
	return &DebitMoneyToTransactionUseCase{
		stateEmitter:            sttEmitter,
		eventStoreWriter:        evtStore,
		stateStoreReader:        sttStoreR,
		stateStoreWriter:        sttStoreW,
		productStateStoreReader: productSttStoreR,
	}
}

type DebitMoneyToTransactionUseCaseRequest struct {
	AccountID string
	ID        string
	Amount    int64
}

// Execute issues the DebitMoneyToTransactionCommand, projects the new state, stores it in the state store
// and emits a message with the new projected state
func (d *DebitMoneyToTransactionUseCase) Execute(request *DebitMoneyToTransactionUseCaseRequest) error {
	cmd := command.DebitMoneyToTransactionCommand{
		ID:                      request.ID,
		Amount:                  request.Amount,
		EventStoreWriter:        d.eventStoreWriter,
		StateStoreReader:        d.stateStoreReader,
		ProductStateStoreReader: d.productStateStoreReader,
	}
	evt, err := cmd.Handle()
	if err != nil {
		return err
	}

	val, _ := d.stateStoreReader.ReadState(request.ID)
	currentState := projector.TransactionState{}
	json.Unmarshal([]byte(val), &currentState)
	prj := projector.DebitMoneyToTransactionProjector{CurrentState: &currentState}
	state := prj.Project(evt)

	stateErr := d.stateStoreWriter.StoreState(evt.Data.StreamId.Hex(), state)
	if stateErr != nil {
		return stateErr
	}

	d.stateEmitter.Emit(*state, evt.Data.StreamId.Hex(), request.AccountID)

	return nil
}
