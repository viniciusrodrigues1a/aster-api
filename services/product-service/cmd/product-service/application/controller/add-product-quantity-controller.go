package controller

import (
	"encoding/json"
	"net/http"
	usecase "product-service/cmd/product-service/application/use-case"
	"product-service/cmd/product-service/domain/command"

	"github.com/gorilla/mux"
)

type AddProductQuantityController struct {
	useCase *usecase.AddProductQuantityUseCase
}

func NewAddProductQuantityController(u *usecase.AddProductQuantityUseCase) *AddProductQuantityController {
	return &AddProductQuantityController{
		useCase: u,
	}
}

func (a *AddProductQuantityController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var body usecase.AddProductQuantityUseCaseRequest

	decoderErr := json.NewDecoder(r.Body).Decode(&body)
	if decoderErr != nil {
		http.Error(w, decoderErr.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	body.ID = id
	body.AccountID = r.Context().Value("account_id").(string)

	err := a.useCase.Execute(&body)
	if err != nil {
		status := http.StatusInternalServerError

		if err == command.ErrProductNotFound {
			status = http.StatusNotFound
		}

		http.Error(w, err.Error(), status)
	}
}
