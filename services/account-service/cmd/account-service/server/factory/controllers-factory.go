package factory

import "account-service/cmd/account-service/application/controller"

func MakeCreateAccountController() *controller.CreateAccountController {
	useCase := makeCreateAccountUseCase()

	return controller.NewCreateAccountController(useCase)
}

func MakeLoginController() *controller.LoginController {
	useCase := makeLoginUseCase()

	return controller.NewLoginController(useCase)
}

func MakeValidateTokenController() *controller.ValidateTokenController {
	useCase := makeValidateTokenUseCase()

	return controller.NewValidateTokenController(useCase)
}
