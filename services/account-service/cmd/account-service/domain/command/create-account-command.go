package command

import (
	"account-service/cmd/account-service/domain/event"
	"errors"
	"net/mail"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidEmail = errors.New("invalid email")

type CreateAccountCommand struct {
	Name     string
	Email    string
	Password string
}

func NewCreateAccountCommand(name, email, password string) *CreateAccountCommand {
	return &CreateAccountCommand{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func (c *CreateAccountCommand) Handle() (*eventlib.BaseEvent, error) {
	if !isEmailValid(c.Email) {
		return nil, ErrInvalidEmail
	}

	hash, err := hashPassword(c.Password)
	if err != nil {
		return nil, err
	}

	return event.NewAccountWasCreatedEvent(c.Name, c.Email, hash), nil
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func hashPassword(plaintext string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plaintext), 14)
	return string(bytes), err
}
