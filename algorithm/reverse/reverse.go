package reverse

func Do[T any](a []T) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func Did[T any](a []T) []T {
	b := make([]T, len(a))
	copy(b, a)
	Do(b)
	return b
}
