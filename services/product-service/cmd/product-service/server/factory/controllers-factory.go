package factory

import "product-service/cmd/product-service/application/controller"

func MakeCreateProductController() *controller.CreateProductController {
	useCase := makeCreateProductUseCase()

	return controller.NewCreateProductController(useCase)
}

func MakeUpdateProductController() *controller.UpdateProductController {
	useCase := makeUpdateProductUseCase()

	return controller.NewUpdateProductController(useCase)
}

func MakeDeleteProductController() *controller.DeleteProductController {
	useCase := makeDeleteProductUseCase()

	return controller.NewDeleteProductController(useCase)
}
