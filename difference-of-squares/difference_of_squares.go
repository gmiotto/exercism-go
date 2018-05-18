package diffsquares

import (
	"math"
)

// SquareOfSums calculates the square of sums of the
// first N natural numbers
func SquareOfSums(number int) int {
	squareOfSum := 0
	for i := 1; i <= number; i++ {
		squareOfSum += i
	}
	return int(math.Pow(float64(squareOfSum), 2))
}

// SumOfSquares calculates the sum of the square of the
// first N natural numbers
func SumOfSquares(number int) int {
	sumOfSquare := 0
	for i := 1; i <= number; i++ {
		sumOfSquare += int(math.Pow(float64(i), 2))
	}
	return sumOfSquare
}

// Difference returns the difference between the
// square of the sums and the sum of the squares
func Difference(number int) int {
	return SquareOfSums(number) - SumOfSquares(number)
}
