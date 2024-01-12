
package layers

import (
	"math"
	"strconv"
	"testing"
)

func TestSigmoidActivate(t *testing.T) {
	var sigmoid = func(x float64) float64 {
		return math.Exp(x) / (1.0 + math.Exp(x))
	}
	for k, v := range []float64{-1000, -100, -10, 0, 10, 100, 1000} {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			s := new(Sigmoid)
			want := sigmoid(v)
			got := s.Activate(v)
			if want-got > 1e-16 {
				t.Errorf("want %g, got %g", want, got)