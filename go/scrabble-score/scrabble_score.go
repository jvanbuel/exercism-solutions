// Package scrabble contains the solution to the scrabble problem of the Exercism go track
package scrabble

import "strings"

// Score calculates the scrabble score of a string s
func Score(s string) int {
	var score int

	for _, c := range strings.ToUpper(s) {
		score += scoreChar(string(c))
	}
	return score
}

func scoreChar(c string) int {

	switch c {
	case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
		return 1
	case "D", "G":
		return 2
	case "B", "C", "M", "P":
		return 3
	case "F", "H", "V", "W", "Y":
		return 4
	case "K":
		return 5
	case "J", "X":
		return 8
	case "Q", "Z":
		return 10
	}
	return 0
}
