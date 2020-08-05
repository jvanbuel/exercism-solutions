// Package diffsquares is a solution to the difference of squares problem of the Exercism go track
package diffsquares

// SquareOfSum calculates the square of the sum of the first i integers
func SquareOfSum(i int) int {
	sum := i * (i + 1) / 2
	return sum * sum
}

// SumOfSquares calculates the sum of the squares of the first i integers
func SumOfSquares(i int) int {
	return i * (i + 1) * (2*i + 1) / 6
}

// Difference calculates the difference between the square of sum and the sum of squares
func Difference(i int) int {
	return SquareOfSum(i) - SumOfSquares(i)
}
