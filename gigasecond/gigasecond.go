// Package gigasecond implements a function that adds
// 1 gigaseconds to a time
package gigasecond

// import path for the time package from the standard library
import "time"

const gigaSecond = 1e9

// AddGigasecond receives a Time instance and add 1 gigaseconds to it
// returns the added time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond * time.Second)
}
