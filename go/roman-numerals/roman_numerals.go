package romannumerals

import (
	"errors"
	"fmt"
)

// ToRomanNumeral returns the roman string representation of an arabic integer
func ToRomanNumeral(a int) (string, error) {
	if a > 3000 {
		return "", errors.New("int should be smaller than 3000")
	}
	for i := 0; i <= 3; i++ {
		fmt.Printf("%d ", (a%10^(i+1))/10^i)
	}
	println("")
	return "", nil
}
