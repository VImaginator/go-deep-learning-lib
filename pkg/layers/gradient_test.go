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
	var l