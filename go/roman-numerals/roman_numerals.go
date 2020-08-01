package romannumerals

import (
	"errors"
	"fmt"
)

// ToRomanNumeral returns the roman string representation of an arabic integer
func ToRomanNumeral(a int) (string, error) {
	var s string
	m := make(map[int]map[rune]string)
	m[0] = map[rune]string{
		'0': "",
		'1': "M",
		'2': "MM",
		'3': "MMM",
	}
	m[1] = map[rune]string{
		'0': "",
		'1': "C",
		'2': "CC",
		'3': "CCC",
		'4': "CD",
		'5': "D",
		'6': "DC",
		'7': "DCC",
		'8': "DCCC",
		'9': "CM",
	}
	m[2] = map[rune]string{
		'0': "",
		'1': "X",
		'2': "XX",
		'3': "XXX",
		'4': "XL",
		'5': "L",
		'6': "LX",
		'7': "LXX",
		'8': "LXXX",
		'9': "XC",
	}
	m[3] = map[rune]string{
		'0': "",
		'1': "I",
		'2': "II",
		'3': "III",
		'4': "IV",
		'5': "V",
		'6': "VI",
		'7': "VII",
		'8': "VIII",
		'9': "IX",
	}

	if a > 3000 || a <= 0 {
		return "", errors.New("int should be smaller than 3000")
	}
	num := fmt.Sprintf("%04d", a) // Pad with 0 to 4 digits
	for i, d := range num {
		s += m[i][d]
	}
	return s, nil
}
