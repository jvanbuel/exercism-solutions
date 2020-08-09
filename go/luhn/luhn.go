package luhn

import (
	"strconv"
	"strings"
)

// Valid determines whether a string s is a valid Luhn number
func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)

	if len(s) <= 1 {
		return false
	}

	var sum int

	even := len(s)%2 == 0

	for _, c := range s {
		digit, err := strconv.Atoi(string(c))
		if err != nil {
			return false
		}
		if even {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		even = !even
		sum += digit
	}
	return sum%10 == 0
}
