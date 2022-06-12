package controller

import (
	"encoding/json"
	"net/http"
	usecase "transaction-service/cmd/transaction-service/application/use-case"
	"transaction-service/cmd/transaction-service/domain/command"

	"github.com/gorilla/mux"
)

type DebitMoneyToTransactionController struct {
	useCase *usecase.DebitMoneyToTransactionUseCase
}

func NewDebitMoneyToTransactionController(u *usecase.DebitMoneyToTransactionUseCase) *DebitMoneyToTransactionController {
	return &DebitMoneyToTransactionController{
		useCase: u,
	}
}

func (d *DebitMoneyToTransactionController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var body usecase.DebitMoneyToTransactionUseCaseRequest

	decoderErr := json.NewDecoder(r.Body).Decode(&body)
	if decoderErr != nil {
		http.Error(w, decoderErr.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	body.ID = id
	body.AccountID = r.Context().Value("account_id").(string)

	err := d.useCase.Execute(&body)
	if err != nil {
		status := http.StatusInternalServerError

		if err == command.ErrTransactionDoesntExist || err == command.ErrProductCouldntBeFound {
			status = http.StatusNotFound
		}
		if err == command.ErrValuePaidCantBeGreaterThanTotalValue {
			status = http.StatusBadRequest
		}

		http.Error(w, err.Error(), status)
	}
}
