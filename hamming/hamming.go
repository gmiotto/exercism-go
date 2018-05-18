// Package hamming exports the hamming Distance function
package hamming

import "errors"

// Distance function verify the hamming distance between two DNA strands.
func Distance(a, b string) (int, error) {
	runeA := []rune(a)
	runeB := []rune(b)
	distance := 0

	if !isValid(runeA, runeB) {
		return -1, errors.New("sequences must have the same size")
	}

	for i, v := range runeA {
		if v != runeB[i] {
			distance++
		}
	}
	return distance, nil
}

func isValid(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	return true
}
