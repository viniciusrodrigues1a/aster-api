package factory

import "account-service/cmd/account-service/application/controller"

func MakeCreateAccountController() *controller.CreateAccountController {
	useCase := makeCreateAccountUseCase()

	return controller.NewCreateAccountController(useCase)
}
