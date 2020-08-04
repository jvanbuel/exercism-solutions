package anagram

import (
	"reflect"
	"strings"
)

// Detect detects whether there exist anagrams of a string s in a list of other strings o
func Detect(s string, os []string) []string {
	s = strings.ToLower(s)
	var detected []string

	if (s == "") || (len(os) == 0) {
		return detected
	}

	for _, O := range os {
		o := strings.ToLower(O)
		if o == s {
			continue
		}
		if len(s) != len(o) {
			continue
		}
		if reflect.DeepEqual(ToMap(s), ToMap(o)) {
			detected = append(detected, O)
		}

	}
	return detected
}

// ToMap returns a map containing the characters of a strings
func ToMap(s string) map[rune]int {
	m := map[rune]int{}
	for _, c := range s {
		m[c]++
	}
	return m
}
