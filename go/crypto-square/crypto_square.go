package cryptosquare

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

// Encode encodes a string s according to the crypto square algorithm
func Encode(s string) string {
	if s == "" {
		return ""
	}
	var encoded string
	s = strings.ToLower(regexp.MustCompile("[^a-zA-Z0-9]").ReplaceAllString(s, ""))
	c, r := GetCR(s)

	s = fmt.Sprintf("%-*v", c*r, s) // Pad string to right with spaces until it has length c*r

	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {

			encoded += string(s[i+j*c])
		}
		encoded += " "
	}
	return encoded[:len(encoded)-1]
}

// GetCR returns the rows and columns for the "square"
func GetCR(s string) (int, int) {
	sq := int(math.Sqrt(float64(len(s))))

	if sq*sq == len(s) {
		return sq, sq
	}
	c := sq + 1
	r := sq

	if c*r < len(s) {
		r++
	}
	return c, r
}
