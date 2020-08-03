// Package acronym is a solution to the acronym problem of the Exercism go track
package acronym

import (
	"strings"
)

// Abbreviate abbreviates a string s to its acronym
func Abbreviate(s string) string {

	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")

	words := strings.Split(s, " ")
	var abbrev string
	for _, w := range words {
		if w != "" {
			abbrev = abbrev + strings.ToUpper(string(w[0]))
		}
	}
	return abbrev
}
