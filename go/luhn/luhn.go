package luhn

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Valid determines whether a string s is a valid Luhn number
func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)

	if len(s) <= 1 || regexp.MustCompile("[^0-9]").MatchString(s) {
		return false
	}

	var count, digit int
	var multip bool
	var err error

	if l := len(s); l%2 == 0 {
		multip = true
	}

	for _, c := range s {
		digit, err = strconv.Atoi(string(c))
		if err != nil {
			log.Fatal("could not convert character to digit")
		}
		if multip {
			dd := digit * 2
			if dd > 9 {
				count += dd - 9
			} else {
				count += dd
			}
			multip = false
		} else {
			count += digit
			multip = true
		}

	}
	return count%10 == 0
}
