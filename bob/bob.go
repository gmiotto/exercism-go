// Package bob contains a function Hey that receives a string and answer it
package bob

import (
	"regexp"
	"strings"
)

// Hey function that receives a string and answer accordingly to Bob rules
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	reg, _ := regexp.Compile("[0-9]*[A-Za-z]+")

	if strings.HasSuffix(remark, "?") {
		if reg.MatchString(remark) && strings.ToUpper(remark) == remark {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}

	if reg.MatchString(remark) && strings.ToUpper(remark) == remark {
		return "Whoa, chill out!"
	}

	if remark == "" {
		return "Fine. Be that way!"
	}

	return "Whatever."

}
