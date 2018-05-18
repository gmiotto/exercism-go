package isogram

import (
	"strings"
)

// IsIsogram is a function to validate if a string do not have duplicates
func IsIsogram(input string) bool {
	if input == "" {
		return true
	}
	input = strings.Replace(input, "-", "", -1)
	input = strings.Replace(input, " ", "", -1)

	deduplicated := removeDuplicates([]rune(strings.ToLower(input)))
	return len(deduplicated) == len(input)
}

func removeDuplicates(input []rune) []string {
	inputSet := make(map[rune]bool)
	for _, v := range input {
		inputSet[v] = true
	}
	output := []string{}
	for key := range inputSet {
		output = append(output, string(key))
	}
	return output
}
