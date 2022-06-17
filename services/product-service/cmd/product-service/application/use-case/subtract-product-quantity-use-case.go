package usecase

import (
	"encoding/json"
	"product-service/cmd/product-service/domain/command"
	"product-service/cmd/product-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type SubtractProductQuantityUseCase struct {
	eventStoreRepository eventstorelib.EventStoreWriter
	stateStoreReader     statestorelib.StateStoreReader
	stateStoreWriter     statestorelib.StateStoreWriter
}

func NewSubtractProductQuantityUseCase(evtStore eventstorelib.EventStoreWriter, sttStoreR statestorelib.StateStoreReader, sttStoreW statestorelib.StateStoreWriter) *SubtractProductQuantityUseCase {
	return &SubtractProductQuantityUseCase{
		eventStoreRepository: evtStore,
		stateStoreReader:     sttStoreR,
		stateStoreWriter:     sttStoreW,
	}
}

type SubtractProductQuantityUseCaseRequest struct {
	ID         string `json:"id"`
	Reason     string `json:"reason"`
	ByQuantity int32  `json:"by_quantity"`
}

func (s *SubtractProductQuantityUseCase) Execute(request *SubtractProductQuantityUseCaseRequest) error {
	cmd := command.SubtractProductQuantityCommand{
		ID:               request.ID,
		Reason:           request.Reason,
		ByQuantity:       request.ByQuantity,
		EventStoreWriter: s.eventStoreRepository,
		StateStoreReader: s.stateStoreReader,
	}
	evt, err := cmd.Handle()
	if err != nil {
		return err
	}

	val, _ := s.stateStoreReader.ReadState(request.ID)
	currentState := projector.ProductState{}
	json.Unmarshal([]byte(val), &currentState)
	prj := projector.ProductQuantitySubtractionProjector{CurrentState: currentState}
	state := prj.Project(evt)

	stateErr := s.stateStoreWriter.StoreState(request.ID, state)
	if stateErr != nil {
		return stateErr
	}

	return nil
}
