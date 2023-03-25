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
			&