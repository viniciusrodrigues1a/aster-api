package event

import (
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Password struct {
	Hash string
}

type AccountWasCreatedEvent struct {
	Name  string
	Email string
	Password
}

func NewAccountWasCreatedEvent(name, email, hash string) *eventlib.BaseEvent {
	payload := AccountWasCreatedEvent{
		Name:  name,
		Email: email,
		Password: Password{
			Hash: hash,
		},
	}

	return eventlib.NewBaseEvent("account-was-created", primitive.NewObjectID(), payload)
}
