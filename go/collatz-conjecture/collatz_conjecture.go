package collatzconjecture

import (
	"errors"
)

// CollatzConjecture counts the number of iterations needed to get to 1 when following the Collatz procedure
func CollatzConjecture(i int) (int, error) {
	if i <= 0 {
		return 0, errors.New("i should be larger than 0")
	}
	var count int
	for i != 1 {
		if i%2 == 0 {
			i = i / 2
			count++
			continue
		}
		i = 3*i + 1
		count++

	}
	return count, nil
}
