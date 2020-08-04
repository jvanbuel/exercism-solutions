package strain

// Ints is a list of integers
type Ints []int

// Lists is a list of lists of integers
type Lists [][]int

// Strings is a list of strings
type Strings []string

// Keep keeps integers that conform to a predicate
func (is Ints) Keep(f func(int) bool) Ints {
	var ints Ints
	for _, i := range is {
		if f(i) {
			ints = append(ints, i)
		}
	}
	return ints
}

// Discard discards integers that conform to a predicate
func (is Ints) Discard(f func(int) bool) Ints {
	var ints Ints
	for _, i := range is {
		if !f(i) {
			ints = append(ints, i)
		}
	}
	return ints
}

// Keep keeps lists of integers that conform to a predicate
func (ls Lists) Keep(f func([]int) bool) Lists {
	var intsList Lists
	for _, ints := range ls {
		if f(ints) {
			intsList = append(intsList, ints)
		}
	}
	return intsList
}

// Keep keeps strings that confrom to a predicate
func (ss Strings) Keep(f func(string) bool) Strings {
	var strings Strings
	for _, s := range ss {
		if f(s) {
			strings = append(strings, s)
		}
	}
	return strings
}
