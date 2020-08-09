package grains

import (
	"errors"
)

// Square returns the i-th power of 2
func Square(i int) (uint64, error) {
	if i <= 0 || i > 64 {
		return 0, errors.New("i should be larger than zero and smaller than 64")
	}
	return 1 << (i - 1), nil
}

// Total returns the total number of grains
func Total() uint64 {
	return ^uint64(0)
}
