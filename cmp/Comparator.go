package cmp

import (
	"bytes"
	"reflect"
	"github.com/cnfree/common/errors"
)

// Comparator compares a and b and returns:
//     0 if they are equal
//     < 0 if a < b
//     > 0 if a > b

type Comparator func(a, b interface{}) int

func Compare(a interface{}, b interface{}) int {
	value1 := reflect.ValueOf(a)
	value2 := reflect.ValueOf(b)

	if value1.Kind() != value2.Kind() {
		panic(errors.Errorf("Compare needs two same type, but %s != %s", value1.Kind(), value2.Kind()))
	}

	switch a.(type) {
	case int, int8, int16, int32, int64:
		a1 := value1.Int()
		a2 := value2.Int()
		if a1 > a2 {
			return 1
		} else if a1 == a2 {
			return 0
		} else {
			return -1
		}
	case uint, uint8, uint16, uint32, uint64:
		a1 := value1.Uint()
		a2 := value2.Uint()
		if a1 > a2 {
			return 1
		} else if a1 == a2 {
			return 0
		} else {
			return -1
		}
	case float32, float64:
		a1 := value1.Float()
		a2 := value2.Float()
		if a1 > a2 {
			return 1
		} else if a1 == a2 {
			return 0
		} else {
			return -1
		}
	case string:
		a1 := value1.String()
		a2 := value2.String()
		if a1 > a2 {
			return 1
		} else if a1 == a2 {
			return 0
		} else {
			return -1
		}
	case []byte:
		a1 := value1.Bytes()
		a2 := value2.Bytes()
		return bytes.Compare(a1, a2)
	default:
		panic(errors.Errorf("type %T is not supported now", a))
	}
}

func Less(v1 interface{}, v2 interface{}) bool {
	n := Compare(v1, v2)
	return n < 0
}

func LessEqual(v1 interface{}, v2 interface{}) bool{
	n := Compare(v1, v2)
	return n <= 0
}

func Greater(v1 interface{}, v2 interface{}) bool {
	n := Compare(v1, v2)
	return n > 0
}

func GreaterEqual(v1 interface{}, v2 interface{}) bool {
	n := Compare(v1, v2)
	return n >= 0
}
