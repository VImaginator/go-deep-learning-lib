
package layers

import (
	"fmt"
	"testing"
)

func TestSoftmaxEstimate(t *testing.T) {
	var Max = func(a []float64) int {
		key := 0
		value := a[key]
		for k, v := range a {
			if v > value {
				key = k