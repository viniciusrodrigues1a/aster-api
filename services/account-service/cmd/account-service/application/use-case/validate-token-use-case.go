package usecase

type TokenVerifier interface {
	Verify(token string) (interface{}, error)
}

type ValidateTokenUseCase struct {
	tokenVerifier TokenVerifier
}

func NewValidateTokenUseCase(tokenVerifier TokenVerifier) *ValidateTokenUseCase {
	return &ValidateTokenUseCase{
		tokenVerifier: tokenVerifier,
	}
}

type ValidateTokenUseCaseRequest struct {
	Token string
}

func (v *ValidateTokenUseCase) Execute(request *ValidateTokenUseCaseRequest) error {
	_, err := v.tokenVerifier.Verify(request.Token)
	return err
}
