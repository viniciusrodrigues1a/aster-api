package controller

import (
	"encoding/json"
	"net/http"
	"product-service/cmd/product-service/application/use-case"

	"github.com/gorilla/mux"
)

type UpdateProductController struct {
	useCase *usecase.UpdateProductUseCase
}

func NewUpdateProductController(useCase *usecase.UpdateProductUseCase) *UpdateProductController {
	return &UpdateProductController{
		useCase: useCase,
	}
}

func (u *UpdateProductController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var body usecase.UpdateProductUseCaseRequest

	decoderErr := json.NewDecoder(r.Body).Decode(&body)
	if decoderErr != nil {
		http.Error(w, decoderErr.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	body.Id = id
	body.AccountID = r.Context().Value("account_id").(string)

	err := u.useCase.Execute(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
