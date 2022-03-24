package factory

import (
	usecase "account-service/cmd/account-service/application/use-case"
	"account-service/cmd/account-service/external/messaging"
	"account-service/cmd/account-service/infrastructure/cryptography"
	jwtmanager "account-service/cmd/account-service/infrastructure/jwt-manager"
	"os"
)

var eventStoreRepository = makeMongoEventStoreRepository()
var stateStoreRepository = makeRedisStateStoreRepository()

func makeCreateAccountUseCase() *usecase.CreateAccountUseCase {
	m := messaging.New()
	emitter := messaging.NewAccountCreationEmitter(m)
	hasher := cryptography.NewHasher()
	return usecase.NewCreateAccountUseCase(emitter, eventStoreRepository, stateStoreRepository, stateStoreRepository, hasher)
}

func makeLoginUseCase() *usecase.LoginUseCase {
	hashComparer := cryptography.NewHashComparer()
	jwtSigner := jwtmanager.NewJWTManager([]byte(os.Getenv("RSA_PRIVATE_KEY")), []byte(os.Getenv("RSA_PUBLIC_KEY")))

	return usecase.NewLoginUseCase(stateStoreRepository, jwtSigner, hashComparer)
}

func makeValidateTokenUseCase() *usecase.ValidateTokenUseCase {
	jwtVerifier := jwtmanager.NewJWTManager([]byte(os.Getenv("RSA_PRIVATE_KEY")), []byte(os.Getenv("RSA_PUBLIC_KEY")))

	return usecase.NewValidateTokenUseCase(jwtVerifier)
}
