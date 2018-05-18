package luhn

import (
	"strconv"
	"strings"
)

// Valid function implements luhn algorith
func Valid(input string) bool {
	input = strings.Replace(input, " ", "", -1)
	if len(input) < 2 {
		return false
	}
	sum := 0
	for i := 0; i < len(input); i++ {
		j := len(input) - 1 - i
		tempDigit, err := strconv.Atoi(string(input[j]))
		if err != nil {
			return false
		}
		if i%2 != 0 {
			tempDigit = tempDigit * 2
			if tempDigit > 9 {
				tempDigit -= 9
			}
		}
		sum += tempDigit
	}

	return sum%10 == 0
}
