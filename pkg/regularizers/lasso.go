package regularizers

import "math"

type Lasso struct {
	Lambda  float64
	weights [][]float64
}

func (r Lasso) Regularize(gradi