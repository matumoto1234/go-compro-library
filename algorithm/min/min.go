package min

import "golang.org/x/exp/constraints"

func Ordered[T constraints.Ordered](a []T) T {
	m := a[0]
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
}
