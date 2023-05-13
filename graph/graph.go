package graph

import "fmt"

type Interface[T any] interface {
	Size() int
	Edges(int) []*Edge[T]
	AddEdge(*Edge[T])
}

func _[T any]() {
	var _ Interface[T] = (*AdjacencyList[T])(nil)
}

type Edge[T any] struct {
	From int
	To   int
	Cost T
}

func (e *Edge[T]) String() string {
	return fmt.Sprintf("From:%v To:%v Cost:%v", e.From, e.To, e.Cost)
}

func NewEdge[T any](from, to int, cost T) *Edge[T] {
	return &Edge[T]{
		From: from,
		To:   to,
		Cost: cost,
	}
}

type AdjacencyList[T any] struct {
	NeighborEdges [][]*Edge[T]
}

func NewAdjacencyList[T any](n int) *AdjacencyList[T] {
	return &AdjacencyList[T]{
		make([][]*Edge[T], n),
	}
}

func (g *AdjacencyList[T]) Size() int {
	return len(g.NeighborEdges)
}

func (g *AdjacencyList[T]) Edges(v int) []*Edge[T] {
	return g.NeighborEdges[v]
}

func (g *AdjacencyList[T]) AddEdge(e *Edge[T]) {
	g.NeighborEdges[e.From] = append(g.NeighborEdges[e.From], e)
}
