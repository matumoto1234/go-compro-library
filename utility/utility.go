package utility

import (
	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func ChMax[T constraints.Ordered](a *T, b T) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}

func ChMin[T constraints.Ordered](a *T, b T) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}
