package command

import (
	"account-service/cmd/account-service/domain/event"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type hasherSpy struct {
	calledTimes int
}

func (h *hasherSpy) Hash(plaintext string) (string, error) {
	h.calledTimes += 1
	return "hash", nil
}

func TestCreateAccountCommand(t *testing.T) {
	cmd := CreateAccountCommand{
		Name:     "Amy",
		Email:    "amy@email.com",
		Password: "pa55",
		Hasher:   &hasherSpy{},
	}
	evt, err := cmd.Handle()
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}

	hash, _ := cmd.Hasher.Hash(cmd.Password)

	got := evt
	want := event.NewAccountWasCreatedEvent(cmd.Name, cmd.Email, hash)

	if !cmp.Equal(got, want, cmpopts.IgnoreFields(eventlib.BaseEvent{}, "Data.StreamId", "Data.Id")) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCreateAccountCommand_InvalidEmail(t *testing.T) {
	cmd := CreateAccountCommand{
		Name:     "Amy",
		Email:    "invalid.email",
		Password: "pa55",
		Hasher:   &hasherSpy{},
	}
	_, err := cmd.Handle()

	got := err
	want := ErrInvalidEmail

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCreateAccountCommand_CallsHasher(t *testing.T) {
	spy := &hasherSpy{}
	cmd := CreateAccountCommand{
		Name:     "Amy",
		Email:    "amy@email.com",
		Password: "pa55",
		Hasher:   spy,
	}
	_, err := cmd.Handle()
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}

	if spy.calledTimes != 1 {
		t.Errorf("called Hasher %d time(s), wanted 1 call", spy.calledTimes)
	}
}

type hasherErrorSpy struct {
	thrown error
}

func (h *hasherErrorSpy) Hash(plaintext string) (string, error) {
	h.thrown = fmt.Errorf("hasher error")
	return "", h.thrown
}

func TestCreateAccountCommand_ReturnHasherError(t *testing.T) {
	spy := &hasherErrorSpy{}
	cmd := CreateAccountCommand{
		Name:     "Amy",
		Email:    "amy@email.com",
		Password: "pa55",
		Hasher:   spy,
	}
	_, err := cmd.Handle()

	got := err
	want := spy.thrown

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
