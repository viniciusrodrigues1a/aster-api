package command

import (
	"expense-service/cmd/expense-service/domain/event"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

func TestCreateExpenseCommand(t *testing.T) {
	cmd := CreateExpenseCommand{
		Title:                  "My expense",
		Description:            "My description",
		Value:                  300,
		EventStoreStreamWriter: &streamWriterSpy{},
	}
	evt, err := cmd.Handle()
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}

	got := evt
	want := event.NewExpenseWasCreatedEvent(cmd.Title, cmd.Description, cmd.Value)

	if !cmp.Equal(got, want, cmpopts.IgnoreFields(eventlib.BaseEvent{}, "Data.StreamId", "Data.Id")) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCreateExpenseCommand_CallsStreamWriterSpy(t *testing.T) {
	spy := &streamWriterSpy{}
	cmd := CreateExpenseCommand{
		Title:                  "My expense",
		Description:            "My description",
		Value:                  300,
		EventStoreStreamWriter: spy,
	}
	_, err := cmd.Handle()
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}

	if spy.calledTimes != 1 {
		t.Errorf("called EventStoreStreamWriter %d time(s), wanted 1 call", spy.calledTimes)
	}
}

func TestCreateExpenseCommand_ReturnStreamWriterError(t *testing.T) {
	spy := &streamWriterErrorSpy{}
	cmd := CreateExpenseCommand{
		Title:                  "My expense",
		Description:            "My description",
		Value:                  300,
		EventStoreStreamWriter: spy,
	}
	_, err := cmd.Handle()

	got := err
	want := spy.thrown

	if err != spy.thrown {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
