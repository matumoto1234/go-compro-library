package graph

type Edge[T any] interface {
	From() int
	To() int
	Cost() T
}
