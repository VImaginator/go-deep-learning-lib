package layers

import (
	"math"
)

type Softmax struct {
	layer
}

func (s *Softmax) Estim