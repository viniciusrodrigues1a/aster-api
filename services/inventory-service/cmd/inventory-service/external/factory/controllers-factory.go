package factory

import "inventory-service/cmd/inventory-service/application/controller"

func MakeListInventoryController() *controller.ListInventoryController {
	return controller.NewListInventoryController(stateStoreRepository)
}
