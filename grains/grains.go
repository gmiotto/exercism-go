package grains

import (
	"errors"
	"math"
	"math/big"
)

const (
	a1           = 1
	ratio        = 2
	totalSquares = 64
)

// Square returns the number of grains on a given square
// implemented using the geometric progression formula
func Square(input int) (uint64, error) {
	if input <= 0 || input > totalSquares {
		return 0, errors.New("square number must be positive")
	}
	aInput := a1 * int(math.Pow(ratio, float64(input-1)))
	return uint64(aInput), nil
}

// Total returns the total number of wheat on the whole board
// implemented using the finite sum of a geometric progression
func Total() uint64 {
	var x, y = big.NewInt(ratio), big.NewInt(totalSquares)
	pow := x.Exp(x, y, nil)
	return pow.Sub(pow, big.NewInt(1)).Uint64()
}
