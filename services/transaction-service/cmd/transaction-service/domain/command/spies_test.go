package command

import (
	"fmt"

	eventlib "github.com/viniciusrodrigues1a/aster-api/pkg/domain/event-lib"
)

type streamWriterSpy struct {
	calledTimes int
}

func (s *streamWriterSpy) StoreEventStream(event *eventlib.BaseEvent) (string, error) {
	s.calledTimes += 1
	return "", nil
}

type streamWriterErrorSpy struct {
	thrown error
}

func (s *streamWriterErrorSpy) StoreEventStream(event *eventlib.BaseEvent) (string, error) {
	s.thrown = fmt.Errorf("error storing event stream")
	return "", s.thrown
}

type storeWriterSpy struct {
	calledTimes int
}

func (s *storeWriterSpy) StoreEvent(event *eventlib.BaseEvent) (string, error) {
	s.calledTimes += 1
	return "", nil
}

type storeWriterErrorSpy struct {
	thrown error
}

func (s *storeWriterErrorSpy) StoreEvent(event *eventlib.BaseEvent) (string, error) {
	s.thrown = fmt.Errorf("error storing event")
	return "", s.thrown
}

type stateReaderSpy struct {
	calledTimes int
}

func (s *stateReaderSpy) ReadState(key string) (string, error) {
	s.calledTimes += 1
	return "", nil
}

type stateReaderErrorSpy struct {
	thrown error
}

func (s *stateReaderErrorSpy) ReadState(key string) (string, error) {
	s.thrown = ErrTransactionDoesntExist
	return "", ErrTransactionDoesntExist
}
