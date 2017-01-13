package common

import (
	"reflect"
	"sync"
	"github.com/cnfree/common/structs"
)

type reflectUtil struct {
	mutex sync.Mutex
}

var Reflect = reflectUtil{}

func (this reflectUtil) IsStruct(s interface{}) bool {
	return structs.IsStruct(s)
}

func (this reflectUtil) GetStructMetaData(s interface{}) *structs.Struct {
	if structs.IsStruct(s) {
		return structs.New(s)
	}
	return nil
}

func (this reflectUtil) GetMethodMetadata(method reflect.Method) *structs.MethodMetadata {
	m, err := structs.GetMethodMetadata(method)
	if err != nil {
		return nil
	} else {
		return &m
	}
}

func (this reflectUtil) IsZero(s interface{}) bool {
	v := reflect.ValueOf(s)
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && this.IsZero(v.Index(i))
		}
		return z
	case reflect.Struct:
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && this.IsZero(v.Field(i))
		}
		return z
	}

	z := reflect.Zero(v.Type())
	return v.Interface() == z.Interface()
}

// DeepEqual returns if the two values are deeply equal like reflect.DeepEqual.
// In addition to this, this method will also dereference the input values if
// possible so the DeepEqual performed will not fail if one parameter is a
// pointer and the other is not.
//
// DeepEqual will not perform indirection of nested values of the input parameters.
func (this reflectUtil) DeepEqual(a, b interface{}) bool {
	ra := reflect.Indirect(reflect.ValueOf(a))
	rb := reflect.Indirect(reflect.ValueOf(b))

	if raValid, rbValid := ra.IsValid(), rb.IsValid(); !raValid && !rbValid {
		// If the elements are both nil, and of the same type the are equal
		// If they are of different types they are not equal
		return reflect.TypeOf(a) == reflect.TypeOf(b)
	} else if raValid != rbValid {
		// Both values must be valid to be equal
		return false
	}

	return reflect.DeepEqual(ra.Interface(), rb.Interface())
}
