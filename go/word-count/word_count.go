package wordcount

import (
	"log"
	"regexp"
	"strings"
)

// Frequency is a map containing the frequencies of each word in a list
type Frequency map[string]int

// WordCount returns a frequency map of the words in a string s
func WordCount(s string) Frequency {
	f := Frequency{}
	s = strings.ToLower(s)

	t := regexp.MustCompile(`[ \t\n\r\,]`)
	words := t.Split(s, -1)

	reg, err := regexp.Compile("[^a-zA-Z0-9\\']+|(\\'\\B)|(^\\')")
	if err != nil {
		log.Fatal(err)
	}

	for _, w := range words {
		if w != "" {
			w = reg.ReplaceAllString(w, "")
			f[w]++
		}

	}

	return f
}
