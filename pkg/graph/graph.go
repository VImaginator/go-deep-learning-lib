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
			f