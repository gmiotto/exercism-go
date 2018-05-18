package flatten

// Flatten 1
func Flatten(input interface{}) []interface{} {
	output := []interface{}{}

	switch e := input.(type) {
	case []interface{}:
		for _, v := range e {
			output = append(output, Flatten(v)...)
		}
	case interface{}:
		output = append(output, e)
	}

	return output
}
