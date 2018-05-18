package secret

const (
	wink          = "wink"
	doubleBlink   = "double blink"
	closeYourEyes = "close your eyes"
	jump          = "jump"
)

// Handshake 1
func Handshake(input uint) []string {
	output := []string{}
	if input&1 == 1 {
		output = append(output, wink)
	}
	if input&2 == 2 {
		output = append(output, doubleBlink)
	}
	if input&4 == 4 {
		output = append(output, closeYourEyes)
	}
	if input&8 == 8 {
		output = append(output, jump)
	}
	if input&16 == 16 {
		output = reverse(output)
	}

	return output
}

func reverse(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverse(input[1:]), input[0])
}
