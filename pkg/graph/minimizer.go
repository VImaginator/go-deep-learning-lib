
package graph

type Minimizer struct {
	Layer
	cursor       int
	batch        [][][]float64
	gradients    [][]float64