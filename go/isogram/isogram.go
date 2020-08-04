// Package isogram is a solution to the isogram problem of the Exercism go track
package isogram

import "unicode"

// IsIsogram return whether the input string s is an isogram or not
func IsIsogram(s string) bool {
	m := map[rune]bool{}
	for _, c := range s {
		c = unicode.ToLower(c)
		if c == '-' || c == ' ' {
			continue
		}
		if m[c] {
			return false
		}
		m[c] = true
	}
	return true
}
