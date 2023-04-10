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
			s = reflect.Typ