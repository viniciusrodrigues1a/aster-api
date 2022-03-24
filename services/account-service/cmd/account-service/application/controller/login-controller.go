package controller

import (
	usecase "account-service/cmd/account-service/application/use-case"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginController struct {
	useCase *usecase.LoginUseCase
}

func NewLoginController(u *usecase.LoginUseCase) *LoginController {
	return &LoginController{
		useCase: u,
	}
}

func (l *LoginController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var body usecase.LoginUseCaseRequest

	decodeErr := json.NewDecoder(r.Body).Decode(&body)
	if decodeErr != nil {
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		return
	}

	token, err := l.useCase.Execute(&body)
	if err != nil {
		var status = http.StatusInternalServerError

		if err == usecase.ErrInvalidCredentials {
			status = http.StatusUnauthorized
		}

		http.Error(w, err.Error(), status)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, `{ "token": "%s" }`, token)
}
