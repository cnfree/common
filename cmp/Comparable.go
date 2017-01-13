package compare

import (
	"errors"
	"reflect"
	"unsafe"
	"fmt"
)

type Comparable interface {
	CompareTo(Comparable) (int, error)
}
type IntComp int
type Int64Comp int64
type FloatComp float32
type Float64Comp float64
type StringComp string

func (this StringComp) String() string {
	return string(this)
}

func (this StringComp) Join(str interface{}) StringComp {
	return StringComp(string(this) + fmt.Sprint(str))
}

var ErrInvalid = errors.New("invalid argument")

type CompareError struct {
	LeftType  string
	RightType string
	Err       error
}

type ConvertError struct {
	Type        string
	ConvertType string
	Err         error
}

func (e *CompareError) Error() string { return e.LeftType + " " + e.RightType + ": " + e.Err.Error() }
func (e *ConvertError) Error() string { return e.Type + " " + e.ConvertType + ": " + e.Err.Error() }

func (this IntComp) CompareTo(o Comparable) (int, error) {
	return compareTo(this, o)
}

func (this Int64Comp) CompareTo(o Comparable) (int, error) {
	return compareTo(this, o)
}

func (this FloatComp) CompareTo(o Comparable) (int, error) {
	return compareTo(this, o)
}

func (this Float64Comp) CompareTo(o Comparable) (int, error) {
	return compareTo(this, o)
}

func (this StringComp) CompareTo(o Comparable) (int, error) {
	return compareTo(this, o)
}

func compareTo(o1 Comparable, o2 Comparable) (int, error) {
	if canCompare(o1, o2) {
		return compare(o1, o2)
	} else {
		return 0, &CompareError{reflect.TypeOf(o1).Name(), reflect.TypeOf(o2).Name(), ErrInvalid}
	}
}

func canCompare(o1 Comparable, o2 Comparable) bool {
	if isSameType(o1, o2) {
		return true
	} else if isNumber(o1) && isNumber(o2) {
		return true
	}
	return false
}

func isNumber(comparable Comparable) bool {
	switch reflect.TypeOf(comparable).Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func isString(comparable Comparable) bool {
	switch reflect.TypeOf(comparable).Kind() {
	case reflect.String:
		return true
	default:
		return false
	}
}

func isSameType(o1 Comparable, o2 Comparable) bool {
	return reflect.TypeOf(o1).Kind() == reflect.TypeOf(o2).Kind()
}

func compare(o1 Comparable, o2 Comparable) (int, error) {
	if isNumber(o1) && isNumber(o2) {
		number1, err := toFloat64(o1)
		if err != nil {
			return 0, nil
		}
		number2, err := toFloat64(o2)
		if err != nil {
			return 0, nil
		}
		if number1 > number2 {
			return 1, nil
		} else if number1 == number2 {
			return 0, nil
		} else {
			return -1, nil
		}
	} else if isString(o1) && isString(o2) {
		string1, err := toString(o1)
		if err != nil {
			return 0, nil
		}
		string2, err := toString(o2)
		if err != nil {
			return 0, nil
		}
		if string1 > string2 {
			return 1, nil
		} else if string1 == string2 {
			return 0, nil
		} else {
			return -1, nil
		}
	}
	return 0, &CompareError{reflect.TypeOf(o1).Name(), reflect.TypeOf(o2).Name(), ErrInvalid}
}

func toFloat64(comparable Comparable) (float64, error) {
	switch value := comparable.(type) {
	case IntComp:
		return float64(int(value)), nil
	case Int64Comp:
		return float64(int64(value)), nil
	case FloatComp:
		return float64(float32(value)), nil
	case Float64Comp:
		return float64(float64(value)), nil
	default:
		return 0, &ConvertError{reflect.ValueOf(comparable).Type().Name(), "float64", ErrInvalid}
	}
}

func toString(comparable Comparable) (string, error) {
	switch value := comparable.(type) {
	case StringComp:
		return string(value), nil
	default:
		return "", &ConvertError{reflect.ValueOf(comparable).Type().Name(), "string", ErrInvalid}
	}
}

func ConvertIntArrayToComparable(array []int) []Comparable {
	value := *(*[]IntComp)(unsafe.Pointer(&array))
	return convertToComparable(reflect.ValueOf(value))
}

func convertToComparable(a reflect.Value) []Comparable {
	b := make([]Comparable, a.Len())
	for i := 0; i < a.Len(); i++ {
		c := a.Index(i).Interface().(Comparable)
		b[i] = c
	}
	return b
}

func ConvertInt64ArrayToComparable(array []int64) []Comparable {
	value := *(*[]Int64Comp)(unsafe.Pointer(&array))
	return convertToComparable(reflect.ValueOf(value))
}

func ConvertFloatArrayToComparable(array []float32) []Comparable {
	value := *(*[]FloatComp)(unsafe.Pointer(&array))
	return convertToComparable(reflect.ValueOf(value))
}

func ConvertFloat64ArrayToComparable(array []float32) []Comparable {
	value := *(*[]Float64Comp)(unsafe.Pointer(&array))
	return convertToComparable(reflect.ValueOf(value))
}

func ConvertStringArrayToComparable(array []string) []Comparable {
	value := *(*[]StringComp)(unsafe.Pointer(&array))
	return convertToComparable(reflect.ValueOf(value))
}

func Map(s []Comparable, handle func(Comparable) (Comparable, error)) ([]Comparable, error) {
	out := []Comparable{}
	for _, i := range s {
		result, err := handle(i)
		if err != nil {
			return nil, err
		}
		out = append(out, result.(Comparable))
	}
	return out, nil
}

func Reduce(s []Comparable, handle func(Comparable, Comparable) (Comparable, error)) (Comparable, error) {
	var out Comparable

	if len(s) > 0 {
		out = s[0]
	}
	for index, i := range s {
		if index == 0 {
			continue
		}
		result, err := handle(out, i)
		if err != nil {
			return out, err
		}
		out = result.(Comparable)
	}
	return out, nil
}

func Filter(s []Comparable, handle func(Comparable) (bool, error)) ([]Comparable, error) {
	out := []Comparable{}

	for _, i := range s {
		result, err := handle(i)
		if err != nil {
			return nil, err
		}
		if ! result {
			continue
		}
		out = append(out, i)
	}
	return out, nil
}

func max(c1 Comparable, c2 Comparable) (Comparable, error) {
	result, err := c1.CompareTo(c2)
	if err != nil {
		return nil, err
	}
	if result >= 0 {
		return c1, nil
	} else {
		return c2, nil
	}
}

func min(c1 Comparable, c2 Comparable) (Comparable, error) {
	result, err := c1.CompareTo(c2)
	if err != nil {
		return nil, err
	}
	if result <= 0 {
		return c1, nil
	} else {
		return c2, nil
	}
}

func Max(s []Comparable) (Comparable, error) {
	return Reduce(s, max)
}

func Min(s []Comparable) (Comparable, error) {
	return Reduce(s, min)
}
