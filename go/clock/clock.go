package clock

import "fmt"

// Clock is structure consisting of an hours an minutes field
type Clock struct {
	minute int
}

// New returns a new clock
func New(hour, minute int) Clock {
	minutes := hour*60 + minute
	minutes %= 24 * 60
	if minutes < 0 {
		minutes += 24 * 60
	}
	return Clock{minutes}
}

// String returns a string representation of a clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minute/60, c.minute%60)
}

// Add adds minutes to a Clock
func (c Clock) Add(minutes int) Clock {
	return New(0, c.minute+minutes)
}

// Subtract substracts minutes from a Clock
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minute-minutes)
}
