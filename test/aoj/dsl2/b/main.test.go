// verification-helper: PROBLEM http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=DSL_2_B

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/matumoto1234/go-compro-library/data_structure"
	"github.com/matumoto1234/go-compro-library/util"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)
	defer stdout.Flush()

	var n, q int
	fmt.Fscan(stdin, &n, &q)

	st := data_structure.NewSegmentTree(n, util.NewMonoid(
		func(a, b int) int {
			return a + b
		},
		func() int {
			return 0
		},
	))

	for i := 0; i < q; i++ {
		var com, x, y int
		fmt.Fscan(stdin, &com, &x, &y)
		x--

		switch com {
		case 0:
			st.Set(x, st.Get(x)+y)
		case 1:
			fmt.Fprintln(stdout, st.Prod(x, y))
		}
	}
}
