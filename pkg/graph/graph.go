package graph

import "math"

type Graph []Layer

func (g Graph) Apply(c Config) {
	for i := range g {
		if layer, ok := g[i].(*Minimizer); ok {
			layer.learningRate = c.LearningRate
		