package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in a given text concurrently and returns this
// data as a FreqMap.
func ConcurrentFrequency(l []string) FreqMap {
	m := FreqMap{}
	bufferSize := 2
	jobs := make(chan string, bufferSize)
	results := make(chan FreqMap, bufferSize)
	const nWorkers int = 3

	defer close(results)
	defer close(jobs)

	for i := 0; i < nWorkers; i++ {
		go func() {
			for j := range jobs {
				results <- Frequency(j)
			}
		}()
	}

	for _, s := range l {
		go func(s string) {
			jobs <- s
		}(s)
	}

	for range l {
		m = Add(m, <-results)
	}

	return m
}

// Add combines two FreqMaps into one
func Add(m FreqMap, n FreqMap) FreqMap {
	if len(m) == 0 {
		return n
	}
	for k, v := range n {
		m[k] += v
	}
	return m
}
