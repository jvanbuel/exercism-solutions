//Package etl is a solution to the etl problem of the side exercises in the Exercism go track
package etl

import "strings"

// Transform converts a map i of scrabble scores to their corresponding letters to a map m of letters and their corresponding scrabble scores
func Transform(i map[int][]string) map[string]int {
	m := make(map[string]int)
	for score, c := range i {
		for _, s := range c {
			m[strings.ToLower(s)] = score
		}
	}
	return m
}
