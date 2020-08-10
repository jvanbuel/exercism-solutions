package atbash

import (
	"regexp"
	"strings"
)

// Atbash uses the atbash cipher to encode a string s
func Atbash(s string) string {
	s = strings.ToLower(regexp.MustCompile("[^a-zA-Z0-9]").ReplaceAllString(s, ""))

	m := map[rune]string{
		'a': "z",
		'b': "y",
		'c': "x",
		'd': "w",
		'e': "v",
		'f': "u",
		'g': "t",
		'h': "s",
		'i': "r",
		'j': "q",
		'k': "p",
		'l': "o",
		'm': "n",
		'n': "m",
		'o': "l",
		'p': "k",
		'q': "j",
		'r': "i",
		's': "h",
		't': "g",
		'u': "f",
		'v': "e",
		'w': "d",
		'x': "c",
		'y': "b",
		'z': "a",
		'0': "0",
		'1': "1",
		'2': "2",
		'3': "3",
		'4': "4",
		'5': "5",
		'6': "6",
		'7': "7",
		'8': "8",
		'9': "9",
	}
	var encoded string
	var i int
	for _, c := range s {
		encoded += m[c]
		i++
		if i%5 == 0 && i != len(s) {
			encoded += " "
		}
	}

	return encoded
}
