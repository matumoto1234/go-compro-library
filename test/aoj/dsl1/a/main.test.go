// verification-helper: PROBLEM https://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=DSL_1_A&lang=ja

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/matumoto1234/go-compro-library/data_structure"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)
	defer stdout.Flush()

	var n, q int
	fmt.Fscan(stdin, &n, &q)

	uf := data_structure.NewUnionFind(n)

	for i:=0;i<q;i++{
		var com, x, y int
		fmt.Fscan(stdin, &com, &x, &y)

		switch com {
		case 0:
			uf.Merge(x, y)
		case 1:
			if uf.Same(x, y) {
				fmt.Fprintln(stdout, 1)
			} else {
				fmt.Fprintln(stdout, 0)
			}
		}
	}
}
