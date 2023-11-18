package layers

import (
	"fmt"
	"math"
	"strings"
)

type Sigmoid struct {
	layer
}

func (s *Sig