package controller

import (
	"encoding/json"
	usecase "expense-service/cmd/expense-service/application/use-case"
	"expense-service/cmd/expense-service/domain/command"
	"net/http"
)

type CreateExpenseController struct {
	useCase *usecase.CreateExpenseUseCase
}

func NewCreateExpenseController(u *usecase.CreateExpenseUseCase) *CreateExpenseController {
	return &CreateExpenseController{
		useCase: u,
	}
}

func (c *CreateExpenseController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var body usecase.CreateExpenseUseCaseRequest

	decoderErr := json.NewDecoder(r.Body).Decode(&body)
	if decoderErr != nil {
		http.Error(w, decoderErr.Error(), http.StatusBadRequest)
		return
	}

	body.AccountId = r.Context().Value("account_id").(string)

	err := c.useCase.Execute(&body)
	if err != nil {
		status := http.StatusInternalServerError

		if err == command.ErrTitleIsRequired ||
			err == command.ErrProductCouldntBeFound ||
			err == command.ErrValueCantBeZero {
			status = http.StatusBadRequest
		}

		http.Error(w, err.Error(), status)
	}
}
