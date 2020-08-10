package cipher

import (
	"regexp"
	"strings"
)

// Caesar is a caesar cipher struct
type Caesar struct{}

// Shift is a shift cipher struct
type Shift struct {
	distance int
}

// Vigenere is a vigenere cipher struct
type Vigenere struct {
	key string
}

// NewCaesar returns a new Caesar cipher structure
func NewCaesar() Cipher {
	return Caesar{}
}

// NewShift returns a new Shift cipher structure
func NewShift(distance int) Cipher {
	if distance*distance >= 26*26 || distance == 0 {
		return nil
	}
	if distance < 0 {
		return Shift{distance: distance + 26}
	}
	return Shift{distance: distance}
}

// NewVigenere returns a new Vigenère cipher structure
func NewVigenere(key string) Cipher {
	if key == "" {
		return nil
	}
	if regexp.MustCompile("[^a-z]").MatchString(key) {
		return nil
	}
	if regexp.MustCompile("^a+$").MatchString(key) {
		return nil
	}
	return Vigenere{key: key}
}

// Decode decodes a string s according to the caesar cipher
func (c Caesar) Decode(s string) string {
	var decoded string
	for _, c := range s {
		decoded += string((c-'a'+rune(23))%26 + 'a')
	}
	return decoded
}

// Encode encodes a string s according to the caesar cipher
func (c Caesar) Encode(s string) string {
	s = regexp.MustCompile("[^a-z]").ReplaceAllString(strings.ToLower(s), "")
	var encoded string
	for _, c := range s {
		encoded += string((c-'a'+rune(3))%26 + 'a')
	}
	return encoded
}

// Decode decodes a string s according to a shift cipher
func (sh Shift) Decode(s string) string {
	var decoded string
	for _, c := range s {
		decoded += string((c-'a'+rune(26-sh.distance))%26 + 'a')
	}
	return decoded

}

// Encode encodes a string s according to a shift cipher
func (sh Shift) Encode(s string) string {
	s = regexp.MustCompile("[^a-z]").ReplaceAllString(strings.ToLower(s), "")
	var encoded string
	for _, c := range s {
		encoded += string((c-'a'+rune(sh.distance))%26 + 'a')
	}
	return encoded
}

// Decode decodes a string s according to a Vigenere cipher
func (v Vigenere) Decode(s string) string {
	s = regexp.MustCompile("[^a-z]").ReplaceAllString(strings.ToLower(s), "")
	var decoded string

	for i, c := range s {
		decoded += string((c-'a'+26-(rune(v.key[i%len(v.key)])-'a'))%26 + 'a')
	}
	return decoded
}

// Encode encodes a string s according to a Vigenere cipher
func (v Vigenere) Encode(s string) string {
	s = regexp.MustCompile("[^a-z]").ReplaceAllString(strings.ToLower(s), "")
	var encoded string

	for i, c := range s {
		encoded += string((c-'a'+rune(v.key[i%len(v.key)])-'a')%26 + 'a')
	}
	return encoded
}
