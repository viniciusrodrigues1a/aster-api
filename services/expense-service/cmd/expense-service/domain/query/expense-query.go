package query

import (
	"encoding/json"
	"errors"
	"expense-service/cmd/expense-service/domain/projector"

	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type ExpenseQuery struct {
	ID                      string
	StateStoreReader        statestorelib.StateStoreReader
	ProductStateStoreReader statestorelib.StateStoreReader
}

var ErrExpenseNotFound = errors.New("expense was not found")

type Product struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	PurchasePrice int64  `json:"purchase_price"`
}

type ExpenseQueryResponse struct {
	Product     *Product `json:"product"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Value       int64    `json:"value"`
	CreatedAt   int64    `json:"created_at"`
}

func (q *ExpenseQuery) ExecuteQuery() (*ExpenseQueryResponse, error) {
	stateString, err := q.StateStoreReader.ReadState(q.ID)
	if err != nil {
		return nil, ErrExpenseNotFound
	}

	expense := &projector.ExpenseState{}
	json.Unmarshal([]byte(stateString), expense)

	var product *Product
	if expense.ProductID != nil {
		productStateString, err := q.ProductStateStoreReader.ReadState(*expense.ProductID)
		if err == nil {
			p := &Product{}
			json.Unmarshal([]byte(productStateString), p)
			product = p
		}
	}

	response := &ExpenseQueryResponse{
		Product:     product,
		Title:       expense.Title,
		Description: expense.Description,
		Value:       expense.Value,
		CreatedAt:   expense.CreatedAt,
	}

	return response, nil
}
