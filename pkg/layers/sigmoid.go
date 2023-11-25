package layers

import (
	"fmt"
	"math"
	"strings"
)

type Sigmoid struct {
	layer
}

func (s *Sigmoid) Activate(z float64) float64 {
	return 1.0 / (1.0 + math.Exp(-z))
}

func (s *Sigmoid) Derive(a float6