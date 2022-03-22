package controller

import (
	usecase "account-service/cmd/account-service/application/use-case"
	"encoding/json"
	"net/http"
)

type CreateAccountController struct {
	useCase *usecase.CreateAccountUseCase
}

func NewCreateAccountController(usecase *usecase.CreateAccountUseCase) *CreateAccountController {
	return &CreateAccountController{
		useCase: usecase,
	}
}

func (c *CreateAccountController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var body usecase.CreateAccountUseCaseRequest

	decodeErr := json.NewDecoder(r.Body).Decode(&body)
	if decodeErr != nil {
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		return
	}

	err := c.useCase.Execute(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
