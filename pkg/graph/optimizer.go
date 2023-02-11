package graph

type Optimizer interface {
	Optimize(gradients [][]float64) [][]f