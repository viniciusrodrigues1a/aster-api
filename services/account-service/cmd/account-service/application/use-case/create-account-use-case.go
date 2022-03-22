package usecase

import (
	"account-service/cmd/account-service/domain/command"
	"account-service/cmd/account-service/domain/projector"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateAccountUseCase struct {
	messageEmitter   MessageEmitter
	eventStoreWriter eventstorelib.EventStoreStreamWriter
	stateStoreWriter statestorelib.StateStoreWriter
}

func NewCreateAccountUseCase(m MessageEmitter, evtStore eventstorelib.EventStoreStreamWriter, sttStore statestorelib.StateStoreWriter) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		messageEmitter:   m,
		eventStoreWriter: evtStore,
		stateStoreWriter: sttStore,
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
	command := command.NewCreateAccountCommand(request.Name, request.Email, request.Password)
	event, err := command.Handle()
	if err != nil {
		return err
	}

	id, err := c.eventStoreWriter.StoreEventStream(event)
	if err != nil {
		return err
	}

	projector := projector.AccountCreationProjector{}
	state := projector.Project(event)

	stateErr := c.stateStoreWriter.StoreState(id, state)
	if stateErr != nil {
		return stateErr
	}

	oid, _ := primitive.ObjectIDFromHex(id)
	c.messageEmitter.Emit(CreateInventoryCommand{Id: oid})

	return nil
}
