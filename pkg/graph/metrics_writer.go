package graph

type MetricsWriter interface {
	Write(Metrics)
}

type M