package controller

import (
	"encoding/json"
	"net/http"
	usecase "transaction-service/cmd/transaction-service/application/use-case"
)

type CreateTransactionController struct {
	useCase *usecase.CreateTransactionUseCase
}

func NewCreateTransactionController(u *usecase.CreateTransactionUseCase) *CreateTransactionController {
	return &CreateTransactionController{
		useCase: u,
	}
}

func (c *CreateTransactionController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var body usecase.CreateTransactionUseCaseRequest

	decoderErr := json.NewDecoder(r.Body).Decode(&body)
	if decoderErr != nil {
		http.Error(w, decoderErr.Error(), http.StatusBadRequest)
		return
	}

	body.AccountID = r.Context().Value("account_id").(string)

	err := c.useCase.Execute(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
