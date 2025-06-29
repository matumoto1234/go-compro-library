// verification-helper: PROBLEM http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_14_A

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/matumoto1234/go-compro-library/string/rollinghash"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var t, p string
	fmt.Fscan(stdin, &t, &p)

	tHash := rollinghash.NewRollingHash(t)
	pHash := rollinghash.NewRollingHash(p)

	for i := 0; i < len(t)-len(p)+1; i++ {
		v1 := tHash.Find(i, i+len(p))
		v2 := pHash.Find(0, len(p))

		if v1 == v2 {
			fmt.Println(i)
		}
	}
}
