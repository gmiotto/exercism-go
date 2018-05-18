// Package twofer exposes the ShareWith function
package twofer

import (
	"fmt"
)

// ShareWith receives a string and return the answer.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
