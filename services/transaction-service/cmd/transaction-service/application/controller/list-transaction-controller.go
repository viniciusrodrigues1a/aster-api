package controller

import (
	"encoding/json"
	"net/http"
	"transaction-service/cmd/transaction-service/domain/query"

	"github.com/gorilla/mux"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type ListTransactionController struct {
	stateStoreReader   statestorelib.StateStoreReader
	productStoreReader statestorelib.StateStoreReader
}

func NewListTransactionController(sttStoreR, productSttStoreR statestorelib.StateStoreReader) *ListTransactionController {
	return &ListTransactionController{
		stateStoreReader:   sttStoreR,
		productStoreReader: productSttStoreR,
	}
}

func (c *ListTransactionController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	qry := &query.TransactionQuery{
		ID:                      id,
		StateStoreReader:        c.stateStoreReader,
		ProductStateStoreReader: c.productStoreReader,
	}

	expense, err := qry.ExecuteQuery()
	if err != nil {
		if err == query.ErrTransactionNotFound {
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
