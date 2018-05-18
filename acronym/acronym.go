// Package acronym contains Abbreviate function
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate returns the acronym of a String
func Abbreviate(s string) string {
	regp := regexp.MustCompile("\\b[A-Za-z]")
	abbrv := regp.FindAllString(s, -1)
	return strings.ToUpper(strings.Join(abbrv, ""))
}
