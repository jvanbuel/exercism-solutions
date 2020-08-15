package sieve

// Sieve finds all primes smaller or equal to n
func Sieve(n int) []int {
	if n <= 0 {
		return []int{}
	}
	r := make([]int, n-1)
	for i := range r {
		r[i] = i + 2
	}

	for i := 0; i < len(r); i++ {
		if r[i] == 0 {
			continue
		}
		j := i + r[i]
		for j < len(r) {
			r[j] = 0
			j += r[i]
		}
	}

	var primes []int
	for _, v := range r {
		if v != 0 {
			primes = append(primes, v)
		}
	}
	return primes
}
