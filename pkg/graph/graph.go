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
	gradients := make([][][]float64, len(g))
	for i := range gradients {
		if l, ok := g[i].(Minimizeable); ok {
			gradients[i] = l.Gradients()
		}
	}
	return gradients
}

func (g Graph) Loss(x, y []float64) []float64 {
	a := g.Estimate(x)
	loss := make([]float64, len(y))
	for i := range loss {
		loss[i] = a[i] - y[i]
	}
	return loss
}

func (g Graph) Minimize(gradients []float64) []float64 {
	j := len(g) - 1
	for i := range g {
		gradients = g[j-i].Mi