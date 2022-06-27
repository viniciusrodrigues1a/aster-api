package command

import (
	"product-service/cmd/product-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type AddProductQuantityCommand struct {
	ID               string
	ByQuantity       int32
	EventStoreWriter eventstorelib.EventStoreWriter
	StateStoreReader statestorelib.StateStoreReader
}

func (i *AddProductQuantityCommand) Handle() (*eventlib.BaseEvent, error) {
	_, err := i.StateStoreReader.ReadState(i.ID)
	if err != nil {
		return nil, ErrProductNotFound
	}

	evt := event.NewProductHadItsQuantityAddedEvent(i.ByQuantity, i.ID)

	_, storeErr := i.EventStoreWriter.StoreEvent(evt)
	if storeErr != nil {
		return nil, err
	}

	return evt, nil
}
