package controller

import (
	"net/http"
	usecase "transaction-service/cmd/transaction-service/application/use-case"
	"transaction-service/cmd/transaction-service/domain/command"

	"github.com/gorilla/mux"
)

type DeleteTransactionController struct {
	useCase *usecase.DeleteTransactionUseCase
}

func NewDeleteTransactionController(u *usecase.DeleteTransactionUseCase) *DeleteTransactionController {
	return &DeleteTransactionController{
		useCase: u,
	}
}

func (u *DeleteTransactionController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	request := &usecase.DeleteTransactionUseCaseRequest{
		ID:        id,
		AccountID: r.Context().Value("account_id").(string),
	}

	err := u.useCase.Execute(request)
	if err != nil {
		status := http.StatusInternalServerError

		if err == command.ErrTransactionDoesntExist {
			status = http.StatusNotFound
		}

		http.Error(w, err.Error(), status)
	}
}
