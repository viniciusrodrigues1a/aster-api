package factory

import usecase "product-service/cmd/product-service/application/use-case"

var eventStoreRepository = makeMongoEventStoreRepository()
var stateStoreRepository = makeRedisStateStoreRepository()

func makeCreateProductUseCase() *usecase.CreateProductUseCase {
	return usecase.NewCreateProductUseCase(eventStoreRepository, stateStoreRepository)
}

func makeUpdateProductUseCase() *usecase.UpdateProductUseCase {
	return usecase.NewUpdateProductUseCase(eventStoreRepository, stateStoreRepository)
}
