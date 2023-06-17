
package layers

type layer struct {
	inputShape  Shape
	outputShape Shape
	output      []float64
}

func (l *layer) Minimize(loss []float64) []float64 {