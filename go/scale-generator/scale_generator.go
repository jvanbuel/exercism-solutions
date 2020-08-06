package scale

// Scale returns a scale based on a tonic t  and an interval string i
func Scale(t string, is string) []string {
	var d []string
	switch t {
	case "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#":
		d = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		d = []string{"F", "Gb", "G", "Ab", "A", "Bb", "B", "C", "Db", "E"}
	}
	if is == "" {
		return d
	}

	m := map[rune]int{'m': 1, 'M': 2}
	var scale []string
	startI := 

	for _, i := range is {
		step := m[i]

	}

	return scale
}
