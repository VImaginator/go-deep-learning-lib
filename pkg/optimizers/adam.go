package optimizers

import "math"

type Adam struct {
	epsilon   float64
	gradients [][]float64
	momentum  Momentum
	rmsprop   RMSprop
}

func (m Adam) Apply(weights [][]float64) Adam {
	gradients := make([][]float64, len(weights))
	for i := range gradients {
		gradients[i] = make([]float64, len(weights[i]))
	}
