package data_structure

import (
	"fmt"

	"github.com/matumoto1234/go-compro-library/internal/bit"
	"github.com/matumoto1234/go-compro-library/util"
)

type FenwickTree[T util.Number] struct {
	n  int
	vs []T
}

func NewFenwickTree[T util.Number](n int) *FenwickTree[T] {
	vs := make([]T, n)
	return &FenwickTree[T]{n, vs}
}

// Add() : add x to p-th element
// p : 0-indexed
func (f *FenwickTree[T]) Add(p int, x T) {
	util.Assert(0 <= p && p < f.n, util.AssertMsg("FenwickTree.Add() : p is out of range"))
	p++ // to 1-indexed
	for p <= f.n {
		// p-1 : 0-indexed
		f.vs[p-1] += x
		p += p & -p
	}
}

// Sum() : return sum of [l, r)
// l, r : 0-indexed
func (f *FenwickTree[T]) Sum(l, r int) T {
	util.Assert(
		0 <= l && l <= r && r <= f.n,
		util.AssertMsg(fmt.Sprintf("FenwickTree.Sum() : l or r is out of range. l : %d, r : %d, f.n : %d", l, r, f.n)),
	)
	// [0, r) = [0, r-1] = f.sum(r-1)
	// [0, l) = [0, l-1] = f.sum(l-1)
	// [l, r) = [0, r) - [0, l)
	return f.sum(r-1) - f.sum(l-1)
}

// LowerBound() : return min({p | w <= f.sum(p)})
// requires f.Sum() to monotonically increasing
//
//	i.e. f.Sum(0, p) <= f.Sum(0, p+1) for all p
func (f *FenwickTree[T]) LowerBound(w T) int {
	if w <= 0 {
		return 0
	}

	p := 0
	r := int(bit.BitCeil(uint64(f.n))) // range length
	s := T(0)
	for ; r > 0; r >>= 1 {
		i := p + r - 1 // 0-indexed
		if i >= f.n {
			continue
		}

		if s+f.vs[i] <= w {
			s += f.vs[i]
			p += r - 1 // 0-indexed
		}
	}
	return p
}

// UpperBound() : return min({p | w < f.sum(p)})
// requires f.Sum() to monotonically increasing
//
//	i.e. f.Sum(0, p) <= f.Sum(0, p+1) for all p
func (f *FenwickTree[T]) UpperBound(w T) int {
	return f.LowerBound(w + 1)
}

// sum() : return sum of [0, p]
// p : 0-indexed
func (f *FenwickTree[T]) sum(p int) T {
	if p < 0 {
		return T(0)
	}

	p++ // to 1-indexed
	s := T(0)
	for p > 0 {
		// p-1 : 0-indexed
		s += f.vs[p-1]
		p -= p & -p
	}
	return s
}
