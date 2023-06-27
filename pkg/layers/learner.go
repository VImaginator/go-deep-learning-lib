package layers

type learner struct {
	weights        [][]float64
	localGradients [][]float64
}

func (l learner) Gradients() []