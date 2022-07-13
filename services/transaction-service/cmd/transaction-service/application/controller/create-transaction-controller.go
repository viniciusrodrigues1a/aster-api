package controller

import (
	"encoding/json"
	"net/http"
	usecase "transaction-service/cmd/transaction-service/application/use-case"
	"transaction-service/cmd/transaction-service/domain/command"
)

type CreateTransactionController struct {
	useCase        *usecase.CreateTransactionUseCase
	commandEmitter CommandEmitter
}

func NewCreateTransactionController(u *usecase.CreateTransactionUseCase, cmdEmitter CommandEmitter) *CreateTransactionController {
	return &CreateTransactionController{
		useCase:        u,
		commandEmitter: cmdEmitter,
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
	status := http.StatusInternalServerError

	if err != nil {
		if err == command.ErrProductIDIsRequired ||
			err == command.ErrQuantityMustBeGreaterThanZero ||
			err == command.ErrProductCouldntBeFound ||
			err == command.ErrThereIsNotEnoughOfThisProduct ||
			err == command.ErrValuePaidCantBeGreaterThanTotalValue {
			status = http.StatusBadRequest
		}

		http.Error(w, err.Error(), status)
	}

	c.commandEmitter.Emit(body)
}
