package encode

import (
	"log"
	"regexp"
	"strconv"
)

// RunLengthEncode encodes a string s ino a runlength string
func RunLengthEncode(s string) string {
	var encoded string
	if s == "" {
		return encoded
	}

	var previous rune
	var current rune
	var count int

	previous = []rune(s)[0]
	for _, c := range s {
		current = c
		if current == previous {
			count++
			continue
		}
		if count > 1 {
			encoded += strconv.Itoa(count) + string(previous)
			previous = current
			count = 1
			continue
		}
		encoded += string(previous)
		previous = current
	}

	// For the last character
	if count > 1 {
		encoded += strconv.Itoa(count) + string(previous)
	} else {
		encoded += string(previous)
	}

	return encoded
}

// RunLengthDecode decodes a runlength string s
func RunLengthDecode(s string) string {
	var decoded string

	if s == "" {
		return decoded
	}

	re := regexp.MustCompile("[0-9]+")

	nums := re.FindAllString(s, -1)
	strings := re.Split(s, -1)
	if len(nums) == 0 {
		return s
	}

	//	fmt.Printf("%q\n", nums)
	//	fmt.Printf("%q\n", strings)

	decoded += strings[0]

	for idx, n := range nums {
		ni, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal()
		}
		for i := 0; i < ni; i++ {
			decoded += string(strings[idx+1][0])
		}
		if len(strings[idx+1]) > 1 {
			decoded += string(strings[idx+1][1:])
		}
	}

	return decoded
}
