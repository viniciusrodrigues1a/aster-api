package factory

import (
	"transaction-service/cmd/transaction-service/application/controller"
	"transaction-service/cmd/transaction-service/external/messaging"
)

func MakeCreateTransactionController() *controller.CreateTransactionController {
	useCase := makeCreateTransactionUseCase()

	return controller.NewCreateTransactionController(useCase, messaging.NewSubtractProductQuantityEmitter(MessagingConn))
}

func MakeUpdateTransactionController() *controller.UpdateTransactionController {
	useCase := makeUpdateTransactionUseCase()

	return controller.NewUpdateTransactionController(useCase)
}

func MakeDeleteTransactionController() *controller.DeleteTransactionController {
	useCase := makeDeleteTransactionUseCase()

	return controller.NewDeleteTransactionController(useCase)
}

func MakeDebitMoneyToTransactionController() *controller.DebitMoneyToTransactionController {
	useCase := makeDebitMoneyToTransactionUseCase()

	return controller.NewDebitMoneyToTransactionController(useCase)
}

func MakeListTransactionController() *controller.ListTransactionController {
	return controller.NewListTransactionController(stateStoreRepository, productStateStoreRepository)
}
