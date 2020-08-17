package letter

import (
	"sync"
)

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
	c := make(chan FreqMap, 3)
	var wg sync.WaitGroup
	wg.Add(3)

	for _, s := range l {
		go func(s string) {
			defer wg.Done()
			c <- Frequency(s)
		}(s)
	}
	wg.Wait()
	close(c)
	for n := range c {
		m = Add(m, n)

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
