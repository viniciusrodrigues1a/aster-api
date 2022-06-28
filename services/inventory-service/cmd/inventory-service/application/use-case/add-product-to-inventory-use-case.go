package usecase

import (
	"encoding/json"
	"inventory-service/cmd/inventory-service/domain/command"
	"inventory-service/cmd/inventory-service/domain/dto"
	"inventory-service/cmd/inventory-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type AddProductToInventoryUseCase struct {
	eventStoreWriter eventstorelib.EventStoreWriter
	stateStoreReader statestorelib.StateStoreReader
	stateStoreWriter statestorelib.StateStoreWriter
}

func NewAddProductToInventoryUseCase(
	evtStore eventstorelib.EventStoreWriter,
	sttStoreR statestorelib.StateStoreReader,
	sttStoreW statestorelib.StateStoreWriter,
) *AddProductToInventoryUseCase {
	return &AddProductToInventoryUseCase{
		eventStoreWriter: evtStore,
		stateStoreReader: sttStoreR,
		stateStoreWriter: sttStoreW,
	}
}

type AddProductToInventoryUseCaseRequest struct {
	ID            string            `json:"id"`
	Email         string            `json:"account_id"`
	Title         string            `json:"title"`
	Description   string            `json:"description"`
	Quantity      int64             `json:"quantity"`
	PurchasePrice int64             `json:"purchase_price"`
	SalePrice     int64             `json:"sale_price"`
	DeletedAt     int64             `json:"deleted_at"`
	Image         *dto.ProductImage `json:"image"`
}

func (a *AddProductToInventoryUseCase) Execute(request *AddProductToInventoryUseCaseRequest) error {
	val, err := a.stateStoreReader.ReadState(request.Email)
	if err != nil {
		return ErrInventoryDoesntExist
	}
	currentState := projector.InventoryState{}
	json.Unmarshal([]byte(val), &currentState)

	cmd := command.NewAddProductToInventoryCommand(request.ID, currentState.ID, request.Title, request.Description, request.Quantity, request.PurchasePrice, request.SalePrice, request.DeletedAt, request.Image, a.eventStoreWriter)
	evt, err := cmd.Handle()
	if err != nil {
		return err
	}

	prj := projector.InventoryProductAdditionProjector{CurrentState: &currentState}
	state := prj.Project(evt)

	storeErr := a.stateStoreWriter.StoreState(request.Email, state)
	if storeErr != nil {
		return storeErr
	}

	return nil
}
