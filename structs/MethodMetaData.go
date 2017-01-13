package structs

import (
	"reflect"
	"runtime"
	"strings"
	"errors"
)

var errMethodIsNil = errors.New("method is nil")
var errBadMethodType = errors.New("method type error")

type MethodMetadata struct {
	Name   string
	File   string
	Line   int
	Method reflect.Method
}

func GetMethodMetadata(method reflect.Method) (metadata MethodMetadata, err error) {
	if method.Func.IsNil() {
		err = errMethodIsNil
		return
	}

	iMethod := method.Func.Interface()

	if reflect.TypeOf(iMethod).Kind() != reflect.Func {
		err = errBadMethodType
		return
	}

	v := reflect.ValueOf(iMethod)

	pc := runtime.FuncForPC(v.Pointer())

	name := strings.TrimRight(pc.Name(), "-fm")
	name = strings.Replace(name, "*", "", 1)

	metadata.Method = method
	metadata.Name = name
	metadata.File, metadata.Line = pc.FileLine(v.Pointer())

	return
}

func (this MethodMetadata) Invoke(args ... interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	return this.Method.Func.Call(inputs)
}

func (this MethodMetadata) SimpleName() string {
	return this.Method.Name
}

func (this MethodMetadata) FullName() string {
	return this.Name
}
