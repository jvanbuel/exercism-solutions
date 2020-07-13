// Package twofer is the solution to the two-fer problem of the Exercism Go track.
package twofer

import "fmt"

// ShareWith takes a string 'name' as input and outputs a formatted string containing that name. If no name is given, a default string is returned.
func ShareWith(name string) string {

	if name == "" {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
