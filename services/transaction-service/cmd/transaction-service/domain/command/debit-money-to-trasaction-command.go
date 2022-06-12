package command

import (
	"encoding/json"
	"transaction-service/cmd/transaction-service/domain/dto"
	"transaction-service/cmd/transaction-service/domain/event"
	"transaction-service/cmd/transaction-service/domain/projector"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type DebitMoneyToTransactionCommand struct {
	ID                      string
	Amount                  int64
	EventStoreWriter        eventstorelib.EventStoreWriter
	StateStoreReader        statestorelib.StateStoreReader
	ProductStateStoreReader statestorelib.StateStoreReader
}

// Handle stores a MoneyWasDebitedToTransactionEvent to the event store and returns the resulting event.
// Returns ErrTransactionDoesntExist if it can't read the state from the state store or if it has been soft deleted.
// Returns ErrProductCouldntBeFound if it can't read the state from the product state store.
// Returns ErrValuePaidCantBeGreaterThanTotalValue if transaction.ValuePaid > product.SalePrice * transaction.Quantity
func (d *DebitMoneyToTransactionCommand) Handle() (*eventlib.BaseEvent, error) {
	stateString, err := d.StateStoreReader.ReadState(d.ID)
	if err != nil {
		return nil, ErrTransactionDoesntExist
	}

	transaction := projector.TransactionState{}
	json.Unmarshal([]byte(stateString), &transaction)

	if transaction.DeletedAt > 0 {
		return nil, ErrTransactionDoesntExist
	}

	productStateString, err := d.ProductStateStoreReader.ReadState(*transaction.ProductID)
	if err != nil {
		return nil, ErrProductCouldntBeFound
	}

	product := dto.ProductState{}
	json.Unmarshal([]byte(productStateString), &product)

	totalValue := product.SalePrice * transaction.Quantity
	valuePaid := transaction.ValuePaid + d.Amount
	if valuePaid > totalValue {
		return nil, ErrValuePaidCantBeGreaterThanTotalValue
	}

	status := "open"
	if valuePaid == totalValue {
		status = "closed"
	}

	evt := event.NewMoneyWasDebitedToTransactionEvent(d.Amount, status, d.ID)

	_, storeErr := d.EventStoreWriter.StoreEvent(evt)
	if storeErr != nil {
		return nil, storeErr
	}

	return evt, nil
}
