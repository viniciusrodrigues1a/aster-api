package controller

import (
	"encoding/json"
	"inventory-service/cmd/inventory-service/domain/query"
	"net/http"

	"github.com/gorilla/mux"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type ListInventoryController struct {
	stateStoreReader statestorelib.StateStoreReader
}

func NewListInventoryController(sttStoreR statestorelib.StateStoreReader) *ListInventoryController {
	return &ListInventoryController{
		stateStoreReader: sttStoreR,
	}
}

func (c *ListInventoryController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["id"]

	qry := &query.InventoryQuery{
		Email:            email,
		StateStoreReader: c.stateStoreReader,
	}

	inventory, err := qry.ExecuteQuery()
	if err != nil {
		if err == query.ErrInventoryNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	inv, err := json.Marshal(inventory)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(inv)

}
