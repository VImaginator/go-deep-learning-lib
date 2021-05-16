package graph

type Fitter struct {
	Epochs               uint64
	Training, Validation Features
}

func (f Fitter) Prepare() Fitter {
	f.Training.Prepare()
	f.Validation.Prepare()
	return f
}

func (f Fitter) Fit(g Graph, w ...MetricsWriter) {
	for i