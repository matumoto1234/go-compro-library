package min

import "github.com/matumoto1234/go-compro-library/math/constraints"

func Ordered[T constraints.Ordered](a []T) T {
	m := a[0]
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
}
