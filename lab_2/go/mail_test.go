package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)



func TestInterpolationNewton(t *testing.T) {

	assert.Equal(t, roundFloat(InterpolationNewton(x, mass, n), 3), 0.588)
	assert.Equal(t, InterpolationNewton(1, mass, n), float64(0))
	assert.Equal(t, InterpolationNewton(0, mass, n), float64(1))
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
