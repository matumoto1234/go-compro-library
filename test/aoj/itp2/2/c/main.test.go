// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/lesson/8/ITP2/2/ITP2_2_C

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/matumoto1234/go-compro-library/datastructure/priorityqueue"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)
	defer stdout.Flush()

	var n, q int
	fmt.Fscan(stdin, &n, &q)

	pqList := make([]*priorityqueue.PriorityQueue[int], n)

	for i := 0; i < n; i++ {
		pqList[i] = priorityqueue.New(func(a, b int) bool {
			return a > b
		})
	}

	for i := 0; i < q; i++ {
		var op, t int
		fmt.Fscan(stdin, &op, &t)

		switch op {
		case 0:
			var x int
			fmt.Fscan(stdin, &x)
			pqList[t].Push(x)
		case 1:
			if pqList[t].Len() != 0 {
				fmt.Fprintln(stdout, pqList[t].Top())
			}
		case 2:
			if pqList[t].Len() != 0 {
				pqList[t].Pop()
			}
		}
	}
}
