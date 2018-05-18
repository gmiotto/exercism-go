package raindrops

import (
	"strconv"
	"strings"
)

// Convert a integer number to a raindrop output
func Convert(number int) string {
	var output strings.Builder

	if number%3 == 0 {
		output.WriteString("Pling")
	}
	if number%5 == 0 {
		output.WriteString("Plang")
	}
	if number%7 == 0 {
		output.WriteString("Plong")
	}
	if output.Len() == 0 {
		output.WriteString(strconv.Itoa(number))
	}
	return output.String()
}
