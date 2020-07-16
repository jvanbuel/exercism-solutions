// Package accumulate is a solution to the accumulate problem of the side exercises of the Exercism go track
package accumulate

// Accumulate applies a function f to each element of the collection of strings c
func Accumulate(c []string, f func(string) string) []string {
	r := make([]string, len(c))
	for i, s := range c {
		r[i] = f(s)
	}
	return r
}
