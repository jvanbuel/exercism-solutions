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
	jobs := make(chan string, len(l))
	results := make(chan FreqMap, len(l))
	const nWorkers int = 3

	defer close(results)
	defer close(jobs)

	for i := 0; i < nWorkers; i++ {
		go func(jobs <-chan string, results chan<- FreqMap) {
			for j := range jobs {
				results <- Frequency(j)
			}
		}(jobs, results)
	}

	for _, s := range l {
		jobs <- s
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
