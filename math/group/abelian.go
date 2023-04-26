package group

var _ Abelian[int] = (*abelianImpl[int])(nil)

type Abelian[T any] interface {
	Op(T, T) T // binary operation
	E() T      // identity element
	Inv(T) T   // inverse element
}

type abelianImpl[T any] struct {
	op  func(T, T) T
	e   func() T
	inv func(T) T
}

func (a *abelianImpl[T]) Op(x, y T) T {
	return a.op(x, y)
}

func (a *abelianImpl[T]) E() T {
	return a.e()
}

func (a *abelianImpl[T]) Inv(x T) T {
	return a.inv(x)
}

func NewAbelian[T any](op func(T, T) T, e func() T, inv func(T) T) *abelianImpl[T] {
	return &abelianImpl[T]{op, e, inv}
}
