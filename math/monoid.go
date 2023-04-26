package math

var _ Monoid[int] = (*monoidImpl[int])(nil)

type Monoid[T any] interface {
	Op(T, T) T // binary operation
	E() T      // identity element
}

type monoidImpl[T any] struct {
	op func(T, T) T
	e  func() T
}

func (m *monoidImpl[T]) Op(a, b T) T {
	return m.op(a, b)
}

func (m *monoidImpl[T]) E() T {
	return m.e()
}

func NewMonoid[T any](op func(T, T) T, e func() T) *monoidImpl[T] {
	return &monoidImpl[T]{op, e}
}
