// verification-helper: PROBLEM https://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=0516

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/matumoto1234/go-compro-library/algorithm"
	"github.com/matumoto1234/go-compro-library/data_structure"
)

var stdin = bufio.NewReader(os.Stdin)
var stdout = bufio.NewWriter(os.Stdout)

func solve() bool {
	var n, k int
	fmt.Fscan(stdin, &n, &k)

	if n == 0 {
		return false
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(stdin, &a[i])
	}

	cs := data_structure.NewCumulativeSum(a)
	ans := 0

	for i := 0; i < n; i++ {
		if i+k > n {
			break
		}
		ans = algorithm.Max([]int{ans, cs.Query(i, i+k)})
	}

	fmt.Fprintln(stdout, ans)

	return true
}

func main() {
	defer stdout.Flush()

	for solve() {
	}
}