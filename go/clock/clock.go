package clock

import "fmt"

// Clock is structure consisting of an hours an minutes field
type Clock struct {
	hour   int
	minute int
}

// New returns a new clock
func New(hour, minute int) Clock {
	if minute < 0 {
		carry := -(minute+1)/60 + 1
		return Clock{
			hour:   ((hour-carry)%24 + 24) % 24,
			minute: (minute%60 + 60) % 60}
	}
	return Clock{hour: ((hour+minute/60)%24 + 24) % 24,
		minute: (minute%60 + 60) % 60}
}

// String returns a string representation of a clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds minutes to a Clock
func (c Clock) Add(minutes int) Clock {
	carry := (c.minute + minutes) / 60
	return Clock{
		hour:   (c.hour + carry) % 24,
		minute: (c.minute + minutes) % 60,
	}
}

// Subtract substracts minutes from a Clock
func (c Clock) Subtract(minutes int) Clock {
	if c.minute < minutes {
		carry := -(c.minute-minutes+1)/60 + 1
		return Clock{
			hour:   ((c.hour-carry)%24 + 24) % 24,
			minute: ((c.minute-minutes)%60 + 60) % 60}
	}
	carry := (c.minute - minutes) / 60
	return Clock{
		hour:   ((c.hour-carry)%24 + 24) % 24,
		minute: ((c.minute-minutes)%60 + 60) % 60,
	}
}
