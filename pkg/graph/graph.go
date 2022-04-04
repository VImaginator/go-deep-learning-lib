package graph

import "math"

type Graph []Layer

func (g Graph) Apply(c Config) {
	for i := range g {
		if layer, ok := g[i].(*Minimizer); ok {
			layer.learningRate = c.LearningRate
			layer.optimizer = c.Optimizer
			layer.regularizer = c.Regularizer
			layer.batch = make([][][]float64, c.BatchSize)
			for j := range layer.batch {
				layer.batch[j] = make([][]float64, len(layer.gradients))
				for k := range layer.gradients {
					layer.batch[j][k] = make([]float64, len(layer.gradients[k]))
				}
			}
		}
	}
}

func (g Graph) Estimate(x []float64) []float64 {
	for i := range g {
		x = g[i].Estimate(x)
	}
	return x
}

func (g Graph) Gradients() [][][]float64 {
	gradients := make([][][]float64, len