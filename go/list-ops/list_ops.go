package listops

// IntList is a list of integers
type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Foldr folds a list of integers starting from the left with a binary reduction function f
func (is IntList) Foldr(f binFunc, i int) int {
	if is.Length() == 0 {
		return i
	}
	return is[:is.Length()-1].Foldr(f, f(is[is.Length()-1], i))
}

// Foldl folds a list of integers starting from the right with a binary reduction function f
func (is IntList) Foldl(f binFunc, i int) int {
	if is.Length() == 0 {
		return i
	}
	return is[1:].Foldl(f, f(i, is[0]))
}

// Filter filter a list of integers according to the predicate function f
func (is IntList) Filter(f predFunc) IntList {
	if is.Length() == 0 {
		return IntList{}
	}
	if f(is[is.Length()-1]) {
		return append(is[:is.Length()-1].Filter(f), is[is.Length()-1])
	}
	return is[:is.Length()-1].Filter(f)
}

// Length returns length of an IntList is
func (is IntList) Length() int {
	return len(is)
}

// Map maps a unary function f to all elements of an IntList
func (is IntList) Map(f unaryFunc) IntList {
	if is.Length() == 0 {
		return IntList{}
	}
	return append(is[:is.Length()-1].Map(f), f(is[is.Length()-1]))
}

// Reverse reverses the entries of an IntList
func (is IntList) Reverse() IntList {
	if is.Length() == 0 {
		return IntList{}
	}
	return append(is[1:].Reverse(), is[0])
}

// Append appends an IntList to another IntList
func (is IntList) Append(js IntList) IntList {
	if js.Length() == 0 {
		return is
	}
	return append(is, js[0]).Append(js[1:])
}

// Concat concatinates a list of IntLists to another IntList
func (is IntList) Concat(lis []IntList) IntList {
	if len(lis) == 0 {
		return is
	}
	return is.Append(lis[0]).Concat(lis[1:])
}
