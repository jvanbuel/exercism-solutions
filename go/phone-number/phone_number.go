package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

// Number returns the number of a telephone string number, removing punctuation and country code
func Number(s string) (string, error) {
	n := regexp.MustCompile("[^0-9]").ReplaceAllString(s, "")
	if len(n) < 10 || len(n) > 11 {
		return "", errors.New("a phone number must have 10 numbers")
	}
	if len(n) == 11 {
		if n[0] != '1' {
			return "", errors.New("a phone number must have 10 numbers, or 11 if it starts with 1")
		}
		n = n[1:]
	}
	if n[0] <= '1' || n[3] <= '1' {
		return "", errors.New("area code cannot start with 0 or 1, same for exchange code")
	}
	return n, nil
}

// AreaCode return the area code of a telephone number string
func AreaCode(s string) (string, error) {
	n, err := Number(s)
	if err != nil {
		return "", err
	}
	return n[:3], nil
}

// Format formats a telephone number string according to a specific format
func Format(s string) (string, error) {
	n, err := Number(s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", n[:3], n[3:6], n[6:]), nil
}
