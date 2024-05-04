package regularizers

import "math"

type Lasso struct {
	Lambda  float64
	weights [][]fl