package prime

import "math"

// Nth returns the nth prime
func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}
	var i int
	count := 0

	for count != n {
		i++
		if IsPrime(i) {
			count++
		}
	}
	return i, true
}

// IsPrime determines whether and integer is prime or not
func IsPrime(i int) bool {
	if i == 1 {
		return false
	}
	if i == 2 {
		return true
	}
	for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}
