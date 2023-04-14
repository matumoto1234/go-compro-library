package data_structure

import (
	"github.com/matumoto1234/go-compro-library/internal"
	"github.com/matumoto1234/go-compro-library/utility"
)

type SegmentTree[T any] struct {
	m       utility.Monoid[T]
	oldSize int
	size    int
	log     int
	Vs      []T
}

func NewSegmentTree[T any](n int, m utility.Monoid[T]) *SegmentTree[T] {
	return NewSegmentTreeWithSlice(make([]T, n), m)
}

func NewSegmentTreeWithSlice[T any](a []T, m utility.Monoid[T]) *SegmentTree[T] {
	size := internal.BitCeil(uint64(len(a)))
	vs := make([]T, 2*size)

	for i := 0; i < len(a); i++ {
		vs[size+i] = a[i]
	}

	for i := size - 1; i >= 1; i-- {
		vs[i] = m.Op(vs[2*i], vs[2*i+1])
	}

	return &SegmentTree[T]{
		m:       m,
		oldSize: len(a),
		size:    size,
		log:     internal.CountRightZero(uint64(size)),
		Vs:      vs,
	}
}

func (s *SegmentTree[T]) Set(p int, x T) {
	utility.Assert(0 <= p && p < s.oldSize, utility.AssertMsg("index out of range"))

	p += s.size
	s.Vs[p] = x
	for i := 1; i <= s.log; i++ {
		s.Vs[p>>i] = s.m.Op(s.Vs[p>>(i-1)], s.Vs[p>>(i-1)^1])
	}
}

func (s *SegmentTree[T]) Get(p int) T {
	utility.Assert(0 <= p && p < s.oldSize, utility.AssertMsg("index out of range"))

	return s.Vs[p+s.size]
}

func (s *SegmentTree[T]) Prod(l, r int) T {
	utility.Assert(0 <= l && l <= r && r <= s.oldSize, utility.AssertMsg("index out of range"))

	l += s.size
	r += s.size

	prodL, prodR := s.m.E(), s.m.E()
	for l < r {
		if l&1 == 1 {
			prodL = s.m.Op(prodL, s.Vs[l])
			l++
		}
		if r&1 == 1 {
			r--
			prodR = s.m.Op(s.Vs[r], prodR)
		}
		l >>= 1
		r >>= 1
	}
	return s.m.Op(prodL, prodR)
}

func (s *SegmentTree[T]) AllProd() T {
	return s.Vs[1]
}

// MaxRight returns the maximum r (l <= r)
// such that f(op(a[l], a[l+1], ..., a[r-1])) = true.
func (s *SegmentTree[T]) MaxRight(l int, f func(T) bool) int {
	utility.Assert(0 <= l && l <= s.oldSize, utility.AssertMsg("index out of range"))
	utility.Assert(f(s.m.E()), utility.AssertMsg("f(e) must be true for some e"))

	if l == s.oldSize {
		return s.oldSize
	}
	l += s.size
	sm := s.m.E()
	for {
		for l%2 == 0 {
			l >>= 1
		}
		if !f(s.m.Op(sm, s.Vs[l])) {
			for l < s.size {
				l = 2 * l
				if f(s.m.Op(sm, s.Vs[l])) {
					sm = s.m.Op(sm, s.Vs[l])
					l++
				}
			}
			return l - s.size
		}
		sm = s.m.Op(sm, s.Vs[l])
		l++
		if l&-l == l {
			break
		}
	}
	return s.oldSize
}

// MinLeft returns the minimum l (l <= r)
// such that f(op(a[l], a[l+1], ..., a[r-1])) = true.
func (s *SegmentTree[T]) MinLeft(r int, f func(T) bool) int {
	utility.Assert(0 <= r && r <= s.oldSize, utility.AssertMsg("index out of range"))
	utility.Assert(f(s.m.E()), utility.AssertMsg("f(e) must be true for some e"))

	if r == 0 {
		return 0
	}
	r += s.size
	sm := s.m.E()
	for {
		r--
		for r > 1 && r%2 == 1 {
			r >>= 1
		}
		if !f(s.m.Op(s.Vs[r], sm)) {
			for r < s.size {
				r = 2*r + 1
				if f(s.m.Op(s.Vs[r], sm)) {
					sm = s.m.Op(s.Vs[r], sm)
					r--
				}
			}
			return r + 1 - s.size
		}
		sm = s.m.Op(s.Vs[r], sm)
		if r&-r == r {
			break
		}
	}
	return 0
}
