package command

import (
	"testing"
	"transaction-service/cmd/transaction-service/domain/event"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

func TestUpdateTransactionCommand(t *testing.T) {
	productID := "product-id-0"
	cmd := UpdateTransactionCommand{
		ProductID:               &productID,
		ID:                      "transaction-id-0",
		ValuePaid:               300,
		Quantity:                3,
		Description:             "My description",
		EventStoreWriter:        &storeWriterSpy{},
		StateStoreReader:        &stateReaderSpy{},
		ProductStateStoreReader: &stateReaderSpy{returnValue: GetJSONOfProductWithPrice(100)},
	}
	evt, err := cmd.Handle()
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}

	got := evt
	want := event.NewTransactionWasUpdatedEvent(cmd.ProductID, "closed", cmd.Quantity, cmd.ValuePaid, cmd.Description, cmd.ID)

	if !cmp.Equal(got, want, cmpopts.IgnoreFields(eventlib.BaseEvent{}, "Data.Id")) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestUpdateTransactionCommand_CallsStoreWriterSpy(t *testing.T) {
	spy := &storeWriterSpy{}
	cmd := UpdateTransactionCommand{
		ID:                      "transaction-id-0",
		ValuePaid:               300,
		Quantity:                3,
		Description:             "My description",
		EventStoreWriter:        spy,
		StateStoreReader:        &stateReaderSpy{returnValue: "{ \"product_id\": \"product-id-0\" }"},
		ProductStateStoreReader: &stateReaderSpy{returnValue: GetJSONOfProductWithPrice(100)},
	}
	_, err := cmd.Handle()
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}

	if spy.calledTimes != 1 {
		t.Errorf("called EventStoreWriter %d time(s), wanted 1 call", spy.calledTimes)
	}
}

func TestUpdateTransactionCommand_CallsStateReaderSpy(t *testing.T) {
	spy := &stateReaderSpy{returnValue: "{ \"product_id\": \"product-id-0\" }"}
	cmd := UpdateTransactionCommand{
		ID:                      "transaction-id-0",
		ValuePaid:               300,
		Quantity:                3,
		Description:             "My description",
		EventStoreWriter:        &storeWriterSpy{},
		StateStoreReader:        spy,
		ProductStateStoreReader: &stateReaderSpy{returnValue: GetJSONOfProductWithPrice(100)},
	}
	_, err := cmd.Handle()
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}

	if spy.calledTimes != 1 {
		t.Errorf("called StateStoreReader %d time(s), wanted 1 call", spy.calledTimes)
	}
}

func TestUpdateTransactionCommand_ReturnStoreWriterError(t *testing.T) {
	spy := &storeWriterErrorSpy{}
	cmd := UpdateTransactionCommand{
		ID:                      "transaction-id-0",
		ValuePaid:               300,
		Quantity:                3,
		Description:             "My description",
		EventStoreWriter:        spy,
		StateStoreReader:        &stateReaderSpy{returnValue: "{ \"product_id\": \"product-id-0\" }"},
		ProductStateStoreReader: &stateReaderSpy{returnValue: GetJSONOfProductWithPrice(100)},
	}
	_, err := cmd.Handle()

	got := err
	want := spy.thrown

	if err != spy.thrown {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestUpdateTransactionCommand_ReturnStateReaderError(t *testing.T) {
	spy := &stateReaderErrorSpy{}
	cmd := UpdateTransactionCommand{
		ID:                      "transaction-id-0",
		ValuePaid:               300,
		Quantity:                3,
		Description:             "My description",
		EventStoreWriter:        &storeWriterSpy{},
		StateStoreReader:        spy,
		ProductStateStoreReader: &stateReaderSpy{returnValue: GetJSONOfProductWithPrice(100)},
	}
	_, err := cmd.Handle()

	got := err
	want := spy.thrown

	if err != spy.thrown {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
