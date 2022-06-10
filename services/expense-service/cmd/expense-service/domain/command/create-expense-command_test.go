package command

import (
	"expense-service/cmd/expense-service/domain/event"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

func TestCreateExpenseCommand(t *testing.T) {
	productID := "product-id-0"
	cmd := CreateExpenseCommand{
		ProductID:               &productID,
		Title:                   "My expense",
		Description:             "My description",
		Value:                   300,
		EventStoreStreamWriter:  &streamWriterSpy{},
		ProductStateStoreReader: &stateReaderSpy{},
	}
	evt, err := cmd.Handle()
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}

	got := evt
	want := event.NewExpenseWasCreatedEvent(&productID, cmd.Title, cmd.Description, cmd.Value)

	if !cmp.Equal(got, want, cmpopts.IgnoreFields(eventlib.BaseEvent{}, "Data.StreamId", "Data.Id")) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCreateExpenseCommand_CallsStreamWriterSpy(t *testing.T) {
	spy := &streamWriterSpy{}
	cmd := CreateExpenseCommand{
		Title:                   "My expense",
		Description:             "My description",
		Value:                   300,
		EventStoreStreamWriter:  spy,
		ProductStateStoreReader: &stateReaderSpy{},
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
		Title:                   "My expense",
		Description:             "My description",
		Value:                   300,
		EventStoreStreamWriter:  spy,
		ProductStateStoreReader: &stateReaderSpy{},
	}
	_, err := cmd.Handle()

	got := err
	want := spy.thrown

	if err != spy.thrown {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCreateExpenseCommand_ReturnsErrTitleIsRequired(t *testing.T) {
	cmd := CreateExpenseCommand{
		Description:             "My description",
		Value:                   300,
		EventStoreStreamWriter:  &streamWriterSpy{},
		ProductStateStoreReader: &stateReaderSpy{},
	}
	_, err := cmd.Handle()

	got := err
	want := ErrTitleIsRequired

	if !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCreateExpenseCommand_ReturnsErrProductCouldntBeFound(t *testing.T) {
	spy := &stateReaderErrorSpy{thrown: ErrProductCouldntBeFound}
	productID := "product-id-0"
	cmd := CreateExpenseCommand{
		ProductID:               &productID,
		Title:                   "My expense",
		Description:             "My description",
		Value:                   300,
		EventStoreStreamWriter:  &streamWriterSpy{},
		ProductStateStoreReader: spy,
	}
	_, err := cmd.Handle()

	got := err
	want := spy.thrown

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
