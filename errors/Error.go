package errors

import (
	"fmt"
	"github.com/cnfree/common/debug"
	errs "errors"
)

var TraceEnabled = true

type Error struct {
	Stack debug.Stack
	Cause error
}

func (e *Error) Error() string {
	return e.Cause.Error()
}

func New(s string) error {
	return errs.New(s)
}

func NewTraceError(s string) error {
	err := errs.New(s)
	if !TraceEnabled {
		return err
	}
	return &Error{
		Stack: debug.TraceN(1, 32),
		Cause: err,
	}
}

func TraceError(err error) error {
	if err == nil || !TraceEnabled {
		return err
	}
	_, ok := err.(*Error)
	if ok {
		return err
	}
	return &Error{
		Stack: debug.TraceN(1, 32),
		Cause: err,
	}
}

func Errorf(format string, v ...interface{}) error {
	err := fmt.Errorf(format, v...)
	if !TraceEnabled {
		return err
	}
	return &Error{
		Stack: debug.TraceN(1, 32),
		Cause: err,
	}
}

func ErrorStack(err error) debug.Stack {
	if err == nil {
		return nil
	}
	e, ok := err.(*Error)
	if ok {
		return e.Stack
	}
	return nil
}

func ErrorCause(err error) error {
	for err != nil {
		e, ok := err.(*Error)
		if ok {
			err = e.Cause
		} else {
			return err
		}
	}
	return nil
}

func Equal(err1, err2 error) bool {
	e1 := ErrorCause(err1)
	e2 := ErrorCause(err2)
	if e1 == e2 {
		return true
	}
	if e1 == nil || e2 == nil {
		return e1 == e2
	}
	return e1.Error() == e2.Error()
}

func NotEqual(err1, err2 error) bool {
	return !Equal(err1, err2)
}
