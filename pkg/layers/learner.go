package layers

type learner struct {
	weights        [][]float64
	localGradients [][]float64
}

fu