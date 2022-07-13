package usecase

import (
	"errors"
	"product-service/cmd/product-service/domain/command"
	"product-service/cmd/product-service/domain/dto"
	"product-service/cmd/product-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

var ErrTitleIsEmpty = errors.New("Title is empty")
var ErrPurchasePriceZero = errors.New("Purchase price can't be $ 0.00")
var ErrSalePriceZero = errors.New("Sale price can't be $ 0.00")

type CreateProductUseCase struct {
	stateEmitter         StateEmitter
	eventStoreRepository eventstorelib.EventStoreStreamWriter
	stateStoreRepository statestorelib.StateStoreWriter
}

func NewCreateProductUseCase(sttEmitter StateEmitter, evtStore eventstorelib.EventStoreStreamWriter, sttStore statestorelib.StateStoreWriter) *CreateProductUseCase {
	return &CreateProductUseCase{
		stateEmitter:         sttEmitter,
		eventStoreRepository: evtStore,
		stateStoreRepository: sttStore,
	}
}

type CreateProductUseCaseRequest struct {
	AccountID     string `json:"account_id"`
	Title         string
	Description   string
	Quantity      int32
	PurchasePrice int64             `json:"purchase_price"`
	SalePrice     int64             `json:"sale_price"`
	Image         *dto.ProductImage `json:"image"`
}

func (c *CreateProductUseCase) Execute(request *CreateProductUseCaseRequest) error {
	if request.Title == "" {
		return ErrTitleIsEmpty
	}

	if request.SalePrice == 0 {
		return ErrSalePriceZero
	}

	if request.PurchasePrice == 0 {
		return ErrPurchasePriceZero
	}

	command := command.CreateProductCommand{
		Title:         request.Title,
		Description:   request.Description,
		Quantity:      request.Quantity,
		PurchasePrice: request.PurchasePrice,
		SalePrice:     request.SalePrice,
		Image:         request.Image,
	}
	event := command.Handle()

	id, err := c.eventStoreRepository.StoreEventStream(event)
	if err != nil {
		return err
	}

	projector := projector.ProductCreationProjector{}
	state := projector.Project(event)

	stateErr := c.stateStoreRepository.StoreState(id, state)
	if stateErr != nil {
		return stateErr
	}

	c.stateEmitter.Emit(*state, event.Data.StreamId.Hex(), request.AccountID)

	return nil
}
