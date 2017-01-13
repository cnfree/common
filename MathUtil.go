package common

import (
	"math"
	"sync"
)

type mathUtil struct {
	mutex sync.Mutex
}

var Math = mathUtil{}

func (this mathUtil) Round(x float64, precision int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(precision))
	intermediate := x * pow

	if intermediate < 0.0 {
		intermediate -= 0.5
	} else {
		intermediate += 0.5
	}
	rounder = float64(int64(intermediate))

	return rounder / float64(pow)
}
