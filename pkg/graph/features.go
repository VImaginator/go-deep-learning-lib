
package graph

import (
	"math/rand"
)

type Features struct {
	ClassWeights        []float64
	DisableClassWeights bool
	DisableShuffle      bool