package usecase

import (
	"account-service/cmd/account-service/domain/projector"
	"encoding/json"
	"errors"

	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type TokenSigner interface {
	Sign(email string) (string, error)
}

type HashComparer interface {
	Compare(plaintext, hash string) bool
}

type LoginUseCase struct {
	stateStoreReader statestorelib.StateStoreReader
	tokenSigner      TokenSigner
	hashComparer     HashComparer
}

func NewLoginUseCase(sttStore statestorelib.StateStoreReader, tokenSigner TokenSigner, hashComparer HashComparer) *LoginUseCase {
	return &LoginUseCase{
		stateStoreReader: sttStore,
		tokenSigner:      tokenSigner,
		hashComparer:     hashComparer,
	}
}

type LoginUseCaseRequest struct {
	Email    string
	Password string
}

func (l *LoginUseCase) Execute(request *LoginUseCaseRequest) (string, error) {
	stateJSON, err := l.stateStoreReader.ReadState(request.Email)
	if err != nil { // email not found
		return "", ErrInvalidCredentials
	}

	var state projector.AccountState

	unmarshalErr := json.Unmarshal([]byte(stateJSON.(string)), &state)
	if unmarshalErr != nil {
		return "", unmarshalErr
	}

	isPasswordCorrect := l.hashComparer.Compare(request.Password, state.Hash)
	if !isPasswordCorrect {
		return "", ErrInvalidCredentials
	}

	return l.tokenSigner.Sign(request.Email)
}
