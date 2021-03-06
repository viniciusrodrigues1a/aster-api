package command

import (
	"account-service/cmd/account-service/domain/event"
	"errors"
	"net/mail"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

var ErrInvalidEmail = errors.New("invalid email")
var ErrPasswordIsTooShort = errors.New("password must be at least 8 characters long")

type Hasher interface {
	Hash(plaintext string) (string, error)
}

type CreateAccountCommand struct {
	Name     string
	Email    string
	Password string
	Hasher   Hasher
}

func NewCreateAccountCommand(name, email, password string, hasher Hasher) *CreateAccountCommand {
	return &CreateAccountCommand{
		Name:     name,
		Email:    email,
		Password: password,
		Hasher:   hasher,
	}
}

// Handle hashes the password and stores an AccountWasCreatedEvent in the event store and returns the resulting event.
// returns an ErrInvalidEmail if email is not valid.
// returns an error if it can't hash the password.
func (c *CreateAccountCommand) Handle() (*eventlib.BaseEvent, error) {
	if !isPasswordValid(c.Password) {
		return nil, ErrPasswordIsTooShort
	}

	if !isEmailValid(c.Email) {
		return nil, ErrInvalidEmail
	}

	hash, err := c.Hasher.Hash(c.Password)
	if err != nil {
		return nil, err
	}

	return event.NewAccountWasCreatedEvent(c.Name, c.Email, hash), nil
}

func isPasswordValid(password string) bool {
	if len(password) < 8 {
		return false
	}

	return true
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
