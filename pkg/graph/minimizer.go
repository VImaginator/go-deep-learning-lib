
package graph

type Minimizer struct {
	Layer
	cursor       int
	batch        [][][]float64
	gradients    [][]float64
	learningRate LearningRate
	optimizer    Optimizer
	regularizer  Regularizer
	weights      [][]float64
}

func (m *Minimizer) Minimize(gradients []float64) []float64 {
	gradients = m.Layer.Minimize(gradients)
	for i := range m.gradients {
		copy(m.batch[m.cursor][i], m.gradients[i])
	}
	m.cursor++
	if m.cursor < len(m.batch) {
		return gradients
	}
	m.cursor = 0
	m.average()
	m.regularize()
	m.optimize()
	lr := m.learningRate.Rate()