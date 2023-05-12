package layers

import (
	"math"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/therfoo/therfoo/pkg/graph"
)

func TestGradient(t *testing.T) {
	var linearGraph = func() graph.Graph {
		return graph.New(
			&Input{2},
			&Dense{Neurons: 2},
			&Sigmoid{},
			&Dense{Neurons: 1},
			&Sigmoid{},
		)
	}

	var polynomialGraph = func() graph.Graph {
		return graph.New(
			&Input{2},
			&Dense{Neurons: 1},
			&Polynomial{Degree: 2},
			&Sigmoid{},
		)
	}

	var name = func(l graph.Layer) string {
		type Stringer interface{ String() string }
		var s string
		if stringer, ok := l.(Stringer); ok {
			s = strings.Split(stringer.String(), "\n")[0]
		} else {
			s = reflect.TypeOf(l).String()
		}
		return s
	}

	var writer graph.MetricsWriterFunc = func(metrics graph.Metrics) {
		t.Log(metrics.String())
	}

	for k, v := range []struct {
		Epochs uint64
		Data   graph.Features
		Graph  graph.Graph
	}{
		{0, AND, linearGraph()},
		{1, NAND, linearGraph()},
		{2, OR, linearGraph()},
		{3, XOR, linearGraph()},
		{0, AND, polynomialGraph()},
		{1, NAND, polynomialGraph()},
		{2, OR, polynomialGraph()},
		{3, XOR, polynomialGraph()},
	} {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			v.Data.DisableClassWeights = true
			v.Data.DisableShuffle = true

			v.Graph.Apply(graph.Config{}.Validate())

			fitter := graph.Fitter{Training: v.Data}
			fitter.Fi