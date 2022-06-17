package factory

import (
	usecase "product-service/cmd/product-service/application/use-case"
	"product-service/cmd/product-service/external/messaging"
)

var eventStoreRepository = makeMongoEventStoreRepository()
var stateStoreRepository = makeRedisStateStoreRepository()
var stateEmitter = messaging.NewProductEventStateEmitter(messaging.New())

func makeCreateProductUseCase() *usecase.CreateProductUseCase {
	return usecase.NewCreateProductUseCase(stateEmitter, eventStoreRepository, stateStoreRepository)
}

func makeUpdateProductUseCase() *usecase.UpdateProductUseCase {
	return usecase.NewUpdateProductUseCase(stateEmitter, eventStoreRepository, stateStoreRepository)
}

func MakeSubtractProductQuantityUseCase() *usecase.SubtractProductQuantityUseCase {
	return usecase.NewSubtractProductQuantityUseCase(eventStoreRepository, stateStoreRepository, stateStoreRepository)
}
