// Package hamming contains the solution to the Hamming problem of the Exercism Go track
package hamming

import "fmt"

// Distance calculates the Hamming distance between two strings a and b
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("String a has length %d and string b has length %d", len(a), len(b))
	}

	var h int = 0
	for i, c := range a {
		if b[i] != byte(c) {
			h++
		}
	}
	return h, nil
}
