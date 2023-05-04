package graph

type Interface interface {
	Size() int
	Neighbors(int) []int
}

// var _ Graph = (*adjacencyList)(nil)

type adjacencyList struct {
	Data [][]int
}

func NewAdjacencyList(n int) *adjacencyList {
	return &adjacencyList{make([][]int, n)}
}
