
package layers

import (
	"fmt"
	"math"
	"strings"
)

type UnbiasedDense struct {
	layer
	learner
	gradients []float64
	input     []float64
	Neurons   uint64
}

func (d *UnbiasedDense) Estimate(input []float64) []float64 {
	d.input = input