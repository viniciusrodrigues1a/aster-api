package factory

import "expense-service/cmd/expense-service/application/controller"

func MakeCreateExpenseController() *controller.CreateExpenseController {
	useCase := makeCreateExpenseUseCase()

	return controller.NewCreateExpenseController(useCase)
}
