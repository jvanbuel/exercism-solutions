// Package hamming contains the solution to the Hamming problem of the Exercism Go track
package hamming

import "fmt"

// Distance calculates the Hamming distance between two strings a and b
func Distance(a, b string) (int, error) {

	var aRune = []rune(a)
	var bRune = []rune(b)

	if len(aRune) != len(bRune) {
		return 0, fmt.Errorf("string a has length %d and string b has length %d", len(a), len(b))
	}

	var h int
	for i, c := range aRune {
		if bRune[i] != c {
			h++
		}
	}

	return h, nil
}
