package permutation

import (
	"github.com/matumoto1234/go-compro-library/algorithm/reverse"
	"golang.org/x/exp/constraints"
)

func Next[T constraints.Ordered](a []T) bool {
	for i := len(a) - 2; i >= 0; i-- {
		if a[i] < a[i+1] {
			j := len(a) - 1
			for a[i] >= a[j] {
				j--
			}
			a[i], a[j] = a[j], a[i]
			reverse.Do(a[i+1:])
			return true
		}
		if i == 0 {
			reverse.Do(a)
		}
	}
	return false
}

func Prev[T constraints.Ordered](a []T) bool {
	for i := len(a) - 2; i >= 0; i-- {
		if a[i] > a[i+1] {
			j := len(a) - 1
			for a[i] <= a[j] {
				j--
			}
			a[i], a[j] = a[j], a[i]
			reverse.Do(a[i+1:])
			return true
		}
		if i == 0 {
			reverse.Do(a)
		}
	}
	return false
}
