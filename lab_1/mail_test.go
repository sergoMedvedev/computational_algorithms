package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

var mass [][]float64 = [][]float64{
	{0.5, 0.707},
	{0.25, 0.924},
	{0, 1},
	{0.25, 0.924},
	{0.5, 0.707},
	{0.75, 0.383},
	{1, 0},
}

var x float64 = 0.6
var n int = 4

func TestInterpolationNewton(t *testing.T) {

	assert.Equal(t, roundFloat(InterpolationNewton(x, mass, n), 3), 0.588)
	assert.Equal(t, InterpolationNewton(1, mass, n), float64(0))
	assert.Equal(t, InterpolationNewton(0, mass, n), float64(1))
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
