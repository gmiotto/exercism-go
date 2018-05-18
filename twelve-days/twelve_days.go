package twelve

import (
	"fmt"
	"strings"
)

var verses = []string{"a Partridge ",
	"two Turtle Doves, and ",
	"three French Hens, ",
	"four Calling Birds, ",
	"five Gold Rings, ",
	"six Geese-a-Laying, ",
	"seven Swans-a-Swimming, ",
	"eight Maids-a-Milking, ",
	"nine Ladies Dancing, ",
	"ten Lords-a-Leaping, ",
	"eleven Pipers Piping, ",
	"twelve Drummers Drumming, "}

var days = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

const (
	prefix = "On the %s day of Christmas my true love gave to me, "
	suffix = "in a Pear Tree."
)

// Song 1
func Song() string {
	var output strings.Builder
	for i := 0; i < len(verses); i++ {
		output.WriteString(Verse(i + 1))
		output.WriteString("\n")
	}
	return output.String()
}

// Verse 1
func Verse(verse int) string {
	var output strings.Builder
	output.WriteString(fmt.Sprintf(prefix, days[verse-1]))
	for i := verse - 1; i >= 0; i-- {
		output.WriteString(verses[i])
	}
	output.WriteString(suffix)
	return output.String()
}
