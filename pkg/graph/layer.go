package graph

type Layer interface {
	Estimate([]float64) []float64
	Minimize([]float64) []float64
	Se