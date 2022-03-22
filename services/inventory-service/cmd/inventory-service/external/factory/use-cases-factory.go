package factory

import usecase "inventory-service/cmd/inventory-service/application/use-case"

func MakeCreateInventoryUseCase() *usecase.CreateInventoryUseCase {
	return usecase.NewCreateInventoryUseCase(
		makeMongoEventStoreRepository(),
		makeRedisStateStoreRepository(),
	)
}
