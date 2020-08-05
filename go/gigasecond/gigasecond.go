// Package gigasecond is a solution to the gigasecond problem of the Exercism go track
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds 1 gigasecond to a time t
func AddGigasecond(t time.Time) time.Time {
	d, err := time.ParseDuration("1000000000s")
	if err != nil {
		return t
	}
	return t.Add(d)
}
