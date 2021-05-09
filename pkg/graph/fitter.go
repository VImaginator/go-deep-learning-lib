package graph

type Fitter struct {
	Epochs               uint64
	Training, Validation Features
}

func (f Fitter) Pr