package pythagorean

// Triplet is a triplet of pythagorean integers
type Triplet [3]int

// Range returns a list of Pythagorean Triplets between the integers min and max
func Range(min, max int) []Triplet {
	var triplets []Triplet
	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			for k := j; k <= max; k++ {
				if i*i+j*j == k*k {
					triplets = append(triplets, Triplet{i, j, k})
				}

			}
		}
	}
	return triplets
}

// Sum returns a list of Pythagorean Tripelets whose perimeter (i.e. sum) equals p
func Sum(p int) []Triplet {
	var triplets []Triplet
	for i := 1; i <= p; i++ {
		for j := i; j <= p-i; j++ {
			if i*i+j*j == (p-i-j)*(p-i-j) {
				triplets = append(triplets, Triplet{i, j, p - i - j})
			}
		}
	}
	return triplets
}
