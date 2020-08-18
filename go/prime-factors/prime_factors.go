package prime

// Factors returns the prime factors of a number n
func Factors(n int64) []int64 {
	factors := make([]int64, 0)
	for i := int64(2); i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	return factors
}
