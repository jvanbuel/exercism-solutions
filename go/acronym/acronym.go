// Package acronym is a solution to the acronym problem of the Exercism go track
package acronym

import (
	"log"
	"regexp"
	"strings"
)

// Abbreviate abbreviates a string s to its acronym
func Abbreviate(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9']+")
	if err != nil {
		log.Fatal(err)
	}

	s = reg.ReplaceAllString(s, " ")

	words := strings.Split(s, " ")
	var abbrev string
	for _, w := range words {
		if w != "" {
			abbrev = abbrev + strings.ToUpper(string(w[0]))
		}
	}
	return abbrev
}
