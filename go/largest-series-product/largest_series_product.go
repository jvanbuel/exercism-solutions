package lsproduct

import (
	"errors"
	"log"
	"regexp"
	"strconv"
)

// LargestSeriesProduct returns the largest product of n consecutive integers in a string of integers ds
func LargestSeriesProduct(ds string, span int) (int64, error) {
	if span < 0 {
		return -1, errors.New("span must be larger than 0")
	}
	if span > len(ds) {
		return -1, errors.New("span must be smaller than length of string")
	}
	if len(regexp.MustCompile("[^0-9]").FindAllString(ds, -1)) != 0 {
		return -1, errors.New("ds can only contain digits")
	}

	if span == len(ds) {
		mult := int64(1)
		for _, c := range ds {
			d, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal()
			}
			mult *= int64(d)
		}
		return mult, nil
	}
	var largest int64

	for i := 0; i <= len(ds)-span; i++ {
		seg := ds[i : i+span]
		mult := int64(1)
		for _, c := range seg {
			d, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal()
			}
			mult *= int64(d)
		}
		if mult > largest {
			largest = mult
		}
	}

	return largest, nil
}
