package common

import (
	"strconv"
	"sync"
	"math/big"
	"github.com/cnfree/common/utils"
)

type stringUtil struct {
	sync.Mutex
}

var String = stringUtil{}

func (_ stringUtil) Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r) - 1; i < len(r)/2; i, j = i + 1, j - 1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func (this stringUtil) ToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// Float32 string to float32
func (this stringUtil) ToFloat32(s string) (float32, error) {
	v, err := strconv.ParseFloat(s, 32)
	return float32(v), err
}

// Float64 string to float64
func (this stringUtil) ToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// Int string to int
func (this stringUtil) ToInt(s string) (int, error) {
	v, err := strconv.ParseInt(s, 10, 32)
	return int(v), err
}

// Int8 string to int8
func (this stringUtil) ToInt8(s string) (int8, error) {
	v, err := strconv.ParseInt(s, 10, 8)
	return int8(v), err
}

// Int16 string to int16
func (this stringUtil) ToInt16(s string) (int16, error) {
	v, err := strconv.ParseInt(s, 10, 16)
	return int16(v), err
}

// Int32 string to int32
func (this stringUtil) ToInt32(s string) (int32, error) {
	v, err := strconv.ParseInt(s, 10, 32)
	return int32(v), err
}

// Int64 string to int64
func (this stringUtil) ToInt64(s string) (int64, error) {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		i := new(big.Int)
		ni, ok := i.SetString(s, 10) // octal
		if !ok {
			return int64(v), err
		}
		return ni.Int64(), nil
	}
	return int64(v), err
}

// Uint string to uint
func (this stringUtil) ToUint(s string) (uint, error) {
	v, err := strconv.ParseUint(s, 10, 32)
	return uint(v), err
}

// Uint8 string to uint8
func (this stringUtil) ToUint8(s string) (uint8, error) {
	v, err := strconv.ParseUint(s, 10, 8)
	return uint8(v), err
}

// Uint16 string to uint16
func (this stringUtil) ToUint16(s string) (uint16, error) {
	v, err := strconv.ParseUint(s, 10, 16)
	return uint16(v), err
}

// Uint32 string to uint31
func (this stringUtil) ToUint32(s string) (uint32, error) {
	v, err := strconv.ParseUint(s, 10, 32)
	return uint32(v), err
}

// Uint64 string to uint64
func (this stringUtil) ToUint64(s string) (uint64, error) {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		i := new(big.Int)
		ni, ok := i.SetString(s, 10)
		if !ok {
			return uint64(v), err
		}
		return ni.Uint64(), nil
	}
	return uint64(v), err
}

func (this stringUtil) min(s1 string, s2 string) string {
	if s1 < s2 {
		return s1
	}
	return s2
}

func (this stringUtil) max(s1 string, s2 string) string {
	if s1 > s2 {
		return s1
	}
	return s2
}

func (this stringUtil) Min(s ... string) string {
	return this.Reduce(s, this.min)
}

func (this stringUtil) Max(s ... string) string {
	return this.Reduce(s, this.max)
}

func (this stringUtil) Map(s []string, handle func(string) string) []string {
	out := make([]string, len(s))
	for index, i := range s {
		out[index] = handle(i)
	}
	return out
}

func (this stringUtil) MapToInt(s []string, handle func(string) int) []int {
	out := make([]int, len(s))
	for index, i := range s {
		out[index] = handle(i)
	}
	return out
}

func (this stringUtil) MapToInt64(s []string, handle func(string) int64) []int64 {
	out := make([]int64, len(s))
	for index, i := range s {
		out[index] = handle(i)
	}
	return out
}

func (this stringUtil) MapToFloat(s []string, handle func(string) float32) []float32 {
	out := make([]float32, len(s))
	for index, i := range s {
		out[index] = handle(i)
	}
	return out
}

func (this stringUtil) MapToFloat64(s []string, handle func(string) float64) []float64 {
	out := make([]float64, len(s))
	for index, i := range s {
		out[index] = handle(i)
	}
	return out
}

func (this stringUtil) Reduce(s []string, handle func(string, string) string) string {
	out := ""
	if len(s) > 0 {
		out = s[0]
	}
	for index, i := range s {
		if index == 0 {
			continue
		}
		out = handle(out, i)
	}
	return out
}

func (this stringUtil) Filter(s []string, handle func(string) bool) []string {
	out := []string{}
	for _, i := range s {
		if !handle(i) {
			continue
		}
		out = append(out, i)
	}
	return out
}

// StringValue returns the string representation of a value.
func (str stringUtil) ToString(i interface{}) string {
	return utils.ToString(i)
}
