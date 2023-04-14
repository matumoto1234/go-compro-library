package algorithm

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a []T) T {
	m := a[0]
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
}

func Max[T constraints.Ordered](a []T) T {
	m := a[0]
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
}
