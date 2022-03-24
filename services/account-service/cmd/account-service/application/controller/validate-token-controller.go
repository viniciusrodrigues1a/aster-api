package controller

import (
	usecase "account-service/cmd/account-service/application/use-case"
	"errors"
	"net/http"
	"strings"
)

var ErrInvalidToken = errors.New("invalid authorization token")

type ValidateTokenController struct {
	useCase *usecase.ValidateTokenUseCase
}

func NewValidateTokenController(u *usecase.ValidateTokenUseCase) *ValidateTokenController {
	return &ValidateTokenController{
		useCase: u,
	}
}

func (v *ValidateTokenController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	authorizationHeader, ok := r.Header["Authorization"]
	if !ok {
		http.Error(w, ErrInvalidToken.Error(), http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorizationHeader[0], " ")[1]

	err := v.useCase.Execute(&usecase.ValidateTokenUseCaseRequest{Token: token})
	if err != nil {
		http.Error(w, ErrInvalidToken.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
