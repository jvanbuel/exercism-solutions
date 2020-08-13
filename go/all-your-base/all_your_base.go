package allyourbase

import "errors"

// ConvertToBase converts a number in base a to a number in base b
func ConvertToBase(a int, digits []int, b int) ([]int, error) {
	if a <= 2 {
		return []int{}, errors.New("input base must be >= 2")
	}
	if b <= 2 {
		return []int{}, errors.New("output base must be >= 2")
	}
	if len(digits) == 0 {
		return []int{0}, nil
	}
	for _, d := range digits {
		if d >= a {
			return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
		}

	}
	return []int{}, nil
}
