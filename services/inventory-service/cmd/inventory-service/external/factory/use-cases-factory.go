package factory

import usecase "inventory-service/cmd/inventory-service/application/use-case"

var eventStoreRepository = makeMongoEventStoreRepository()
var stateStoreRepository = makeRedisStateStoreRepository()

func MakeCreateInventoryUseCase() *usecase.CreateInventoryUseCase {
	return usecase.NewCreateInventoryUseCase(eventStoreRepository, stateStoreRepository)
}

func MakeAddExpenseToInventoryUseCase() *usecase.AddExpenseToInventoryUseCase {
	return usecase.NewAddExpenseToInventoryUseCase(eventStoreRepository, stateStoreRepository, stateStoreRepository)
}

func MakeAddTransactionToInventoryUseCase() *usecase.AddTransactionToInventoryUseCase {
	return usecase.NewAddTransactionToInventoryUseCase(eventStoreRepository, stateStoreRepository, stateStoreRepository)
}
