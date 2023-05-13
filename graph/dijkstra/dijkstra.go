package dijkstra

import (
	"github.com/matumoto1234/go-compro-library/algorithm/reverse"
	"github.com/matumoto1234/go-compro-library/datastructure/priorityqueue"
	"github.com/matumoto1234/go-compro-library/graph"
	"golang.org/x/exp/constraints"
)

type vAndDist[T any] struct {
	v    int
	dist T
}

type Ordered[T constraints.Ordered] struct {
	G   graph.Interface[T]
	Inf T

	// 始点s
	S int

	// Distances[v] : 始点sから頂点vへの最短経路長
	Distances []T

	// PreviewEdges[v] : ある頂点uから頂点vへ遷移するために使用した辺(u,v,cost)
	PreviewEdges []*graph.Edge[T]
}

func NewOrdered[T constraints.Ordered](g graph.Interface[T], s int, inf T) *Ordered[T] {
	dists, prevs := dijkstra(g, s, inf)

	return &Ordered[T]{
		G:            g,
		S:            s,
		Inf:          inf,
		Distances:    dists,
		PreviewEdges: prevs,
	}
}

// RestorePath(v) : 始点sから頂点vへの最短経路を復元する
//
//	計算量 : O(N)
//	始点sが渡された場合、{s}を返す
func (o *Ordered[T]) RestorePath(v int) []int {
	path := []int{v}

	for v != o.S {
		path = append(path, v)
		v = o.PreviewEdges[v].From
	}

	reverse.Slice(path)
	return path
}

// ShortestPathTree() : 始点sを根とした最短経路有向木を返す
//
//	計算量 : O(N + M)
//	木の辺は根から葉に向かって有向である
func (o *Ordered[T]) ShortestPathTree() *graph.AdjacencyList[T] {
	n := o.G.Size()
	g := graph.NewAdjacencyList[T](n)

	used := make([]bool, n)

	for s := 0; s < n; s++ {
		v := s
		for v != o.S {
			if used[v] {
				break
			}

			used[v] = true

			prev := o.PreviewEdges[v].From

			g.AddEdge(graph.NewEdge(
				prev,
				v,
				o.PreviewEdges[v].Cost,
			))

			v = prev
		}
	}

	return g
}

func dijkstra[T constraints.Ordered](g graph.Interface[T], s int, inf T) ([]T, []*graph.Edge[T]) {
	n := g.Size()
	dists := make([]T, n)
	prevs := make([]*graph.Edge[T], n)
	for i := 0; i < n; i++ {
		dists[i] = inf
		prevs[i] = graph.NewEdge(-1, -1, T(0))
	}

	pq := priorityqueue.New(func(a, b vAndDist[T]) bool {
		return a.dist < b.dist
	})

	dists[s] = T(0)

	pq.Push(vAndDist[T]{
		v:    s,
		dist: dists[s],
	})

	for pq.Len() > 0 {
		d := pq.Pop()

		if dists[d.v] < d.dist {
			continue
		}

		edges := g.Edges(d.v)

		for _, e := range edges {
			if dists[e.To] > dists[d.v]+e.Cost {
				dists[e.To] = dists[d.v] + e.Cost
				prevs[e.To] = e
				pq.Push(vAndDist[T]{
					v:    e.To,
					dist: dists[e.To],
				})
			}
		}
	}

	return dists, prevs
}
