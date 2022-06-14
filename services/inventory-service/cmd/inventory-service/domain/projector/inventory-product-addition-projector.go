package projector

import (
	"inventory-service/cmd/inventory-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type InventoryProductAdditionProjector struct {
	CurrentState *InventoryState
}

func (i *InventoryProductAdditionProjector) Project(e *eventlib.BaseEvent) *InventoryState {
	payload := e.Payload.(event.ProductWasAddedToInventoryEvent)

	index := findIndexOfProductByID(i.CurrentState.Products, payload.ProductID)
	newProducts := i.CurrentState.Products
	newProduct := Product{
		ID:            payload.ProductID,
		Title:         payload.Title,
		Description:   payload.Description,
		Quantity:      payload.Quantity,
		PurchasePrice: payload.PurchasePrice,
		SalePrice:     payload.SalePrice,
	}

	if index == -1 { // add a new Product to the slice
		newProducts = append(newProducts, newProduct)
	} else { // update Product in the slice
		newProducts[index] = newProduct
	}

	if payload.DeletedAt > 0 && index > -1 { // remove Product from the slice
		newProducts = append(newProducts[:index], newProducts[index+1:]...)
	}

	return &InventoryState{
		ID:           i.CurrentState.ID,
		Participants: i.CurrentState.Participants,
		Expenses:     i.CurrentState.Expenses,
		Transactions: i.CurrentState.Transactions,
		Products:     newProducts,
	}
}

func findIndexOfProductByID(products []Product, id string) int {
	for index, v := range products {
		if v.ID == id {
			return index
		}
	}

	return -1
}
