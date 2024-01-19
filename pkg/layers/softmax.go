package layers

import (
	"math"
)

type Softmax struct {
	layer
}

func (s *Softmax) Estimate(z []float64) []float64 {
