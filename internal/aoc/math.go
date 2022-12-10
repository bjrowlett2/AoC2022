package aoc

import "math"

func Abs(value int) int {
	return int(math.Abs(float64(value)))
}

func Between(min, x, max int) bool {
	return (min <= x) && (x <= max)
}

func Sign(value int) int {
	return int(math.Copysign(1, float64(value)))
}
