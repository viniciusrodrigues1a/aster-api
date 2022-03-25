package factory

import (
	usecase "expense-service/cmd/expense-service/application/use-case"
	"expense-service/cmd/expense-service/external/messaging"
)

var eventStoreRepository = makeMongoEventStoreRepository()
var stateStoreRepository = makeRedisStateStoreRepository()
var stateEmitter = messaging.NewExpenseEventStateEmitter(messaging.New())

func makeCreateExpenseUseCase() *usecase.CreateExpenseUseCase {
	return usecase.NewCreateExpenseUseCase(stateEmitter, eventStoreRepository, stateStoreRepository)
}

func makeUpdateExpenseUseCase() *usecase.UpdateExpenseUseCase {
	return usecase.NewUpdateExpenseUseCase(stateEmitter, eventStoreRepository, stateStoreRepository, stateStoreRepository)
}

func makeDeleteExpenseUseCase() *usecase.DeleteExpenseUseCase {
	return usecase.NewDeleteExpenseUseCase(stateEmitter, eventStoreRepository, stateStoreRepository, stateStoreRepository)
}
