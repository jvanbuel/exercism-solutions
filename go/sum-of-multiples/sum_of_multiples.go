package summultiples

// SumMultiples return the sum of the multiples of the integers is below a limit
func SumMultiples(limit int, is ...int) int {
	var sum int

	if len(is) == 0 {
		return sum
	}

	for i := 1; i < limit; i++ {
		for _, d := range is {
			if d == 0 {
				break
			}
			if i%d == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
