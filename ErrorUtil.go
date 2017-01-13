package common

import (
	"sync"
	"github.com/cnfree/common/errors"
	"github.com/cnfree/common/debug"
)

type errorUtil struct {
	mutex sync.Mutex
}

var Error = errorUtil{}

func (this errorUtil) NewError(err string) error {
	return errors.New(err)
}

func (this errorUtil) NewTraceError(err string) error {
	return errors.NewTraceError(err)
}

func (this errorUtil) TraceError(err error) error {
	return errors.TraceError(err)
}

func (this errorUtil) Errorf(format string, v ...interface{}) error {
	return errors.Errorf(format, v...)
}

func (this errorUtil) SetTraceEnable(trace bool) {
	errors.TraceEnabled = trace
}

func (this errorUtil) ErrorStack(err error) debug.Stack {
	return errors.ErrorStack(err)
}

func (this errorUtil) ErrorCause(err error) error {
	return errors.ErrorCause(err)
}
