package armstrong

import (
	"strconv"
)

// IsNumber determines whether an integer n is an Armstrong number or not
func IsNumber(n int) bool {
	var sum int
	l := len(strconv.Itoa(n))
	copy := n

	for i := 0; i < l; i++ {
		digit := copy % 10
		copy /= 10
		add := 1
		for k := 0; k < l; k++ {
			add *= digit
		}
		sum += add
	}

	return sum == n
}
