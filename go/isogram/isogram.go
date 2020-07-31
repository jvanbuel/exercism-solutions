// Package isogram is a solution to the isogram problem of the Exercism go track
package isogram

import (
	"strings"
)

// IsIsogram return whether the input string s is an isogram or not
func IsIsogram(s string) bool {
	m := map[rune]bool{}
	for _, c := range strings.ToLower(s) {
		if c == '-' || c == ' ' {
			continue
		}
		if m[c] == true {
			return false
		}
		m[c] = true
	}
	return true
}
