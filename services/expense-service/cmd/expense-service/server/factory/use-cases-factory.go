package factory

import usecase "expense-service/cmd/expense-service/application/use-case"

var eventStoreRepository = makeMongoEventStoreRepository()
var stateStoreRepository = makeRedisStateStoreRepository()

func makeCreateExpenseUseCase() *usecase.CreateExpenseUseCase {
	return usecase.NewCreateExpenseUseCase(eventStoreRepository, stateStoreRepository)
}

func makeUpdateExpenseUseCase() *usecase.UpdateExpenseUseCase {
	return usecase.NewUpdateExpenseUseCase(eventStoreRepository, stateStoreRepository, stateStoreRepository)
}

func makeDeleteExpenseUseCase() *usecase.DeleteExpenseUseCase {
	return usecase.NewDeleteExpenseUseCase(eventStoreRepository, stateStoreRepository, stateStoreRepository)
}
