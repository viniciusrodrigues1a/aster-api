package projector

import (
	"account-service/cmd/account-service/domain/event"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type AccountCreationProjector struct{}

func (a *AccountCreationProjector) Project(e *eventlib.BaseEvent) *AccountState {
	payload := e.Payload.(event.AccountWasCreatedEvent)

	return &AccountState{
		Name:  payload.Name,
		Email: payload.Email,
	}
}
