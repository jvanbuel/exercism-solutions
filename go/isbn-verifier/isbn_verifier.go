package isbn

import (
	"log"
	"regexp"
	"strconv"
	"unicode"
)

// IsValidISBN determines whether a strings is a valid ISBN number
func IsValidISBN(s string) bool {
	if s == "" {
		return false
	}
	check := s[len(s)-1]
	if check != 'X' && !unicode.IsDigit(rune(check)) {
		return false
	}
	ns := regexp.MustCompile("[^0-9]").ReplaceAllString(s[:len(s)-1], "")

	if len(ns) != 9 {
		return false
	}

	var sum int
	for i, c := range ns {
		d, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatal()
		}
		sum += d * (10 - i)
	}

	if check == 'X' {
		sum += 10
	} else {
		d, err := strconv.Atoi(string(check))
		if err != nil {
			log.Fatal()
		}
		sum += d
	}
	return sum%11 == 0
}
