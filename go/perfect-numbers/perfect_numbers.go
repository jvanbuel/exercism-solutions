package perfect

import "errors"

// Classification is one of abundant, perect, or deficient
type Classification int

// Definition of the different classification types
const (
	ClassificationAbundant  = 0
	ClassificationPerfect   = 1
	ClassificationDeficient = 2
)

// ErrOnlyPositive is an error
var ErrOnlyPositive error = errors.New("n should be a natural number")

// Classify classifies an integer as abundant, perfect, or deficient
func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}
	var sum int64
	for i := int64(1); i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	if sum > n {
		return ClassificationAbundant, nil
	}
	if sum == n {
		return ClassificationPerfect, nil
	}
	return ClassificationDeficient, nil
}
