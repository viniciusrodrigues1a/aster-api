package command

import (
	"expense-service/cmd/expense-service/domain/event"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

func TestDeleteExpenseCommand(t *testing.T) {
	cmd := DeleteExpenseCommand{
		Id:               "expense-id-0",
		EventStoreWriter: &storeWriterSpy{},
		StateStoreReader: &stateReaderSpy{},
	}
	evt, err := cmd.Handle()
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}

	got := evt
	want := event.NewExpenseWasDeletedEvent(cmd.Id)

	if !cmp.Equal(got, want, cmpopts.IgnoreFields(eventlib.BaseEvent{}, "Data.Id")) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestDeleteExpenseCommand_CallsStoreWriterSpy(t *testing.T) {
	spy := &storeWriterSpy{}
	cmd := DeleteExpenseCommand{
		Id:               "expense-id-0",
		EventStoreWriter: spy,
		StateStoreReader: &stateReaderSpy{},
	}
	_, err := cmd.Handle()
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}

	if spy.calledTimes != 1 {
		t.Errorf("called EventStoreWriter %d time(s), wanted 1 call", spy.calledTimes)
	}
}

func TestDeleteExpenseCommand_CallsStateReaderSpy(t *testing.T) {
	spy := &stateReaderSpy{}
	cmd := DeleteExpenseCommand{
		Id:               "expense-id-0",
		EventStoreWriter: &storeWriterSpy{},
		StateStoreReader: spy,
	}
	_, err := cmd.Handle()
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}

	if spy.calledTimes != 1 {
		t.Errorf("called StateStoreReader %d time(s), wanted 1 call", spy.calledTimes)
	}
}

func TestDeleteExpenseCommand_ReturnStoreWriterError(t *testing.T) {
	spy := &storeWriterErrorSpy{}
	cmd := DeleteExpenseCommand{
		Id:               "expense-id-0",
		EventStoreWriter: spy,
		StateStoreReader: &stateReaderSpy{},
	}
	_, err := cmd.Handle()

	got := err
	want := spy.thrown

	if err != spy.thrown {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDeleteExpenseCommand_ReturnStateReaderError(t *testing.T) {
	spy := &stateReaderErrorSpy{}
	cmd := DeleteExpenseCommand{
		Id:               "expense-id-0",
		EventStoreWriter: &storeWriterSpy{},
		StateStoreReader: spy,
	}
	_, err := cmd.Handle()

	got := err
	want := spy.thrown

	if err != spy.thrown {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
