
package graph

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestFeaturesClassWeights(t *testing.T) {
	for k, v := range []struct {
		samples  int
		classes  int
		features int
		minority int
		weights  []float64
	}{
		{43400, 5, 21, 783, []float64{0.8159428463996992, 0.8102305610006534, 0.8234512854567878, 0.8093240093240093, 11.085568326947637}},
	} {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			features := Features{
				X: make([][]float64, v.samples),
				Y: make([][]float64, v.samples),
			}
			rand.Seed(80)
			n := v.classes - 1
			for i := 0; i < v.samples; i++ {
				features.X[i] = make([]float64, v.features)
				features.Y[i] = make([]float64, v.classes)
				if i < v.minority {
					features.Y[i][n] = 1