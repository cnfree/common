package common

import (
	"sync"
)

type numberUtil struct {
	mutex sync.Mutex
}

var Number = numberUtil{}

func (this numberUtil) ReduceInt(s []int, handle func(int, int) int) int {
	out := 0
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

func (this numberUtil) ReduceInt64(s []int64, handle func(int64, int64) int64) int64 {
	var out int64 = 0
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

func (this numberUtil) ReduceFloat(s []float32, handle func(float32, float32) float32) float32 {
	var out float32 = 0
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

func (this numberUtil) ReduceFloat64(s []float64, handle func(float64, float64) float64) float64 {
	var out float64 = 0
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

func (this numberUtil) minInt(s1 int, s2 int) int {
	if s1 < s2 {
		return s1
	}
	return s2
}

func (this numberUtil) maxInt(s1 int, s2 int) int {
	if s1 > s2 {
		return s1
	}
	return s2
}

func (this numberUtil) sumInt(s1 int, s2 int) int {
	return s1 + s2
}

func (this numberUtil) MinInt(s ...int) int {
	return this.ReduceInt(s, this.minInt)
}

func (this numberUtil) MaxInt(s ...int) int {
	return this.ReduceInt(s, this.maxInt)
}

func (this numberUtil) SumInt(s ...int) int {
	return this.ReduceInt(s, this.sumInt)
}

func (this numberUtil) minInt64(s1 int64, s2 int64) int64 {
	if s1 < s2 {
		return s1
	}
	return s2
}

func (this numberUtil) maxInt64(s1 int64, s2 int64) int64 {
	if s1 > s2 {
		return s1
	}
	return s2
}

func (this numberUtil) sumInt64(s1 int64, s2 int64) int64 {
	return s1 + s2
}

func (this numberUtil) MinInt64(s ...int64) int64 {
	return this.ReduceInt64(s, this.minInt64)
}

func (this numberUtil) MaxInt64(s ...int64) int64 {
	return this.ReduceInt64(s, this.maxInt64)
}

func (this numberUtil) SumInt64(s ...int64) int64 {
	return this.ReduceInt64(s, this.sumInt64)
}

func (this numberUtil) minFloat(s1 float32, s2 float32) float32 {
	if s1 < s2 {
		return s1
	}
	return s2
}

func (this numberUtil) maxFloat(s1 float32, s2 float32) float32 {
	if s1 > s2 {
		return s1
	}
	return s2
}

func (this numberUtil) sumFloat(s1 float32, s2 float32) float32 {
	return s1 + s2
}

func (this numberUtil) MinFloat(s ...float32) float32 {
	return this.ReduceFloat(s, this.minFloat)
}

func (this numberUtil) MaxFloat(s ...float32) float32 {
	return this.ReduceFloat(s, this.maxFloat)
}

func (this numberUtil) SumFloat(s ...float32) float32 {
	return this.ReduceFloat(s, this.sumFloat)
}

func (this numberUtil) minFloat64(s1 float64, s2 float64) float64 {
	if s1 < s2 {
		return s1
	}
	return s2
}

func (this numberUtil) maxFloat64(s1 float64, s2 float64) float64 {
	if s1 > s2 {
		return s1
	}
	return s2
}

func (this numberUtil) sumFloat64(s1 float64, s2 float64) float64 {
	return s1 + s2
}

func (this numberUtil) MinFloat64(s ...float64) float64 {
	return this.ReduceFloat64(s, this.minFloat64)
}

func (this numberUtil) MaxFloat64(s ...float64) float64 {
	return this.ReduceFloat64(s, this.maxFloat64)
}

func (this numberUtil) SumFloat64(s ...float64) float64 {
	return this.ReduceFloat64(s, this.sumFloat64)
}

func (this numberUtil) MapInt(s []int, handle func(int) int) []int {
	out := []int{}
	for _, i := range s {
		out = append(out, handle(i))
	}
	return out
}

func (this numberUtil) MapInt64(s []int64, handle func(int64) int64) []int64 {
	out := []int64{}
	for _, i := range s {
		out = append(out, handle(i))
	}
	return out
}

func (this numberUtil) MapFloat(s []float32, handle func(float32) float32) []float32 {
	out := []float32{}
	for _, i := range s {
		out = append(out, handle(i))
	}
	return out
}

func (this numberUtil) MapFloat64(s []float64, handle func(float64) float64) []float64 {
	out := []float64{}
	for _, i := range s {
		out = append(out, handle(i))
	}
	return out
}

func (this numberUtil) FilterInt(s []int, handle func(int) bool) []int {
	out := []int{}
	for _, i := range s {
		if !handle(i) {
			continue
		}
		out = append(out, i)
	}
	return out
}

func (this numberUtil) FilterInt64(s []int64, handle func(int64) bool) []int64 {
	out := []int64{}
	for _, i := range s {
		if !handle(i) {
			continue
		}
		out = append(out, i)
	}
	return out
}

func (this numberUtil) FilterFloat(s []float32, handle func(float32) bool) []float32 {
	out := []float32{}
	for _, i := range s {
		if !handle(i) {
			continue
		}
		out = append(out, i)
	}
	return out
}

func (this numberUtil) FilterFloat64(s []float64, handle func(float64) bool) []float64 {
	out := []float64{}
	for _, i := range s {
		if !handle(i) {
			continue
		}
		out = append(out, i)
	}
	return out
}
