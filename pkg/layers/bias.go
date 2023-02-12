
package layers

type Bias struct {
	layer
	learner
}

func (b *Bias) Estimate(input []float64) []float64 {
	for i := range input {
		b.output[i] = input[i] + b.weights[0][i]
	}
	return b.output
}

func (b *Bias) Minimize(gradients []float64) []float64 {
	for k, v := range gradients {
		b.localGradients[0][k] = v
	}
	return gradients
}

func (b *Bias) SetShape(shape []uint64) {
	b.layer.SetShape(shape)