package command

import (
	"product-service/cmd/product-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type SubtractProductQuantityCommand struct {
	ID               string
	Reason           string
	ByQuantity       int32
	EventStoreWriter eventstorelib.EventStoreWriter
	StateStoreReader statestorelib.StateStoreReader
}

func (s *SubtractProductQuantityCommand) Handle() (*eventlib.BaseEvent, error) {
	_, err := s.StateStoreReader.ReadState(s.ID)
	if err != nil {
		return nil, ErrProductNotFound
	}

	evt := event.NewProductHadItsQuantitySubtractedEvent(s.Reason, s.ByQuantity, s.ID)

	_, storeErr := s.EventStoreWriter.StoreEvent(evt)
	if storeErr != nil {
		return nil, err
	}

	return evt, nil

}
