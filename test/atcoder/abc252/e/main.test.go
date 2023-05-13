// verification-helper: IGNORE
// verification-helper: PROBLEM https://atcoder.jp/contests/abc252/tasks/abc252_e

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/matumoto1234/go-compro-library/graph"
	"github.com/matumoto1234/go-compro-library/graph/dijkstra"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)
	defer stdout.Flush()

	var n, m int
	fmt.Fscan(stdin, &n, &m)

	g := graph.NewAdjacencyList[int](n)

	edgeToIndex := make(map[graph.Edge[int]]int)

	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(stdin, &a, &b, &c)

		a--
		b--

		e := graph.NewEdge(a, b, c)
		re := graph.NewEdge(b, a, c)

		g.AddEdge(e)
		g.AddEdge(re)
		edgeToIndex[*e] = i
		edgeToIndex[*re] = i
	}

	d := dijkstra.NewOrdered[int](g, 0, 1<<60)

	t := d.ShortestPathTree()

	edges := make(map[graph.Edge[int]]bool)

	for v := 0; v < n; v++ {
		for _, e := range t.Edges(v) {
			edges[*e] = true
		}
	}

	ans := make([]int, 0, len(edges))

	for e := range edges {
		ans = append(ans, edgeToIndex[e]+1)
	}

	for i, v := range ans {
		if i > 0 {
			fmt.Fprint(stdout, " ")
		}
		fmt.Fprint(stdout, v)
	}
	fmt.Fprint(stdout, "\n")
}
