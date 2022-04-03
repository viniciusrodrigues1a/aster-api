package controller

import (
	"encoding/json"
	"net/http"
	usecase "transaction-service/cmd/transaction-service/application/use-case"
	"transaction-service/cmd/transaction-service/domain/command"

	"github.com/gorilla/mux"
)

type UpdateTransactionController struct {
	useCase *usecase.UpdateTransactionUseCase
}

func NewUpdateTransactionController(u *usecase.UpdateTransactionUseCase) *UpdateTransactionController {
	return &UpdateTransactionController{
		useCase: u,
	}
}

func (u *UpdateTransactionController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var body usecase.UpdateTransactionUseCaseRequest

	decoderErr := json.NewDecoder(r.Body).Decode(&body)
	if decoderErr != nil {
		http.Error(w, decoderErr.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	body.ID = id
	body.AccountID = r.Context().Value("account_id").(string)

	err := u.useCase.Execute(&body)
	if err != nil {
		status := http.StatusInternalServerError

		if err == command.ErrTransactionDoesntExist {
			status = http.StatusNotFound
		}

		http.Error(w, err.Error(), status)
	}
}
