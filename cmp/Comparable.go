package cmp

import (
	"errors"
	"reflect"
	"unsafe"
	"fmt"
	"strconv"
)

var _ Comparable = IntComp(0)
var _ Comparable = Int64Comp(0)
var _ Comparable = FloatComp(0)
var _ Comparable = Float64Comp(0)
var _ Comparable = StringComp("")

type Comparable interface {
	CompareTo(Comparable) (int, error)
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

func ConvertArrayToComparable(array interface{}) []Comparable {
	return convertToComparable(reflect.ValueOf(array))
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
	out := make([]Comparable, len(s))
	for index, i := range s {
		result, err := handle(i)
		if err != nil {
			return nil, err
		}
		out[index] = result.(Comparable)
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

func ToComp(i interface{}) (Comparable, error) {
	if m, ok := i.(Comparable); ok {
		return m, nil
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Float32:
		return Comparable(FloatComp(i.(float32))), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return IntComp(int(reflect.ValueOf(i).Uint())), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		return IntComp(int(reflect.ValueOf(i).Int())), nil
	case reflect.Float64:
		return Comparable(Float64Comp(i.(float64))), nil
	case reflect.Uint64:
		return Comparable(Int64Comp(int64(i.(uint64)))), nil
	case reflect.Int64:
		return Comparable(Int64Comp(i.(int64))), nil
	case reflect.String:
		return Comparable(StringComp(i.(string))), nil
	default:
		return nil, ErrInvalid
	}
}

func ToCompArray(i interface{}) ([]Comparable, error) {
	if m, ok := i.([]Comparable); ok {
		return m, nil
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Slice:
		values := reflect.ValueOf(i)
		b := make([]Comparable, values.Len())
		for i := 0; i < values.Len(); i++ {
			c, err := ToComp(values.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			b[i] = c
		}
		return b, nil
	default:
		return nil, ErrInvalid
	}
}

func ToIntArray(s []Comparable) ([]int, error) {
	b := make([]int, len(s))
	for index, elem := range s {
		c, err := ToInt(elem)
		if err != nil {
			return nil, err
		}
		b[index] = c
	}
	return b, nil
}

func ToInt(i Comparable) (int, error) {
	if m, ok := i.(StringComp); ok {
		return strconv.Atoi(string(m))
	}
	if m, ok := i.(IntComp); ok {
		return int(m), nil
	}
	if m, ok := i.(Int64Comp); ok {
		return int(int64(m)), nil
	}
	if m, ok := i.(FloatComp); ok {
		return int(float32(m)), nil
	}
	if m, ok := i.(Float64Comp); ok {
		return int(float64(m)), nil
	}
	return 0, ErrInvalid
}

func ToInt64Array(s []Comparable) ([]int64, error) {
	b := make([]int64, len(s))
	for index, elem := range s {
		c, err := ToInt64(elem)
		if err != nil {
			return nil, err
		}
		b[index] = c
	}
	return b, nil
}

func ToInt64(i Comparable) (int64, error) {
	if m, ok := i.(StringComp); ok {
		return strconv.ParseInt(string(m), 10, 64)
	}
	if m, ok := i.(IntComp); ok {
		return int64(int(m)), nil
	}
	if m, ok := i.(Int64Comp); ok {
		return int64(m), nil
	}
	if m, ok := i.(FloatComp); ok {
		return int64(float32(m)), nil
	}
	if m, ok := i.(Float64Comp); ok {
		return int64(float64(m)), nil
	}
	return 0, ErrInvalid
}

func ToFloatArray(s []Comparable) ([]float32, error) {
	b := make([]float32, len(s))
	for index, elem := range s {
		c, err := ToFloat(elem)
		if err != nil {
			return nil, err
		}
		b[index] = c
	}
	return b, nil
}

func ToFloat(i Comparable) (float32, error) {
	if m, ok := i.(StringComp); ok {
		result, err := strconv.ParseFloat(string(m), 64)
		if err != nil {
			return 0, err
		}
		return float32(result), nil
	}
	if m, ok := i.(IntComp); ok {
		return float32(int(m)), nil
	}
	if m, ok := i.(Int64Comp); ok {
		return float32(int64(m)), nil
	}
	if m, ok := i.(FloatComp); ok {
		return float32(m), nil
	}
	if m, ok := i.(Float64Comp); ok {
		return float32(float64(m)), nil
	}
	return 0, ErrInvalid
}

func ToFloat64Array(s []Comparable) ([]float64, error) {
	b := make([]float64, len(s))
	for index, elem := range s {
		c, err := ToFloat64(elem)
		if err != nil {
			return nil, err
		}
		b[index] = c
	}
	return b, nil
}

func ToFloat64(i Comparable) (float64, error) {
	if m, ok := i.(StringComp); ok {
		return strconv.ParseFloat(string(m), 64)
	}
	if m, ok := i.(IntComp); ok {
		return float64(int(m)), nil
	}
	if m, ok := i.(Int64Comp); ok {
		return float64(int64(m)), nil
	}
	if m, ok := i.(FloatComp); ok {
		return float64(float32(m)), nil
	}
	if m, ok := i.(Float64Comp); ok {
		return float64(m), nil
	}
	return 0, ErrInvalid
}

func ToStringArray(s []Comparable) ([]string, error) {
	b := make([]string, len(s))
	for index, elem := range s {
		c, err := ToString(elem)
		if err != nil {
			return nil, err
		}
		b[index] = c
	}
	return b, nil
}

func ToElementArray(s []Comparable, elementType reflect.Type) (interface{}, error) {
	b := reflect.Zero(reflect.SliceOf(elementType))
	for _, elem := range s {
		if reflect.TypeOf(elem) == elementType {
			b = reflect.Append(b, reflect.ValueOf(elem))
			continue
		}

		if reflect.TypeOf(elem).Kind() == elementType.Kind() {
			if elementType.Kind() == reflect.Ptr {
				if reflect.TypeOf(elem).Elem().Kind() == elementType.Elem().Kind() {
					b = reflect.Append(b, reflect.ValueOf(elem).Convert(elementType))
					continue
				} else {
					return nil, ErrInvalid
				}
			} else {
				b = reflect.Append(b, reflect.ValueOf(elem).Convert(elementType))
				continue
			}
		}

		if elementType.Kind() == reflect.Ptr && reflect.TypeOf(elem).Kind() != reflect.Ptr {
			if reflect.TypeOf(elem).Kind() == elementType.Elem().Kind() {
				if reflect.ValueOf(elem).CanAddr() {
					b = reflect.Append(b, reflect.ValueOf(elem).Addr().Convert(elementType))
					continue
				}
			}
		}

		if elementType.Kind() != reflect.Ptr && reflect.TypeOf(elem).Kind() == reflect.Ptr {
			if reflect.TypeOf(elem).Elem().Kind() == elementType.Kind() {
				b = reflect.Append(b, reflect.ValueOf(elem).Elem().Convert(elementType))
				continue
			}
		}

		return nil, ErrInvalid
	}
	return b.Interface(), nil
}

func ToString(i Comparable) (string, error) {
	if m, ok := i.(StringComp); ok {
		return string(m), nil
	}
	if m, ok := i.(IntComp); ok {
		return strconv.Itoa(int(m)), nil
	}
	if m, ok := i.(Int64Comp); ok {
		return strconv.FormatInt(int64(m), 10), nil
	}
	if m, ok := i.(FloatComp); ok {
		return fmt.Sprint(float32(m)), nil
	}
	if m, ok := i.(Float64Comp); ok {
		return fmt.Sprint(float64(m)), nil
	}
	return "", ErrInvalid
}
