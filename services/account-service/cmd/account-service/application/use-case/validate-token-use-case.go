package usecase

import (
	"fmt"

	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

type TokenVerifier interface {
	Verify(token string) (interface{}, error)
}

type ValidateTokenUseCase struct {
	stateStoreReader statestorelib.StateStoreReader
	tokenVerifier    TokenVerifier
}

func NewValidateTokenUseCase(sttStore statestorelib.StateStoreReader, tokenVerifier TokenVerifier) *ValidateTokenUseCase {
	return &ValidateTokenUseCase{
		stateStoreReader: sttStore,
		tokenVerifier:    tokenVerifier,
	}
}

type ValidateTokenUseCaseRequest struct {
	Token string
}

func (v *ValidateTokenUseCase) Execute(request *ValidateTokenUseCaseRequest) error {
	payload, err := v.tokenVerifier.Verify(request.Token)
	if err != nil {
		return err
	}

	_, readErr := v.stateStoreReader.ReadState(payload.(string))
	if readErr != nil {
		return fmt.Errorf("validateToken: account not found")
	}

	return nil
}
