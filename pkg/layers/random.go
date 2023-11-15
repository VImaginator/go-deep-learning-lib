package layers

import (
	"math/rand"
)

func Random(size int) []float64 {
	w := make([]float64, size)
	r := rand