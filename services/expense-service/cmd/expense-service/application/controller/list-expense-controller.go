package controller

import (
	"encoding/json"
	"expense-service/cmd/expense-service/domain/query"
	"net/http"

	"github.com/gorilla/mux"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type ListExpenseController struct {
	stateStoreReader   statestorelib.StateStoreReader
	productStoreReader statestorelib.StateStoreReader
}

func NewListExpenseController(sttStoreR, productSttStoreR statestorelib.StateStoreReader) *ListExpenseController {
	return &ListExpenseController{
		stateStoreReader:   sttStoreR,
		productStoreReader: productSttStoreR,
	}
}

func (c *ListExpenseController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	qry := &query.ExpenseQuery{
		ID:                      id,
		StateStoreReader:        c.stateStoreReader,
		ProductStateStoreReader: c.productStoreReader,
	}

	expense, err := qry.ExecuteQuery()
	if err != nil {
		if err == query.ErrExpenseNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	exp, err := json.Marshal(expense)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(exp)
}
