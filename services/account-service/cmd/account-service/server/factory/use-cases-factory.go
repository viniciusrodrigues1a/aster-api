package factory

import (
	usecase "account-service/cmd/account-service/application/use-case"
	"account-service/cmd/account-service/external/messaging"
)

var eventStoreRepository = makeMongoEventStoreRepository()
var stateStoreRepository = makeRedisStateStoreRepository()

func makeCreateAccountUseCase() *usecase.CreateAccountUseCase {
	m := messaging.New()
	emitter := messaging.NewAccountCreationEmitter(m)
	return usecase.NewCreateAccountUseCase(emitter, eventStoreRepository, stateStoreRepository)
}
