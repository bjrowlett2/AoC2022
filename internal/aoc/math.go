package aoc

import "math"

func Abs(value int) int {
	return int(math.Abs(float64(value)))
}

func Sign(value int) int {
	return int(math.Copysign(1, float64(value)))
}
