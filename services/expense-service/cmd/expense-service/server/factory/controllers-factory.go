package factory

import (
	"expense-service/cmd/expense-service/application/controller"
)

func MakeCreateExpenseController() *controller.CreateExpenseController {
	useCase := makeCreateExpenseUseCase()

	return controller.NewCreateExpenseController(useCase)
}

func MakeUpdateExpenseController() *controller.UpdateExpenseController {
	useCase := makeUpdateExpenseUseCase()

	return controller.NewUpdateExpenseController(useCase)
}

func MakeDeleteExpenseController() *controller.DeleteExpenseController {
	useCase := makeDeleteExpenseUseCase()

	return controller.NewDeleteExpenseController(useCase)
}
