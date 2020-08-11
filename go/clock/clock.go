package clock

import "fmt"

// Clock is structure consisting of an hours an minutes field
type Clock struct {
	hour   int
	minute int
}

// New returns a new clock
func New(hour, minute int) Clock {
	return Clock{
		hour:   posRemainder(hour+carry(minute), 24),
		minute: posRemainder(minute, 60),
	}
}

// String returns a string representation of a clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds minutes to a Clock
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

// Subtract substracts minutes from a Clock
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute-minutes)
}

// Carry calculates the number of hours that need to be carried to the hour field of a clock
func carry(minutes int) int {
	if minutes < 0 {
		return (minutes+1)/60 - 1
	}
	return minutes / 60
}

// posRemainder calculates the positive remainder of the division of i by m
func posRemainder(i int, m int) int {
	return (i%m + m) % m
}
