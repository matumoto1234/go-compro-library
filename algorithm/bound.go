package algorithm

import "golang.org/x/exp/constraints"

// LowerBound() : return min({p | v <= a[p]})
// requires a to be sorted
func LowerBound[T constraints.Ordered](a []T, v T) int {
	if a[len(a)-1] < v {
		return len(a)
	}

	// [l, r)
	l, r := 0, len(a)
	for r-l > 1 {
		m := (l + r) / 2
		if a[m] <= v {
			l = m
		} else {
			r = m
		}
	}
	return l
}

// UpperBound() : return min({p | v < a[p]})
// requires a to be sorted
func UpperBound[T constraints.Ordered](a []T, v T) int {
	if a[len(a)-1] <= v {
		return len(a)
	}

	// (l, r]
	l, r := -1, len(a)-1
	for r-l > 1 {
		m := (l + r) / 2
		if a[m] <= v {
			l = m
		} else {
			r = m
		}
	}
	return r
}
