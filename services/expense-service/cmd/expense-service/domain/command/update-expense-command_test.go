package command

import (
	"expense-service/cmd/expense-service/domain/event"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

func TestUpdateExpenseCommand(t *testing.T) {
	productID := "product-id-0"
	cmd := UpdateExpenseCommand{
		ProductID:               &productID,
		ID:                      "expense-id-0",
		Title:                   "My expense",
		Description:             "My description",
		Value:                   300,
		EventStoreWriter:        &storeWriterSpy{},
		StateStoreReader:        &stateReaderSpy{},
		ProductStateStoreReader: &stateReaderSpy{},
	}
	evt, err := cmd.Handle()
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}

	got := evt
	want := event.NewExpenseWasUpdatedEvent(cmd.ProductID, cmd.Title, cmd.Description, cmd.Value, cmd.ID)

	if !cmp.Equal(got, want, cmpopts.IgnoreFields(eventlib.BaseEvent{}, "Data.Id")) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestUpdateExpenseCommand_CallsStoreWriterSpy(t *testing.T) {
	spy := &storeWriterSpy{}
	cmd := UpdateExpenseCommand{
		ID:               "expense-id-0",
		Title:            "My expense",
		Description:      "My description",
		Value:            300,
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

func TestUpdateExpenseCommand_CallsStateReaderSpy(t *testing.T) {
	spy := &stateReaderSpy{}
	cmd := UpdateExpenseCommand{
		ID:               "expense-id-0",
		Title:            "My expense",
		Description:      "My description",
		Value:            300,
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

func TestUpdateExpenseCommand_ReturnStoreWriterError(t *testing.T) {
	spy := &storeWriterErrorSpy{}
	cmd := UpdateExpenseCommand{
		ID:               "expense-id-0",
		Title:            "My expense",
		Description:      "My description",
		Value:            300,
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

func TestUpdateExpenseCommand_ReturnStateReaderError(t *testing.T) {
	spy := &stateReaderErrorSpy{}
	cmd := UpdateExpenseCommand{
		ID:               "expense-id-0",
		Title:            "My expense",
		Description:      "My description",
		Value:            300,
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
