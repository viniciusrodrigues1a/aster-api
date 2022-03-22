package controller

import (
	"encoding/json"
	"net/http"
	"product-service/cmd/product-service/application/use-case"
)

type CreateProductController struct {
	useCase *usecase.CreateProductUseCase
}

func NewCreateProductController(useCase *usecase.CreateProductUseCase) *CreateProductController {
	return &CreateProductController{
		useCase: useCase,
	}
}

func (c *CreateProductController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var body usecase.CreateProductUseCaseRequest

	decoderErr := json.NewDecoder(r.Body).Decode(&body)
	if decoderErr != nil {
		http.Error(w, decoderErr.Error(), http.StatusBadRequest)
		return
	}

	err := c.useCase.Execute(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
