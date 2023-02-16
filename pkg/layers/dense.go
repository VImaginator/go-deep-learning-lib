package layers

type Dense struct {
	bias    Bias
	dense   UnbiasedDense
	Neurons uint64
}

func (l *Dense) Estimate(in