// verification-helper: PROBLEM http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_5_D

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/matumoto1234/go-compro-library/datastructure/compressor"
	"github.com/matumoto1234/go-compro-library/datastructure/fenwicktree"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)
	defer stdout.Flush()

	var n int
	fmt.Fscan(stdin, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(stdin, &a[i])
	}

	c := compressor.New(a)
	f := fenwicktree.New[int](n)

	var ans int64
	for i, v := range a {
		cv := c.Do(v)
		f.Add(cv, 1)
		ans += int64(i + 1 - f.Sum(0, cv+1))
	}

	fmt.Fprintln(stdout, ans)
}
