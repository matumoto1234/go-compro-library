package algorithm

func Reverse[T any](a []T) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func Reversed[T any](a []T) []T {
	b := make([]T, len(a))
	copy(b, a)
	Reverse(b)
	return b
}
