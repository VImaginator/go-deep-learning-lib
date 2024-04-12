package optimizers

import "math"

type RMSprop struct {
	gradients [][]float64
	momentum  float6