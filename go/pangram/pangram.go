// Package pangram is a solution to the pangram problem of the Exercism go track
package pangram

import "unicode"

// IsPangram determines whether a given string s is a pangram or not
func IsPangram(s string) bool {
	if s == "" {
		return false
	}

	m := map[rune]bool{
		'a': true,
		'b': true,
		'c': true,
		'd': true,
		'e': true,
		'f': true,
		'g': true,
		'h': true,
		'i': true,
		'j': true,
		'k': true,
		'l': true,
		'm': true,
		'n': true,
		'o': true,
		'p': true,
		'q': true,
		'r': true,
		's': true,
		't': true,
		'u': true,
		'v': true,
		'w': true,
		'x': true,
		'y': true,
		'z': true,
	}
	for _, c := range s {
		c = unicode.ToLower(c)
		if _, ok := m[c]; ok {
			delete(m, c)
		}
	}
	if len(m) == 0 {
		return true
	}
	return false
}
