package command

import (
	"testing"
	"transaction-service/cmd/transaction-service/domain/event"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

func TestCreateTransactionCommand(t *testing.T) {
	cmd := CreateTransactionCommand{
		ValuePaid:              10000,
		Quantity:               3,
		Description:            "My description",
		EventStoreStreamWriter: &streamWriterSpy{},
	}
	evt, err := cmd.Handle()
	if err != nil {
		t.Errorf("got error: %s", err.Error())
	}

	got := evt
	want := event.NewTransactionWasCreatedEvent(cmd.Quantity, cmd.ValuePaid, cmd.Description)

	if !cmp.Equal(got, want, cmpopts.IgnoreFields(eventlib.BaseEvent{}, "Data.StreamId", "Data.Id")) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCreateTransactionCommand_CallsStreamWriterSpy(t *testing.T) {
	spy := &streamWriterSpy{}
	cmd := CreateTransactionCommand{
		ValuePaid:              10000,
		Quantity:               1,
		Description:            "My description",
		EventStoreStreamWriter: spy,
	}
	_, err := cmd.Handle()
	if err != nil {
		t.Errorf("got error: %s", err.Error())
	}

	if spy.calledTimes != 1 {
		t.Errorf("called EventStoreStreamWriter %d time(s), wanted 1 call", spy.calledTimes)
	}
}

func TestCreateTransactionCommand_ReturnStreamWriterError(t *testing.T) {
	spy := &streamWriterErrorSpy{}
	cmd := CreateTransactionCommand{
		ValuePaid:              10000,
		Quantity:               1,
		Description:            "My description",
		EventStoreStreamWriter: spy,
	}
	_, err := cmd.Handle()

	got := err
	want := spy.thrown

	if err != spy.thrown {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCreateTransactionCommand_ReturnErrQuantityMustBeGreaterThanZero(t *testing.T) {
	cmd := CreateTransactionCommand{
		ValuePaid:              10000,
		Quantity:               0,
		Description:            "My description",
		EventStoreStreamWriter: &streamWriterSpy{},
	}
	_, err := cmd.Handle()

	got := err
	want := ErrQuantityMustBeGreaterThanZero

	if !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Errorf("got %q, want %q", got, want)
	}
}
