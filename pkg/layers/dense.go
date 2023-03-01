package layers

type Dense struct {
	bias    Bias
	dense   UnbiasedDense
	Neurons uint64
}

func (l *Dense) Estimate(input []float64) []float64 {
	input = l.dense.Estimate(input)
	return l.bias.Estimate(input)
}

func (l *Dense) Gradients() [][]float64 {
	return append(l.bias.localGradients, l.dense.localGradients...)
}

func (l *Dense) Minimize(gradients []float64) []float64 {
	gradients = l.bias.Minimize(gradients)
	return l.dense.Minimize(gradien