package layers

import (
	"math"
)

type Softmax struct {
	layer
}

func (s *Softmax) Estimate(z []float64) []float64 {
	var max float64
	for _, v := range z {
		max = math.Max(max, v)
	}
	var sum float64
	for k, v := range z {
		z[k] = math.Exp(v - m