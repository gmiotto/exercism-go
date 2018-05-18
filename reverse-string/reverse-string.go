package reverse

import (
	"strings"
)

// String function that reverses a string
func String(input string) string {
	output := []string{}
	for _, v := range input {
		output = append([]string{string(v)}, output...)
	}
	return strings.Join(output, "")
}
