package accumulator

import "github.com/matumoto1234/go-compro-library/math/group"

type Accumulator[T any] struct {
	Sum []T
	Abe group.Abelian[T]
}

func New[T any](a []T, abe group.Abelian[T]) *Accumulator[T] {
	sum := make([]T, len(a)+1)

	for i, v := range a {
		sum[i+1] = abe.Op(sum[i], v)
	}

	return &Accumulator[T]{
		Sum: sum,
		Abe: abe,
	}
}

// [l, r)
func (a *Accumulator[T]) Range(l, r int) T {
	return a.Abe.Op(a.Sum[r], a.Abe.Inv(a.Sum[l]))
}

// NewSum() : constructs a new Accumulator for int.
// op : +
// e : 0
// inv : -
func NewSum(a []int) *Accumulator[int] {
	return New[int](
		a,
		group.NewAbelian(
			func(x, y int) int {
				return x + y
			},
			func() int {
				return 0
			},
			func(x int) int {
				return -x
			},
		),
	)
}
