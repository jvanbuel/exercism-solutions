package sieve

// Sieve finds all primes smaller or equal to n
func Sieve(n int) []int {
	if n <= 0 {
		return []int{}
	}
	r := make([]int, n-1)
	var primes []int

	for i := range r {
		r[i] = i + 2
	}

	for i := 0; i < len(r); i++ {
		if r[i] == 0 {
			continue
		}
		primes = append(primes, r[i])
		j := i + r[i]
		for j < len(r) {
			r[j] = 0
			j += r[i]
		}
	}

	return primes
}
