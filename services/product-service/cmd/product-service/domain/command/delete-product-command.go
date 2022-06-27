package command

import (
	"product-service/cmd/product-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type DeleteProductCommand struct {
	ID               string
	EventStoreWriter eventstorelib.EventStoreWriter
	StateStoreReader statestorelib.StateStoreReader
}

func (d *DeleteProductCommand) Handle() (*eventlib.BaseEvent, error) {
	evt := event.NewProductWasDeletedEvent(d.ID)

	_, err := d.StateStoreReader.ReadState(d.ID)
	if err != nil {
		return nil, ErrProductNotFound
	}

	_, storeErr := d.EventStoreWriter.StoreEvent(evt)
	if storeErr != nil {
		return nil, storeErr
	}

	return evt, nil
}
