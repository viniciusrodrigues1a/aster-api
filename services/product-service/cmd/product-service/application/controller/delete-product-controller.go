package controller

import (
	"net/http"
	usecase "product-service/cmd/product-service/application/use-case"
	"product-service/cmd/product-service/domain/command"

	"github.com/gorilla/mux"
)

type DeleteProductController struct {
	useCase *usecase.DeleteProductUseCase
}

func NewDeleteProductController(u *usecase.DeleteProductUseCase) *DeleteProductController {
	return &DeleteProductController{
		useCase: u,
	}
}

func (d *DeleteProductController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	request := &usecase.DeleteProductUseCaseRequest{
		ID:        id,
		AccountID: r.Context().Value("account_id").(string),
	}

	err := d.useCase.Execute(request)
	if err != nil {
		status := http.StatusInternalServerError

		if err == command.ErrProductNotFound {
			status = http.StatusNotFound
		}

		http.Error(w, err.Error(), status)
	}
}
