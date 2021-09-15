package allyourbase

import "errors"

// ConvertToBase converts a number in base a to a number in base b
func ConvertToBase(a int, digits []int, b int) ([]int, error) {
	var converted []int
	if a <= 2 {
		return converted, errors.New("input base must be >= 2")
	}
	if b <= 2 {
		return converted, errors.New("output base must be >= 2")
	}
	if len(digits) == 0 {
		return []int{0}, nil
	}
	for _, d := range digits {
		if d >= a {
			return converted, errors.New("all digits must satisfy 0 <= d < input base")
		}

	}
	return converted, nil
}
