package optimizers

import "math"

type RMSprop struct {
	gradients [][]float64
	momentum  float64
}

func (m RMSprop) Apply(weights [][]float64) RMSprop {
	gradients := make([][]float64, len(weights))
	for i := range gradients {
		gradients[i] = make([]floa