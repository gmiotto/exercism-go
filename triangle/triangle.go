package triangle

import (
	"math"
)

// Kind 1
type Kind string

// Types of triangle
const (
	NaT = "not a triangle"
	Equ = "equilateral"
	Iso = "isosceles"
	Sca = "scalene"
)

// KindFromSides 1
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	}
	if a == b && a == c && b == c {
		return Equ
	}
	if a == b || b == c || c == a || c == b {
		return Iso
	}
	return Sca
}

func isTriangle(a, b, c float64) bool {
	if !isSideValid(a) || !isSideValid(b) || !isSideValid(c) ||
		a+b < c || a+c < b || c+b < a {
		return false
	}

	return true
}

func isSideValid(side float64) bool {
	return !math.IsInf(side, 0) && !math.IsNaN(side) && side > 0
}
