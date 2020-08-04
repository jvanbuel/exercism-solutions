package darts

import "math"

// Score calculates the score of darts with coordinates x and y
func Score(x float64, y float64) int {
	var score int
	if math.Pow(x, 2)+math.Pow(y, 2) <= 1 {
		return 10
	}
	if math.Pow(x, 2)+math.Pow(y, 2) <= 25 {
		return 5
	}
	if math.Pow(x, 2)+math.Pow(y, 2) <= 100 {
		return 1
	}
	return score
}
