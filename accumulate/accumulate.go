package accumulate

// Accumulate 1
func Accumulate(input []string, converter func(string) string) []string {
	output := []string{}
	for _, str := range input {
		output = append(output, converter(str))
	}
	return output
}
