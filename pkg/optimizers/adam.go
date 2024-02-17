package optimizers

import "math"

type Adam struct {
	epsilon   float64
	gradients [][]float64
	momentum  M