package controller

import (
	"encoding/json"
	usecase "expense-service/cmd/expense-service/application/use-case"
	"expense-service/cmd/expense-service/domain/command"
	"net/http"

	"github.com/gorilla/mux"
)

type UpdateExpenseController struct {
	useCase *usecase.UpdateExpenseUseCase
}

func NewUpdateExpenseController(u *usecase.UpdateExpenseUseCase) *UpdateExpenseController {
	return &UpdateExpenseController{
		useCase: u,
	}
}

func (u *UpdateExpenseController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var body usecase.UpdateExpenseUseCaseRequest

	decoderErr := json.NewDecoder(r.Body).Decode(&body)
	if decoderErr != nil {
		http.Error(w, decoderErr.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	body.Id = id

	err := u.useCase.Execute(&body)
	if err != nil {
		status := http.StatusInternalServerError

		if err == command.ErrExpenseDoesntExist {
			status = http.StatusNotFound
		}

		http.Error(w, err.Error(), status)
	}
}
