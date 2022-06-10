package factory

import (
	usecase "transaction-service/cmd/transaction-service/application/use-case"
	"transaction-service/cmd/transaction-service/external/messaging"
)

var eventStoreRepository = makeMongoEventStoreRepository()
var stateStoreRepository = makeRedisStateStoreRepository()
var stateEmitter = messaging.NewTransactionEventStateEmitter(messaging.New())

func makeCreateTransactionUseCase() *usecase.CreateTransactionUseCase {
	return usecase.NewCreateTransactionUseCase(stateEmitter, eventStoreRepository, stateStoreRepository)
}

func makeUpdateTransactionUseCase() *usecase.UpdateTransactionUseCase {
	return usecase.NewUpdateTransactionUseCase(stateEmitter, eventStoreRepository, stateStoreRepository, stateStoreRepository)
}

func makeDeleteTransactionUseCase() *usecase.DeleteTransactionUseCase {
	return usecase.NewDeleteTransactionUseCase(stateEmitter, eventStoreRepository, stateStoreRepository, stateStoreRepository)
}
