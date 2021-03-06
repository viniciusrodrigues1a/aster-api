package controller

import (
	usecase "expense-service/cmd/expense-service/application/use-case"
	"expense-service/cmd/expense-service/domain/command"
	"net/http"

	"github.com/gorilla/mux"
)

type DeleteExpenseController struct {
	useCase *usecase.DeleteExpenseUseCase
}

func NewDeleteExpenseController(u *usecase.DeleteExpenseUseCase) *DeleteExpenseController {
	return &DeleteExpenseController{
		useCase: u,
	}
}

func (d *DeleteExpenseController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	request := &usecase.DeleteExpenseUseCaseRequest{
		ID:        id,
		AccountID: r.Context().Value("account_id").(string),
	}

	err := d.useCase.Execute(request)
	if err != nil {
		status := http.StatusInternalServerError

		if err == command.ErrExpenseDoesntExist {
			status = http.StatusNotFound
		}

		http.Error(w, err.Error(), status)
	}
}
