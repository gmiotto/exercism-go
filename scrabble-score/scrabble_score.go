package scrabble

import (
	"strings"
)

// Score function to get string score
func Score(input string) int {
	count := 0
	for _, letter := range []rune(strings.ToUpper(input)) {
		count += getScorePerRune(letter)
	}
	return count
}

func getScorePerRune(letter rune) int {
	switch letter {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	}
	return 0
}
