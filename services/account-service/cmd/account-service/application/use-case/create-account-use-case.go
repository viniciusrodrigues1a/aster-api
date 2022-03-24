package usecase

import (
	"account-service/cmd/account-service/domain/command"
	"account-service/cmd/account-service/domain/projector"
	"errors"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrEmailAlreadyInUse = errors.New("email is already in use")

type CreateAccountUseCase struct {
	messageEmitter   MessageEmitter
	eventStoreWriter eventstorelib.EventStoreStreamWriter
	stateStoreReader statestorelib.StateStoreReader
	stateStoreWriter statestorelib.StateStoreWriter
	hasher           command.Hasher
}

func NewCreateAccountUseCase(m MessageEmitter, evtStore eventstorelib.EventStoreStreamWriter, sttStoreW statestorelib.StateStoreWriter, sttStoreR statestorelib.StateStoreReader, hasher command.Hasher) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		messageEmitter:   m,
		eventStoreWriter: evtStore,
		stateStoreWriter: sttStoreW,
		stateStoreReader: sttStoreR,
		hasher:           hasher,
	}
}

type CreateAccountUseCaseRequest struct {
	Name     string
	Email    string
	Password string
}

type CreateInventoryCommand struct {
	Id primitive.ObjectID
}

func (c *CreateAccountUseCase) Execute(request *CreateAccountUseCaseRequest) error {
	command := command.NewCreateAccountCommand(request.Name, request.Email, request.Password, c.hasher)
	event, err := command.Handle()
	if err != nil {
		return err
	}

	if _, err := c.stateStoreReader.ReadState(request.Email); err == nil {
		return ErrEmailAlreadyInUse
	}

	id, err := c.eventStoreWriter.StoreEventStream(event)
	if err != nil {
		return err
	}

	projector := projector.AccountCreationProjector{}
	state := projector.Project(event)

	stateErr := c.stateStoreWriter.StoreState(request.Email, state)
	if stateErr != nil {
		return stateErr
	}

	oid, _ := primitive.ObjectIDFromHex(id)
	c.messageEmitter.Emit(CreateInventoryCommand{Id: oid})

	return nil
}
