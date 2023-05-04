package tuple2

import "golang.org/x/exp/constraints"

type Tuple2[T1, T2 constraints.Ordered] struct {
	V1 T1
	V2 T2
}

func New[T1, T2 constraints.Ordered](v1 T1, v2 T2) *Tuple2[T1, T2] {
	return &Tuple2[T1, T2]{
		V1: v1,
		V2: v2,
	}
}

func Less[T1, T2 constraints.Ordered](a, b *Tuple2[T1, T2]) bool {
	if a.V1 != b.V1 {
		return a.V1 < b.V1
	}
	return a.V2 < b.V2
}
