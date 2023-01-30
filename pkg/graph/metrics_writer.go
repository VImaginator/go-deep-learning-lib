package graph

type MetricsWriter interface {
	Write(Metrics)
}

type MetricsWriterFunc func(Metrics)

func (fn MetricsW