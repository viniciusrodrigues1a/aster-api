package command

import (
	"testing"
	"transaction-service/cmd/transaction-service/domain/event"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

func TestCreateTransactionCommand(t *testing.T) {
	productID := "product-id-0"
	cmd := CreateTransactionCommand{
		ProductID:               &productID,
		ValuePaid:               200,
		Quantity:                3,
		Description:             "My description",
		EventStoreStreamWriter:  &streamWriterSpy{},
		ProductStateStoreReader: &stateReaderSpy{returnValue: GetJSONOfProductWithPrice(100)},
	}
	evt, err := cmd.Handle()
	if err != nil {
		t.Errorf("got error: %s", err.Error())
	}

	got := evt
	want := event.NewTransactionWasCreatedEvent(cmd.ProductID, "open", cmd.Quantity, cmd.ValuePaid, cmd.Description)

	if !cmp.Equal(got, want, cmpopts.IgnoreFields(eventlib.BaseEvent{}, "Data.StreamId", "Data.Id")) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCreateTransactionCommand_CallsStreamWriterSpy(t *testing.T) {
	spy := &streamWriterSpy{}
	productID := "product-id-0"
	cmd := CreateTransactionCommand{
		ProductID:               &productID,
		ValuePaid:               200,
		Quantity:                3,
		Description:             "My description",
		EventStoreStreamWriter:  spy,
		ProductStateStoreReader: &stateReaderSpy{returnValue: GetJSONOfProductWithPrice(100)},
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
	productID := "product-id-0"
	cmd := CreateTransactionCommand{
		ProductID:               &productID,
		ValuePaid:               200,
		Quantity:                3,
		Description:             "My description",
		EventStoreStreamWriter:  spy,
		ProductStateStoreReader: &stateReaderSpy{returnValue: GetJSONOfProductWithPrice(100)},
	}
	_, err := cmd.Handle()

	got := err
	want := spy.thrown

	if err != spy.thrown {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCreateTransactionCommand_ReturnErrQuantityMustBeGreaterThanZero(t *testing.T) {
	productID := "product-id-0"
	cmd := CreateTransactionCommand{
		ProductID:               &productID,
		ValuePaid:               200,
		Quantity:                0,
		Description:             "My description",
		EventStoreStreamWriter:  &streamWriterSpy{},
		ProductStateStoreReader: &stateReaderSpy{returnValue: GetJSONOfProductWithPrice(100)},
	}
	_, err := cmd.Handle()

	got := err
	want := ErrQuantityMustBeGreaterThanZero

	if !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCreateTransactionCommand_ReturnErrValuePaidCantBeGreaterThanTotalValue(t *testing.T) {
	productID := "product-id-0"
	cmd := CreateTransactionCommand{
		ProductID:               &productID,
		ValuePaid:               400,
		Quantity:                3,
		Description:             "My description",
		EventStoreStreamWriter:  &streamWriterSpy{},
		ProductStateStoreReader: &stateReaderSpy{returnValue: GetJSONOfProductWithPrice(100)},
	}
	_, err := cmd.Handle()

	got := err
	want := ErrValuePaidCantBeGreaterThanTotalValue

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCreateTransactionCommand_StatusShouldBeClosed(t *testing.T) {
	productID := "product-id-0"
	quantity := int64(3)
	productSalePrice := int64(100)
	cmd := CreateTransactionCommand{
		ProductID:               &productID,
		ValuePaid:               productSalePrice * quantity,
		Quantity:                quantity,
		Description:             "My description",
		EventStoreStreamWriter:  &streamWriterSpy{},
		ProductStateStoreReader: &stateReaderSpy{returnValue: GetJSONOfProductWithPrice(productSalePrice)},
	}
	evt, err := cmd.Handle()
	if err != nil {
		t.Errorf("got error: %s", err.Error())
	}

	expectedEvent := event.NewTransactionWasCreatedEvent(cmd.ProductID, "closed", cmd.Quantity, cmd.ValuePaid, cmd.Description)

	got := evt.Payload.(*event.TransactionWasCreatedEvent).Status
	want := expectedEvent.Payload.(*event.TransactionWasCreatedEvent).Status

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
