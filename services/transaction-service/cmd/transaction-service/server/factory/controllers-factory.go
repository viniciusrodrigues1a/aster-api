package factory

import "transaction-service/cmd/transaction-service/application/controller"

func MakeCreateTransactionController() *controller.CreateTransactionController {
	useCase := makeCreateTransactionUseCase()

	return controller.NewCreateTransactionController(useCase)
}

func MakeUpdateTransactionController() *controller.UpdateTransactionController {
	useCase := makeUpdateTransactionUseCase()

	return controller.NewUpdateTransactionController(useCase)
}

func MakeDeleteTransactionController() *controller.DeleteTransactionController {
	useCase := makeDeleteTransactionUseCase()

	return controller.NewDeleteTransactionController(useCase)
}
