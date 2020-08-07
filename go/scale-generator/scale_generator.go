package scale

import (
	"strings"
)

// Scale returns a scale based on a tonic t  and an interval string i
func Scale(tonic string, steps string) []string {
	var diatomic []string
	switch tonic {
	case "C", "G", "D", "A", "E", "B", "F#", "a", "e", "b", "f#", "c#", "g#", "d#":
		diatomic = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		diatomic = []string{"F", "Gb", "G", "Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E"}
	}
	if steps == "" {
		steps = strings.Repeat(" ", 12)
	}

	var currentI int
	var scale []string

	if len(tonic) == 2 {
		tonic = strings.ToUpper(string(tonic[0])) + string(tonic[1])
	} else {
		tonic = strings.ToUpper(string(tonic[0]))
	}

	for i, v := range diatomic {
		if tonic == v {
			currentI = i
			scale = append(scale, v)
		}
	}

	for _, step := range steps[:len(steps)-1] {
		switch step {
		case 'M':
			currentI += 2
		case 'A':
			currentI += 3
		default:
			currentI++
		}
		scale = append(scale, diatomic[(currentI)%12])

	}

	return scale
}
