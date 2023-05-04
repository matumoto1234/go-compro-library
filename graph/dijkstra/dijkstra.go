package dijkstra

import (
	"github.com/matumoto1234/go-compro-library/datastructure/priorityqueue"
	"github.com/matumoto1234/go-compro-library/graph"
	"golang.org/x/exp/constraints"
)

type vAndDist[T any] struct {
	v    int
	dist T
}

func Ordered[T constraints.Ordered](g graph.Interface, inf T) ([]T, []int) {
	n := g.Size()
	dists := make([]T, n)
	prevs := make([]int, n)
	for i := 0; i < n; i++ {
		dists[i] = inf
		prevs[i] = -1
	}

	pq := priorityqueue.New(func(a, b vAndDist[T]) bool {
		return a.dist < b.dist
	})

	dists[0] = T(0)
	pq.Push(vAndDist[T]{
		v:    0,
		dist: dists[0],
	})

	for pq.Len() > 0 {
		d := pq.Pop()

		if dists[d.v] < d.dist {
			continue
		}

		for _, to := range g.Neighbors(d.v) {
			if dists[to] > dists[d.v]+d.dist {
				dists[to] = dists[d.v] + d.dist
				prevs[to] = d.v
				pq.Push(vAndDist[T]{
					v:    to,
					dist: dists[to],
				})
			}
		}
	}

	return dists, prevs
}
