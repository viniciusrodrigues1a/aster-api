package query

import (
	"encoding/json"
	"inventory-service/cmd/inventory-service/domain/projector"

	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type InventoryQuery struct {
	Email            string
	StateStoreReader statestorelib.StateStoreReader
}

func (q *InventoryQuery) ExecuteQuery() (*projector.InventoryState, error) {
	stateString, err := q.StateStoreReader.ReadState(q.Email)
	if err != nil {
		return nil, ErrInventoryNotFound
	}

	inventory := &projector.InventoryState{}
	json.Unmarshal([]byte(stateString), &inventory)

	return inventory, nil
}
