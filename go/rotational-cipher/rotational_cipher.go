package rotationalcipher

import (
	"unicode"
)

// RotationalCipher encodes a string s using the Caeser cipher with a given shift
func RotationalCipher(s string, shift int) string {
	var encoded string

	for _, c := range s {
		if unicode.IsLetter(c) {
			start := 'a'
			if unicode.IsUpper(c) {
				start = 'A'
			}
			encoded += string((c-start+rune(shift))%26 + start)
			continue
		}
		encoded += string(c)

	}
	return encoded
}
