package clock

import "fmt"

// Clock type to be exported
type Clock struct {
	hour   int
	minute int
}

const (
	secondsInHours   = 3600
	secondsInMinutes = 60
)

// New 1
func New(hour, minute int) Clock {
	seconds := secondsInHours*hour + secondsInMinutes*minute
	c := Clock{0, 0}
	return c.addSecondsToClock(seconds)
}

// String 1
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add 1
func (c Clock) Add(minutes int) Clock {
	return c.addSecondsToClock(minutes * secondsInMinutes)
}

// Subtract 1
func (c Clock) Subtract(minutes int) Clock {
	return c.addSecondsToClock(minutes * secondsInMinutes * -1)
}

func (c Clock) addSecondsToClock(seconds int) Clock {
	_seconds := secondsInHours*c.hour + secondsInMinutes*c.minute + seconds
	hours := (_seconds / secondsInHours) % 24
	if hours < 0 {
		hours += 24
	}
	minutes := (_seconds % secondsInHours) / secondsInMinutes
	if minutes < 0 {
		hours = (hours + 23) % 24
		minutes = minutes + 60
	}
	c.hour = hours
	c.minute = minutes
	return c
}
