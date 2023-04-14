package algorithm

import "golang.org/x/exp/constraints"

func NextPermutation[T constraints.Ordered](a []T) bool {
	for i := len(a) - 2; i >= 0; i-- {
		if a[i] < a[i+1] {
			j := len(a) - 1
			for a[i] >= a[j] {
				j--
			}
			a[i], a[j] = a[j], a[i]
			Reverse(a[i+1:])
			return true
		}
		if i == 0 {
			Reverse(a)
		}
	}
	return false
}

func PrevPermutation[T constraints.Ordered](a []T) bool {
	for i := len(a) - 2; i >= 0; i-- {
		if a[i] > a[i+1] {
			j := len(a) - 1
			for a[i] <= a[j] {
				j--
			}
			a[i], a[j] = a[j], a[i]
			Reverse(a[i+1:])
			return true
		}
		if i == 0 {
			Reverse(a)
		}
	}
	return false
}
