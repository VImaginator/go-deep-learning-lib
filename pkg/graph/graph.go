package graph

import "math"

type Graph []Layer

func (g Graph) Apply(c Config) {
	for i := range g