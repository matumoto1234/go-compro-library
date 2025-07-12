package suffixarray

import (
	"sort"

	"github.com/matumoto1234/go-compro-library/string/rollinghash"
)

func NewSuffixArray(s string) []int {
	rh := rollinghash.NewRollingHash(s)

	sa := make([]int, len(s))
	for i := range sa {
		sa[i] = i
	}

	sort.Slice(sa, func(i, j int) bool {
		return rh.Less(sa[i], len(s), sa[j], len(s))
	})

	return sa
}
