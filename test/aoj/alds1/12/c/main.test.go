// verification-helper: PROBLEM http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_12_C

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

	var n int
	fmt.Fscan(stdin, &n)

	g := graph.NewAdjacencyList[int](n)

	for i := 0; i < n; i++ {
		var u, k int
		fmt.Fscan(stdin, &u, &k)

		for j := 0; j < k; j++ {
			var v, c int
			fmt.Fscan(stdin, &v, &c)

			g.AddEdge(graph.NewEdge(
				u,
				v,
				c,
			))
		}
	}

	d := dijkstra.NewOrdered[int](g, 0, 1<<60)

	for v := 0; v < n; v++ {
		fmt.Fprintf(stdout, "%d %d\n", v, d.Distances[v])
	}
}
