package optimizers

import "math"

type Adam struct {
	epsilon   float64
	gradients [][]float64
	momentum  Momentum
	rmsprop   RMSprop
}

func (m Adam) Apply(weights [][]float64) Adam {
	gradien