package query

import (
	"encoding/json"
	"errors"
	"transaction-service/cmd/transaction-service/domain/projector"

	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type TransactionQuery struct {
	ID                      string
	StateStoreReader        statestorelib.StateStoreReader
	ProductStateStoreReader statestorelib.StateStoreReader
}

var ErrTransactionNotFound = errors.New("transaction was not found")

type Product struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	SalePrice int64  `json:"sale_price"`
}

type TransactionQueryResponse struct {
	Product     *Product `json:"product"`
	Status      string   `json:"status"`
	Description string   `json:"description"`
	Quantity    int64    `json:"quantity"`
	ValuePaid   int64    `json:"value_paid"`
	CreatedAt   int64    `json:"created_at"`
}

func (q *TransactionQuery) ExecuteQuery() (*TransactionQueryResponse, error) {
	stateString, err := q.StateStoreReader.ReadState(q.ID)
	if err != nil {
		return nil, ErrTransactionNotFound
	}

	transaction := &projector.TransactionState{}
	json.Unmarshal([]byte(stateString), transaction)

	var product *Product
	if transaction.ProductID != nil {
		productStateString, err := q.ProductStateStoreReader.ReadState(*transaction.ProductID)
		if err == nil {
			p := &Product{}
			json.Unmarshal([]byte(productStateString), p)
			product = p
		}
	}

	response := &TransactionQueryResponse{
		Product:     product,
		Status:      transaction.Status,
		Description: transaction.Description,
		Quantity:    transaction.Quantity,
		ValuePaid:   transaction.ValuePaid,
		CreatedAt:   transaction.CreatedAt,
	}

	return response, nil
}
