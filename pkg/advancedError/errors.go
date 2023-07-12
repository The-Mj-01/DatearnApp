package advancedError

import (
	"fmt"
	"runtime/debug"
)

// publisher
// ExpertError defines custom type which is new error to system with additional and advanced information
type ExpertError struct {
	inner      error
	message    string
	stackTrace string
	observers  []observer
}

// New creates new error and returns it
func New(underlayingErr error, msg string) *ExpertError {
	return &ExpertError{
		inner:      underlayingErr,
		message:    msg,
		stackTrace: string(debug.Stack()),
	}
}

// GetMessage generates message in Error method that satisfies error interface and calling observers
func (e *ExpertError) GetMessage() string {
	return fmt.Sprintf("Error with inner message:%s\n with message:%s\n stack:%s", e.inner.Error(), e.message, e.stackTrace)
}

// Error satisfies error interface with the ability invoking Observers
func (e *ExpertError) Error() string {
	msg := e.GetMessage()
	e.dispatchObservers()
	return msg
}

// dispatchObservers Dispatches registered observers
func (e *ExpertError) dispatchObservers() {
	fLogger := fileLogger{}
	e.RegisterObserver(&fLogger)
	e.Notify()
}
